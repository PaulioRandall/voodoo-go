package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSpace scans tokens that start with a unicode whitespace property rune
// returning a token representing all whitespace between two non-whitespace
// tokens. Line feeds and carriage returns that preceed a line feed are an
// exempt; they have a token of their own.
func scanSpace(r *Runer) (*token.Token, ParseToken, *token.Token) {
	sb := strings.Builder{}
	start := r.NextCol()

	for {
		ru1, ru2, err := r.LookAhead()
		if err != nil {
			return nil, nil, runerErrorToken(r, err)
		}

		if isCarriageReturn(ru1) && isNewline(ru2) {
			break
		}

		if !isSpace(ru1) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru1)
	}

	tk := spaceToken(r, start, sb.String())
	return scanNext(r, tk)
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
