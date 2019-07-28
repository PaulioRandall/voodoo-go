package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// scanWord scans symbols that start with a unicode category L
// rune returning a keyword or identifier. Note that anything
// not a literal within voodoo is NOT case sensitive.
func scanWord(itr *runer.RuneItr) *symbol.Lexeme {

	start := itr.Index()
	s := scanWordStr(itr)
	t := symbol.UNDEFINED

	switch strings.ToLower(s) {
	case `func`:
		t = symbol.KEYWORD_FUNC
	case `loop`:
		t = symbol.KEYWORD_LOOP
	case `when`:
		t = symbol.KEYWORD_WHEN
	case `end`:
		t = symbol.KEYWORD_END
	case `key`:
		t = symbol.KEYWORD_KEY
	case `val`:
		t = symbol.KEYWORD_VALUE
	case `true`:
		t = symbol.BOOLEAN_TRUE
	case `false`:
		t = symbol.BOOLEAN_FALSE
	default:
		t = symbol.IDENTIFIER_EXPLICIT
	}

	return &symbol.Lexeme{
		Val:   s,
		Start: start,
		End:   itr.Index(),
		Type:  t,
	}
}
