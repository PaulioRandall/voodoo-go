
package lexer

import (
	"strings"
	"unicode"
)

// unicodeCat represents a unicode category.
type unicodeCat int

const (
	none unicodeCat = iota	// Not assigned
	letter										// L
	digit										// Nd
	other										// Anything not defined
)

// unicodeCatOf returns the unicodeCat of the input rune.
func unicodeCatOf(r rune) unicodeCat {
	if unicode.IsLetter(r) {
		return letter
	}
	
	if unicode.IsDigit(r) {
		return digit
	}
	
	return other
}

// isComment returns true if the passed string begins
// with a `//`.
func isComment(s string) bool {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, `//`) {
		return true
	}
	return false
}

// isDigitStr returns true if the string only contains
// decimal digit runes.
func isDigitStr(s string) bool {
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

// isLetterStr returns true if the string only contains
// letter runes.
func isLetterStr(s string) bool {
	for _, v := range s {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}
