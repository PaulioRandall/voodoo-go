package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanWord scans word tokens returning a keyword or identifier.
func scanWord(r *Runer) (token.Token, fault.Fault) {

	s, size, err := scanWordStr(r)
	if err != nil {
		return token.EMPTY, err
	}

	tk := token.Token{
		Val:   s,
		Start: r.Col() - size + 1,
		End:   r.Col() + 1,
		Type:  findWordType(s),
	}

	return tk, nil
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
