package scanner

import (
	"strings"
	"unicode"
)

// startsWith returns true if the input array starts with
// the runes represented by the string.
func startsWith(in []rune, s string) bool {
	if len(in) < len(s) {
		return false
	}

	for i, r := range s {
		if in[i] != r {
			return false
		}
	}

	return true
}

// scanFrac iterates a rune array until a single fractional
// has been extracted returning the fractional slice followed
// by a slice of the remaining input.
func scanFrac(in []rune) (frac []rune, out []rune) {
	if len(in) < 1 || !isDecimalSeparator(in[0]) {
		out = in
		frac = []rune{}
		return
	}

	var tail []rune

	frac = in[0:1]
	tail, out = scanInt(in[1:])
	frac = append(frac, tail...)

	return
}

// scanInt iterates a rune array until a single integer has
// been extracted returning the integer slice followed by a
// slice of the remaining input.
func scanInt(in []rune) (num []rune, out []rune) {
	for i, r := range in {
		if isDigit(r) || isUnderscore(r) {
			num = append(num, r)
			continue
		}

		out = in[i:]
		return
	}

	out = []rune{}
	return
}

// scanWordStr iterates a rune array until a single word has
// been extracted returning the word followed by a slice
// of the remaining input or nil if the whole input represented
// a single word.
func scanWordStr(in []rune) (string, []rune) {
	sb := strings.Builder{}

	for i, r := range in {
		if isLetter(r) || isDigit(r) || isUnderscore(r) {
			sb.WriteRune(r)
			continue
		}

		return sb.String(), in[i:]
	}

	return sb.String(), []rune{}
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
