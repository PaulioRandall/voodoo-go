package strim

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// strim removes or normalises tokens.
func Strim(tk token.Token) token.Token {
	if filter(tk) {
		return nil
	}
	return strim(tk)
}

// filter removes redundant Tokens.
func filter(in token.Token) bool {
	switch in.Kind() {
	case token.TT_SHEBANG:
	case token.TT_SPACE:
	default:
		return false
	}

	return true
}

// strim normalises a token. This may involve modifying the Token ready for
// parsing.
func strim(in token.Token) token.Token {
	switch in.Kind() {
	case token.TT_ID, token.TT_BOOL:
		return toLower(in)
	case token.TT_STRING:
		return stripEnds(in)
	default:
		return in
	}
}

// stripEnds returns the token but with the first and last rune of the tokens
// text removed.
func stripEnds(tk token.Token) token.Token {
	s := tk.Text()[1 : len(tk.Text())-1]
	return token.UpdateText(tk, s)
}

// toLower returns the input token but with all the characters in the value
// field converted to lowercase.
func toLower(tk token.Token) token.Token {
	s := strings.ToLower(tk.Text())
	return token.UpdateText(tk, s)
}
