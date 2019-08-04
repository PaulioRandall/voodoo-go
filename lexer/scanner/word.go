package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/token"
)

// scanWord scans tokens that start with a unicode category L
// rune returning a keyword or identifier.
func scanWord(in []rune, col int) (*token.Token, []rune) {

	s, out := scanWordStr(in)

	tk := &token.Token{
		Val:   s,
		Start: col,
		Type:  findWordType(s),
	}

	return tk, out
}

// findWordType finds the type of the word.
func findWordType(s string) token.TokenType {
	switch strings.ToLower(s) {
	case `func`:
		return token.KEYWORD_FUNC
	case `loop`:
		return token.KEYWORD_LOOP
	case `when`:
		return token.KEYWORD_WHEN
	case `done`:
		return token.KEYWORD_DONE
	case `true`:
		return token.BOOLEAN_TRUE
	case `false`:
		return token.BOOLEAN_FALSE
	default:
		return token.IDENTIFIER
	}
}
