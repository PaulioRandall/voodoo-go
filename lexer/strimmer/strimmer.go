package strimmer

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

// Strim normalises an array of tokens ready for the
// the token parsing, this involves:
// -> Removing whitespace tokens
// -> Removing comment tokens
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Removing now redundant punctuation
// -> Converting all letters to lowercase (Except string literals)
func Strim(in []symbol.Token) []symbol.Token {

	out := []symbol.Token{}

	for _, tk := range in {
		switch {
		case tk.Type == symbol.WHITESPACE:
			continue
		case tk.Type == symbol.COMMENT:
			continue
		case tk.Type == symbol.LITERAL_STRING:
			penultimate := len(tk.Val) - 1
			tk.Val = tk.Val[1:penultimate]
		case tk.Type == symbol.LITERAL_NUMBER:
			tk.Val = strings.ReplaceAll(tk.Val, `_`, ``)
		case tk.Type > symbol.ALPHABETIC_START && tk.Type < symbol.ALPHABETIC_END:
			tk.Val = strings.ToLower(tk.Val)
		}

		out = append(out, tk)
	}

	return out
}
