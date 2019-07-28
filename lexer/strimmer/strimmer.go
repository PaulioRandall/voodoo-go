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
		case tk.Type > token.ALPHABETIC_START && tk.Type < token.ALPHABETIC_END:
			tk.Val = strings.ToLower(tk.Val)
		}

		out = append(out, tk)
	}

	return out
}
