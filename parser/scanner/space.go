package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSpace scans tokens that start with a unicode whitespace property rune
// returning a token representing all whitespace between two non-whitespace
// tokens. Newlines are the exception as they have a token of there own.
func scanSpace(r *Runer) token.Token {
	sb := strings.Builder{}
	start := r.Col() + 1

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			sErr := readerFaultToStringArray(err)
			return errorToken(r, start, sErr)
		}

		if !isSpace(ru) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return spaceToken(r, start, sb.String())
}

// spaceToken creates a new number token.
func spaceToken(r *Runer, start int, val string) token.Token {
	return token.Token{
		Val:   val,
		Line:  r.Line(),
		Start: start,
		End:   r.Col() + 1,
		Type:  token.TT_SPACE,
	}
}
