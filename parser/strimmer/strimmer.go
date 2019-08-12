package strimmer

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Strim normalises an array of tokens ready for the the token parsing, this
// involves:
// -> Removing shebang token
// -> Removing whitespace tokens
// -> Removing comment tokens
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Removing newlines or converting them to end of statement tokens
// -> Converting all letters to lowercase (except literals)
func Strim(in chan token.Token, out chan token.Token) {
	defer close(out)

	var prevType token.TokenType
	var tk token.Token
	var keep bool

	for tk = range in {
		keep = true

		switch {
		case tk.Type == token.SHEBANG:
			keep = false
		case tk.Type == token.WHITESPACE:
			keep = false
		case tk.Type == token.COMMENT:
			keep = false
		case tk.Type == token.NEWLINE:
			tk, keep = whenNewline(tk, prevType)
		case tk.Type == token.LITERAL_STRING:
			tk.Val = trimQuotes(tk.Val)
		case tk.Type == token.LITERAL_NUMBER:
			tk.Val = stripUnderscores(tk.Val)
		case isAlphabeticType(tk.Type):
			tk.Val = strings.ToLower(tk.Val)
		}

		if keep {
			out <- tk
		}

		prevType = tk.Type
	}
}

// stripUnderscores removes redudant underscores from numbers.
func stripUnderscores(s string) string {
	return strings.ReplaceAll(s, `_`, ``)
}

// trimQuotes removes the leading and trailing quotes on string literals.
func trimQuotes(s string) string {
	return s[1 : len(s)-1]
}

// isAlphabeticType returns true if input token type is for
// tokens that may have alphabetic values.
func isAlphabeticType(t token.TokenType) bool {
	switch t {
	case token.KEYWORD_FUNC:
	case token.KEYWORD_LOOP:
	case token.KEYWORD_WHEN:
	case token.KEYWORD_DONE:
	case token.BOOLEAN_TRUE:
	case token.BOOLEAN_FALSE:
	case token.SPELL:
	default:
		return false
	}

	return true
}

// isMultiLineType returns true if the input type allows for the following type
// to be a newline without ending the statement.
func isMultiLineType(t token.TokenType) bool {
	switch t {
	case token.SHEBANG:
	case token.UNDEFINED:
	case token.VALUE_DELIM:
	case token.NEWLINE:
	case token.END_OF_STATEMENT:
	case token.PAREN_CURVY_OPEN:
	case token.PAREN_SQUARE_OPEN:
	default:
		return false
	}

	return true
}

// whenNewline handles newline tokens.
func whenNewline(tk token.Token, prevType token.TokenType) (token.Token, bool) {
	if isMultiLineType(prevType) {
		return tk, false
	}

	tk.Type = token.END_OF_STATEMENT
	return tk, true
}