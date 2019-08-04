package scanner

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSpace scans tokens that start with a unicode whitespace
// property rune returning a token representing all whitespace
// between two non-whitespace tokens.
//
// Note that there is an intention to switch to stream based
// scanning. When this change happens newline runes will
// become the one exception to the rule as they will become
// a token all by themselves used to delimit statements
// and the bodies of different context.
func scanSpace(in []rune, col int) (tk *token.Token, out []rune) {

	end := findNonWhitespace(in)
	if end == -1 {
		end = len(in)
	}

	tk = &token.Token{
		Val:   string(in[:end]),
		Start: col,
		Type:  token.WHITESPACE,
	}

	out = in[end:]
	return
}

// findNonWhitespace finds the first non-whitespace index.
func findNonWhitespace(in []rune) int {
	for i, v := range in {
		if !isSpace(v) {
			return i
		}
	}

	return -1
}
