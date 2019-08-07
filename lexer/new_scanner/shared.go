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

// isStrStart returns true if the language considers the rune
// to be the start of a string literal.
func isStrStart(r rune) bool {
	return r == '"'
}

// isSpell returns true if the language considers the rune
// to be the start of a spell.
func isSpellStart(r rune) bool {
	return r == '@'
}

// isDecimalSeparator returns true if the language considers the
// rune to be a separator between the integer part of a number
// and the fractional part.
func isDecimalSeparator(r rune) bool {
	return r == '.'
}

// isSpace returns true if the language considers the rune
// to be whitespace.
func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

// isLetter returns true if the language considers the rune
// to be a letter.
func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}

// isDigit returns true if the language considers the rune
// to be a digit.
func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// isUnderscore returns true if the rune is an underscore.
func isUnderscore(r rune) bool {
	return r == '_'
}
