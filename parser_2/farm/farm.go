package farm

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser_2/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// Farm represents an allotment of tokens representing a statement.
type Farm struct {
	multiline bool          // True if the next line feed can be ignore
	mature    bool          // True if a statement is ready for harvesting
	tokens    []token.Token // Set of tokens representing a statement
}

// New returns a new Farm.
func New() *Farm {
	return &Farm{
		tokens: []token.Token{},
	}
}

// Copy performs a deep copy of the farm.
func Copy(f *Farm) *Farm {
	return &Farm{
		multiline: f.multiline,
		mature:    f.mature,
		tokens:    token.Copy(f.tokens),
	}
}

// Feed processes a Token to remove redundant runes, converting its Kind if
// certain criteria are meet, and placing it in the Farms token field as part of
// the next statement if it's not redundant. Newline Tokens may cause the
// allotment to mature and the Farm be flagged as ready for harvesting; upon
// which, attempting to feed more Tokens will result in a panic unless a call to
// Harvest() is performed first.
func (ev *Farm) Feed(in token.Token) (bool, error) {
	if ev.mature {
		panic("Can't feed with any more Tokens until Harvest() is invoked")
	}

	if ev.tokens == nil {
		panic("Can't use a salted farm")
	}

	if ev.filter(in) {
		return ev.mature, nil
	}

	in = ev.strim(in)
	ev.tokens = append(ev.tokens, in)
	return ev.mature, nil
}

// Harvest returns a set of Tokens that make up a statement reseting the Farm
// in the process.
func (ev *Farm) Harvest() []token.Token {
	if !ev.mature {
		panic("Can't harvest until a newline Token has been sown")
	}

	defer ev.reset(false)
	return ev.tokens
}

// SaltHarvest returns the set of Tokens that make up a statement even if a
// newline Token was not supplied in the previous invocation of Feed(). The
// tokens is salted in the process rendering the Farm unusable.
func (ev *Farm) SaltHarvest() []token.Token {
	defer ev.reset(true)
	return ev.tokens
}

// reset resets the Farm so a new statement can be started. If the input is true
// then Farm will be rendered unusable.
func (ev *Farm) reset(salt bool) {
	ev.multiline = false
	ev.mature = false
	if salt {
		ev.tokens = nil
	} else {
		ev.tokens = []token.Token{}
	}
}

// filter removes redundant Tokens.
func (ev *Farm) filter(in token.Token) bool {
	switch in.Kind() {
	case token.TT_NEWLINE:
		if !ev.multiline {
			ev.mature = true
		}
	case token.TT_SHEBANG:
	case token.TT_SPACE:
	default:
		return false
	}

	return true
}

// strim normalises a token. This may involve modifying the Token ready for
// parsing.
func (ev *Farm) strim(in token.Token) token.Token {
	switch in.Kind() {
	case token.TT_ID:
		return toLower(in)
	default:
		return in
	}
}

// toLower returns the input token but with all the characters in the value
// field converted to lowercase.
func toLower(tk token.Token) token.Token {
	s := strings.ToLower(tk.Text())
	return scantok.UpdateText(tk, s)
}
