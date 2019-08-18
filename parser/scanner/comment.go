package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanComment scans comments. Comments can start anywhere and continue to the
// end of the line
func scanComment(r *Runer) token.Token {
	start := r.NextCol()
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return *runerErrorToken(r, err)
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
func commentToken(r *Runer, start int, v string) token.Token {
	return token.Token{
		Val:   v,
		Line:  r.Line(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_COMMENT,
	}
}
