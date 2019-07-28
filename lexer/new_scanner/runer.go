package scanner

import (
	"strings"
	"unicode"
)

// scanInt iterates a rune array until a single integer
// has been extracted returning the integer followed by a
// slice of the remaining input or nil if the whole input
// represented a single integer.
func scanInt(in []rune) (string, []rune) {
	sb := strings.Builder{}

	for i, r := range in {
		if isDigit(r) || isUnderscore(r) {
			sb.WriteRune(r)
			continue
		}

		return sb.String(), in[i:]
	}

	return sb.String(), nil
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

	return sb.String(), nil
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
