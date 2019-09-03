package space

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// ScanSpace scans space or newline runes into a space or newline Token.
func ScanSpace(r *runer.Runer) (token.Token, perror.Perror) {
	start := r.NextCol()

	isSpace := func(ru1, ru2 rune) bool {
		isNewline := ru1 == '\n' || (ru1 == '\r' && ru2 == '\n')
		return !isNewline && unicode.IsSpace(ru1)
	}

	s, e := r.ReadWhile(isSpace)
	if e != nil {
		return nil, e
	}

	if len(s) > 0 {
		return spaceToken(r, start, s), nil
	}

	return scanNewline(r, start)
}

// scanNewline scans a newline symbol set, i.e. `\n` or `\r\n`.
func scanNewline(r *runer.Runer, start int) (token.Token, perror.Perror) {
	ru1, _, e := r.Read()
	if e != nil {
		return nil, e
	}

	if ru1 == '\r' {
		ru2, _, e := r.Read()
		if e != nil {
			return nil, e
		}

		return newlineToken(r, start, string(ru1)+string(ru2)), nil
	}

	return newlineToken(r, start, string(ru1)), nil
}

// spaceToken returns a new space token.
func spaceToken(r *runer.Runer, start int, s string) token.Token {
	return token.New(
		s,
		r.Line(),
		start,
		r.NextCol(),
		token.TT_SPACE,
	)
}

// newlineToken returns a new newline token.
func newlineToken(r *runer.Runer, start int, s string) token.Token {
	return token.New(
		s,
		r.Line()-1,
		start,
		r.Col()+1,
		token.TT_NEWLINE,
	)
}
