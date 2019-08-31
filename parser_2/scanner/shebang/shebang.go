package shebang

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ScanShebang scans a shebang line.
func ScanShebang(r *runer.Runer) (token.Token, err.ScanError) {
	start := r.NextCol()

	s, e := r.ReadWhile(func(ru, _ rune) (bool, error) {
		return ru != '\n', nil
	})

	if e != nil {
		return nil, err.NewByRuner(r, e)
	}

	return newShebangToken(r, s, start), nil
}

// newShebangToken returns a new SHEBANG token.
func newShebangToken(r *runer.Runer, text string, start int) token.Token {
	return scantok.New(
		text,
		r.Line(),
		start,
		r.NextCol(),
		token.TT_SHEBANG,
	)
}
