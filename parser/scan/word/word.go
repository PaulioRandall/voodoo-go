package word

import (
	"strings"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// ScanWord scans word tokens returning a keyword or identifier.
func ScanWord(r *runer.Runer) (token.Token, perror.Perror) {
	start := r.NextCol()

	isWordLetter := func(ru, _ rune) bool {
		ok := unicode.IsLetter(ru) || unicode.IsDigit(ru) || ru == '_'
		return ok
	}

	w, e := r.ReadWhile(isWordLetter)

	if e != nil {
		return nil, e
	}

	return wordToken(r, start, w), nil
}

// wordToken returns a new word Token.
func wordToken(r *runer.Runer, start int, w string) token.Token {
	return token.New(
		w,
		r.Line(),
		start,
		r.NextCol(),
		findWordKind(w),
	)
}

// findWordKind finds the kind of the word.
func findWordKind(w string) token.Kind {
	switch strings.ToLower(w) {
	case `true`, `false`:
		return token.TT_BOOL
	default:
		return token.TT_ID
	}

}
