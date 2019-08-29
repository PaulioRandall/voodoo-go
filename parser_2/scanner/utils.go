package scanner

import (
	"strings"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
)

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
