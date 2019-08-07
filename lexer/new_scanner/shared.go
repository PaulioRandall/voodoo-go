package new_scanner

import (
	"strings"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/fault"
)

// scanWordStr reads a full word from a Runer.
func scanWordStr(r *Runer) (string, int, fault.Fault) {
	sb := strings.Builder{}
	size := 0

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return ``, -1, err
		}

		if !isLetter(ru) && !isDigit(ru) && !isUnderscore(ru) {
			break
		}

		r.SkipRune()
		size++
		sb.WriteRune(ru)
	}

	return sb.String(), size, nil
}

// isDecimalDelim returns true if the language considers the rune to be a
// separator between the integer part of a number and the fractional part.
func isDecimalDelim(r rune) bool {
	return r == '.'
}

// isNewline returns true if the language considers the rune to be a newline.
func isNewline(r rune) bool {
	return r == '\n'
}

// isSpace returns true if the language considers the rune to be whitespace.
// Note that newlines are not considered whitespace by this function, the
// isNewline() function should be used to check for newlines.
func isSpace(r rune) bool {
	return !isNewline(r) && unicode.IsSpace(r)
}

// isLetter returns true if the language considers the rune to be a letter.
func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}

// isDigit returns true if the language considers the rune to be a digit.
func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// isUnderscore returns true if the rune is an underscore.
func isUnderscore(r rune) bool {
	return r == '_'
}
