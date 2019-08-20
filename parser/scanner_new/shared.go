package scanner_new

import (
	"strings"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// errorToken creates a new error token.
func errorToken(r *Runer, err []string) *token.Token {
	return &token.Token{
		Line:   r.Line(),
		End:    r.NextCol(),
		Type:   token.TT_ERROR_UPSTREAM,
		Errors: err,
	}
}

// runerErrorToken creates a new error token from an error returned by a Runer.
func runerErrorToken(r *Runer, err error) *token.Token {
	return &token.Token{
		Line:   r.Line(),
		Start:  -1,
		End:    r.NextCol(),
		Type:   token.TT_ERROR_UPSTREAM,
		Errors: []string{err.Error()},
	}
}

// scanWordStr reads a full word from a Runer.
func scanWordStr(r *Runer) (string, error) {
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return ``, err
		}

		if !isWordLetter(ru) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return sb.String(), nil
}

// isDecimalDelim returns true if the language considers the rune to be a
// separator between the integer part of a number and the fractional part.
func isDecimalDelim(ru rune) bool {
	return ru == '.'
}

// isNewline returns true if the language considers the rune to be a newline.
func isNewline(ru rune) bool {
	return ru == '\n'
}

// isSpace returns true if the language considers the rune to be whitespace.
// Note that newlines are not considered whitespace by this function, the
// isNewline() function should be used to check for newlines.
func isSpace(ru rune) bool {
	return !isNewline(ru) && unicode.IsSpace(ru)
}

// isLetter returns true if the language considers the rune to be a letter.
func isLetter(ru rune) bool {
	return unicode.IsLetter(ru)
}

// isNaturalDigit returns true if the digit is a positive interger.
func isNaturalDigit(ru rune) bool {
	return ru != 0 && unicode.IsDigit(ru)
}

// isDigit returns true if the language considers the rune to be a digit.
func isDigit(ru rune) bool {
	return unicode.IsDigit(ru)
}

// isUnderscore returns true if the rune is an underscore.
func isUnderscore(ru rune) bool {
	return ru == '_'
}

// isWordLetter returns true if the rune can be used within a word.
func isWordLetter(ru rune) bool {
	return isLetter(ru) || isDigit(ru) || isUnderscore(ru)
}

// isSpellPrefix returns true if the rune signifies a spell name coming.
func isSpellPrefix(ru rune) bool {
	return ru == '@'
}

// isStringPrefix returns true if the rune signifies a string literal body is
// coming.
func isStringPrefix(ru rune) bool {
	return ru == '"'
}

// isCommentPrefix returns true if the runes signify a comment coming.
func isCommentPrefix(ru1, ru2 rune) bool {
	return ru1 == '/' && ru2 == '/'
}
