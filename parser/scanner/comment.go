package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanComment scans comments. Comments can start anywhere and continue to the
// end of the line
func scanComment(r *Runer) (token.Token, fault.Fault) {

	// TODO: Needs testing

	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return token.EMPTY, err
		}

		if isNewline(ru) || ru == EOF {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	s := sb.String()
	size := len(s)

	tk := token.Token{
		Val:   s,
		Start: r.Col() - size + 1,
		End:   r.Col() + 1,
		Type:  token.COMMENT,
	}

	return tk, nil
}
