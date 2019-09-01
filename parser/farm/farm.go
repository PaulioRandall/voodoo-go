package farm

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/scantok"
	"github.com/PaulioRandall/voodoo-go/parser/token"
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
// which, attempting to feed more Tokens will result in a panic; a new farm must
// be created.
func (ev *Farm) Feed(in token.Token) (bool, error) {
	if ev.mature {
		panic("Can't feed with any more Tokens, harvest is required")
	}

	if ev.tokens == nil {
		panic("Can't use a harvested farm")
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

	defer func() {
		ev.multiline = false
		ev.mature = false
		ev.tokens = nil
	}()
	return ev.tokens
}

// FinalHarvest performs a harvest that will not panic if the allotment is not
// mature.
func (ev *Farm) FinalHarvest() []token.Token {
	ev.mature = true
	return ev.Harvest()
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
