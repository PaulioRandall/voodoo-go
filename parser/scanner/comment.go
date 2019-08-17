package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanComment scans comments. Comments can start anywhere and continue to the
// end of the line
func scanComment(r *Runer) token.Token {
	start := r.Col() + 1
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			errs := readerFaultToStringArray(err)
			return errorToken(r, start, errs)
		}

		if isNewline(ru) || ru == EOF {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return commentToken(r, start, sb.String())
}

// commentToken creates a new comment token.
func commentToken(r *Runer, start int, val string) token.Token {
	return token.Token{
		Val:   val,
		Start: start,
		End:   r.Col() + 1,
		Type:  token.TT_COMMENT,
	}
}
