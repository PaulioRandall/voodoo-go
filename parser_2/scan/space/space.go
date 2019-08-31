package space

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scan/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ScanSpace scans space or newline runes into a space or newline Token.
func ScanSpace(r *runer.Runer) (token.Token, err.ScanError) {
	start := r.NextCol()

	isSpace := func(ru1, ru2 rune) (bool, error) {
		isNewline := ru1 == '\n' || (ru1 == '\r' && ru2 == '\n')
		return !isNewline && unicode.IsSpace(ru1), nil
	}

	s, e := r.ReadWhile(isSpace)
	if e != nil {
		return nil, err.NewByRuner(r, e)
	}

	if len(s) > 0 {
		return spaceToken(r, start, s), nil
	}

	return scanNewline(r, start)
}

// scanNewline scans a newline symbol set, i.e. `\n` or `\r\n`.
func scanNewline(r *runer.Runer, start int) (token.Token, err.ScanError) {
	ru1, _, e := r.Read()
	if e != nil {
		return nil, err.NewByRuner(r, e)
	}

	if ru1 == '\r' {
		ru2, _, e := r.Read()
		if e != nil {
			return nil, err.NewByRuner(r, e)
		}

		return newlineToken(r, start, string(ru1)+string(ru2)), nil
	}

	return newlineToken(r, start, string(ru1)), nil
}

// spaceToken returns a new space token.
func spaceToken(r *runer.Runer, start int, s string) token.Token {
	return scantok.New(
		s,
		r.Line(),
		start,
		r.NextCol(),
		token.TT_SPACE,
	)
}

// newlineToken returns a new newline token.
func newlineToken(r *runer.Runer, start int, s string) token.Token {
	return scantok.New(
		s,
		r.Line()-1,
		start,
		r.Col()+1,
		token.TT_NEWLINE,
	)
}
