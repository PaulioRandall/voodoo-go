package word

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ScanWord scans word tokens returning a keyword or identifier.
func ScanWord(r *runer.Runer) (token.Token, err.ScanError) {
	start := r.NextCol()

	w, e := r.ReadWhile(func(ru rune) (bool, error) {
		ok := unicode.IsLetter(ru) || unicode.IsDigit(ru) || ru == '_'
		return ok, nil
	})

	if e != nil {
		return nil, err.NewByRuner(r, e)
	}

	tk := scantok.New(
		w,
		r.Line(),
		start,
		r.NextCol(),
		findWordKind(w),
	)

	return tk, nil
}

// findWordKind finds the kind of the word.
func findWordKind(w string) token.Kind {
	return token.TT_ID
	/*
		switch strings.ToLower(w) {
		case `func`:
			return token.TT_FUNC
		case `loop`:
			return token.TT_LOOP
		case `match`:
			return token.TT_MATCH
		case `true`:
			return token.TT_TRUE
		case `false`:
			return token.TT_FALSE
		default:
			return token.TT_ID
		}
	*/
}
