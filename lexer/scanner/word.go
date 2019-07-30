package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/token"
)

// scanWord scans tokens that start with a unicode category L
// rune returning a keyword or identifier.
func scanWord(in []rune, col int) (*token.Token, []rune) {

	s, out := scanWordStr(in)
	t := token.IDENTIFIER_EXPLICIT

	switch strings.ToLower(s) {
	case `func`:
		t = token.KEYWORD_FUNC
	case `loop`:
		t = token.KEYWORD_LOOP
	case `when`:
		t = token.KEYWORD_WHEN
	case `done`:
		t = token.KEYWORD_DONE
	case `true`:
		t = token.BOOLEAN_TRUE
	case `false`:
		t = token.BOOLEAN_FALSE
	}

	tk := &token.Token{
		Val:   s,
		Start: col,
		Type:  t,
	}

	return tk, out
}
