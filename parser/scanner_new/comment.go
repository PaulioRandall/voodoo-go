package scanner_new

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanComment scans comments. Comments can start anywhere and continue to the
// end of the line
func scanComment(r *Runer) (*token.Token, ParseToken) {
	start := r.NextCol()
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return runerErrorToken(r, err), nil
		}

		if isNewline(ru) || ru == EOF {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	tk := commentToken(r, start, sb.String())
	return ScanNext(r, tk)
}

// commentToken creates a new comment token.
func commentToken(r *Runer, start int, v string) *token.Token {
	return &token.Token{
		Val:   v,
		Line:  r.Line(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_COMMENT,
	}
}
