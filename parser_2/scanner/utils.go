package scanner

import (
	"strings"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ParseToken represents a function produces a single specific type of Token
// from an input stream. Only the first part of the stream that provides longest
// match against the production rules of the specific Token will be read. If an
// error occurs then an error token will be returned instead.
//
// The first returned token alwyas represents a valid parsed token while the
// last always represents an error. On return, one should be nil and the other
// non-nil.
type ParseToken func(*runer.Runer) (token.Token, ScanError)

// runerError creates a new ScanError from an error returned by a Runer.
func runerError(r *runer.Runer, err error) ScanError {
	return &scanErr{
		l: r.Line(),
		i: r.NextCol(),
		e: []string{err.Error()},
	}
}

// scanWordStr reads a full word from a Runer.
func scanWordStr(r *runer.Runer) (string, ScanError) {
	sb := strings.Builder{}

	for {
		ru, eof, err := r.Peek()
		if err != nil {
			return ``, runerError(r, err)
		}

		if eof || !isWordRune(ru) {
			break
		}

		_, err = r.Skip()
		if err != nil {
			return ``, runerError(r, err)
		}

		sb.WriteRune(ru)
	}

	return sb.String(), nil
}

// isWordRune returns true if the rune is a valid member of a word string.
func isWordRune(ru rune) bool {
	return unicode.IsLetter(ru) || unicode.IsDigit(ru) || ru == '_'
}
