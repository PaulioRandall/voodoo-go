package word

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scan/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ScanWord scans word tokens returning a keyword or identifier.
func ScanWord(r *runer.Runer) (token.Token, err.ScanError) {
	start := r.NextCol()

	isWordLetter := func(ru, _ rune) (bool, error) {
		ok := unicode.IsLetter(ru) || unicode.IsDigit(ru) || ru == '_'
		return ok, nil
	}

	w, e := r.ReadWhile(isWordLetter)

	if e != nil {
		return nil, err.NewByRuner(r, e)
	}

	return wordToken(r, start, w), nil
}

// wordToken returns a new word Token.
func wordToken(r *runer.Runer, start int, w string) token.Token {
	return scantok.New(
		w,
		r.Line(),
		start,
		r.NextCol(),
		findWordKind(w),
	)
}

// findWordKind finds the kind of the word.
func findWordKind(w string) token.Kind {
	return token.TT_ID
}
