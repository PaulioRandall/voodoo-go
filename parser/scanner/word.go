package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
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
		return token.TT_FUNC
	case `loop`:
		return token.TT_LOOP
	case `when`:
		return token.TT_WHEN
	case `done`:
		return token.TT_DONE
	case `true`:
		return token.TT_TRUE
	case `false`:
		return token.TT_FALSE
	default:
		return token.TT_ID
	}
}
