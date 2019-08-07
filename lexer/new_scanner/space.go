package new_scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSpace scans tokens that start with a unicode whitespace property rune
// returning a token representing all whitespace between two non-whitespace
// tokens. Newlines are the exception as they have a token of there own.
func scanSpace(r *Runer) (token.Token, fault.Fault) {
	sb := strings.Builder{}
	start := r.Col() + 1

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return token.EMPTY, err
		}

		if !isSpace(ru) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	tk := token.Token{
		Val:   sb.String(),
		Start: start,
		End:   r.Col() + 1,
		Type:  token.WHITESPACE,
	}

	return tk, nil
}
