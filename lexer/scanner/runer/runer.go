package scanner

import (
	"strings"
	"unicode"
)

// scanWordStr iterates a rune array until a single word has
// been extracted returning the word followed by a slice
// of the remaining input or nil if there's the whole input
// represented a single word.
func scanWordStr(in []rune) (string, []rune) {
	sb := strings.Builder{}

	for i, r := range in {
		if IsLetter(r) || IsDigit(r) || IsUnderscore(r) {
			sb.WriteRune(r)
			continue
		}

		return sb.String(), in[i:]
	}

	return sb.String(), nil
}

// IsLetter returns true if the language considers the rune
// to be a letter.
func IsLetter(r rune) bool {
	return unicode.IsLetter(r)
}

// IsDigit returns true if the language considers the rune
// to be a digit.
func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// IsUnderscore returns true if the rune is an underscore.
func IsUnderscore(r rune) bool {
	return r == '_'
}
