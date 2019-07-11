
package cleaver

import (
	"unicode"
)

// runeType represents a type of a rune.
type runeType int

const (
	none runeType = iota
	letter
	number
	other
)

// runeTypeOf returns the runeType of the input rune.
func runeTypeOf(r rune) runeType {
	if isLetter(r) {
		return letter
	}
	
	if isNumber(r) {
		return number
	}
	
	return other
}

// isLetter returns true if the rune is a letter.
func isLetter(r rune) bool {
	// (UTF-8) 65-90 == A-Z, 97-122 == a-z
	return (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
}

// isNumber returns true if the rune is a number.
func isNumber(r rune) bool {
	// (UTF-8) 48-57 == 0-9
	return r >= 48 && r <= 57
}

// isDoubleQuote returns true if the rune is a '"'.
func isDoubleQuote(r rune) bool {
	// (UTF-8) 34 == "
	return r == 34
}

// isUnderscore returns true if the rune is an '_'.
func isUnderscore(r rune) bool {
	// (UTF-8) 95 == _
	return r == 95
}

// isAtSign returns true if the rune is an '@'.
func isAtSign(r rune) bool {
	// (UTF-8) 64 == @
	return r == 64
}

// isSpace returns true if the rune is an ' '.
func isSpace(r rune) bool {
	// (UTF-8) 32 == ` `
	return r == 32
}

// isComma returns true if the rune is an ','.
func isComma(r rune) bool {
	// (UTF-8) 44 == ,
	return r == 44
}

// isPoint returns true if the rune is an '.'.
func isPoint(r rune) bool {
	// (UTF-8) 46 == .
	return r == 46
}

// isWhitespace returns true if the rune is whitespace.
func isWhitespace(r rune) bool {
	// TODO: Correct this when internet is back up
	return unicode.IsWhitespace(r)
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

// isNumStr returns true if the string only contains
// number runes.
func isNumStr(s string) bool {
	for _, v := range s {
		if !isNumber(v) {
			return false
		}
	}
	return true
}

// isLetterStr returns true if the string only contains
// letter runes.
func isLetterStr(s string) bool {
	for _, v := range s {
		if !isLetter(v) {
			return false
		}
	}
	return true
}
