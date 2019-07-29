package strimmer

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/token"
)

// Strim normalises an array of tokens ready for the
// the token parsing, this involves:
// -> Removing whitespace tokens
// -> Removing comment tokens
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Removing now redundant punctuation
// -> Converting all letters to lowercase (Except string literals)
func Strim(in []token.Token) []token.Token {

	out := []token.Token{}

	for _, tk := range in {
		switch {
		case tk.Type == token.WHITESPACE:
			continue
		case tk.Type == token.COMMENT:
			continue
		case tk.Type == token.LITERAL_STRING:
			penultimate := len(tk.Val) - 1
			tk.Val = tk.Val[1:penultimate]
		case tk.Type == token.LITERAL_NUMBER:
			tk.Val = strings.ReplaceAll(tk.Val, `_`, ``)
		case isAlphabeticType(tk.Type):
			tk.Val = strings.ToLower(tk.Val)
		}

		out = append(out, tk)
	}

	return out
}

// isAlphabeticType returns true if input token type is for
// tokens that may have alphabetic values.
func isAlphabeticType(t token.TokenType) bool {
	switch t {
	case token.KEYWORD_FUNC:
	case token.KEYWORD_LOOP:
	case token.KEYWORD_WHEN:
	case token.KEYWORD_DONE:
	case token.BOOLEAN_TRUE:
	case token.BOOLEAN_FALSE:
	case token.SPELL:
	default:
		return false
	}

	return true
}
