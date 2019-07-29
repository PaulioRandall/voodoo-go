package scanner

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanComment scans symbols that start with a two consecutive
// forward slashes `//` returning the whole input as a comment.
func scanComment(in []rune) *token.Token {
	return &token.Token{
		Val:  string(in),
		Type: token.COMMENT,
	}
}
