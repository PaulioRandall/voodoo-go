package shebang

import (
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// ScanShebang scans a shebang line.
func ScanShebang(r *runer.Runer) (token.Token, perror.Perror) {
	start := r.NextCol()

	s, e := r.ReadWhile(func(ru, _ rune) bool {
		return ru != '\n'
	})

	if e != nil {
		return nil, e
	}

	return newShebangToken(r, s, start), nil
}

// newShebangToken returns a new SHEBANG token.
func newShebangToken(r *runer.Runer, text string, start int) token.Token {
	return token.New(
		text,
		r.Line(),
		start,
		r.NextCol(),
		token.TT_SHEBANG,
	)
}
