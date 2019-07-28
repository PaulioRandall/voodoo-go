package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanWord scans symbols that start with a unicode category L
// rune returning a keyword or identifier. Note that anything
// not a literal within voodoo is NOT case sensitive.
func scanWord(itr *runer.RuneItr) *token.Token {

	start := itr.Index()
	s := scanWordStr(itr)
	t := token.UNDEFINED

	switch strings.ToLower(s) {
	case `func`:
		t = token.KEYWORD_FUNC
	case `loop`:
		t = token.KEYWORD_LOOP
	case `when`:
		t = token.KEYWORD_WHEN
	case `end`:
		t = token.KEYWORD_END
	case `key`:
		t = token.KEYWORD_KEY
	case `val`:
		t = token.KEYWORD_VALUE
	case `true`:
		t = token.BOOLEAN_TRUE
	case `false`:
		t = token.BOOLEAN_FALSE
	default:
		t = token.IDENTIFIER_EXPLICIT
	}

	return &token.Token{
		Val:   s,
		Start: start,
		End:   itr.Index(),
		Type:  t,
	}
}

// scanWordStr iterates a rune iterator until a single word has
// been extracted retruning the string.
func scanWordStr(itr *runer.RuneItr) string {
	sb := strings.Builder{}

	for itr.HasNext() {
		switch {
		case itr.IsNextLetter():
			fallthrough
		case itr.IsNextDigit():
			fallthrough
		case itr.IsNext('_'):
			sb.WriteRune(itr.NextRune())
		default:
			return sb.String()
		}
	}

	return sb.String()
}
