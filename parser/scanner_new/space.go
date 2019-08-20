package scanner_new

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSpace scans tokens that start with a unicode whitespace property rune
// returning a token representing all whitespace between two non-whitespace
// tokens. Newlines are an exception here and considered as a non-whitespace
// character; they have a token of their own.
func scanSpace(r *Runer) (*token.Token, ParseToken, *token.Token) {
	sb := strings.Builder{}
	start := r.NextCol()

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return nil, nil, runerErrorToken(r, err)
		}

		if !isSpace(ru) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	tk := spaceToken(r, start, sb.String())
	return ScanNext(r, tk)
}

// spaceToken creates a new space token.
func spaceToken(r *Runer, start int, v string) *token.Token {
	return &token.Token{
		Val:   v,
		Line:  r.Line(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_SPACE,
	}
}
