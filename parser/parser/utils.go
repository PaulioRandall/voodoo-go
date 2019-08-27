package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// indexOf returns the index of the first token with the specified type.
func indexOf(in []token.Token, t token.TokenType) int {
	for i, tk := range in {
		if tk.Type == t {
			return i
		}
	}
	return -1
}
