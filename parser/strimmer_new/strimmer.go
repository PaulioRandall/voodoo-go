package strimmer_new

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
		case tk.Type == token.TT_ERROR_UPSTREAM:
			out <- tk
			return
		case tk.Type == token.TT_SHEBANG:
			keep = false
		case tk.Type == token.TT_SPACE:
			keep = false
		case tk.Type == token.TT_COMMENT:
			keep = false
		case tk.Type == token.TT_NEWLINE:
			tk, keep = whenNewline(tk, prevType)
		case tk.Type == token.TT_STRING:
			tk.Val = trimQuotes(tk.Val)
		case tk.Type == token.TT_NUMBER:
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
	case token.TT_FUNC:
	case token.TT_LOOP:
	case token.TT_WHEN:
	case token.TT_DONE:
	case token.TT_TRUE:
	case token.TT_FALSE:
	case token.TT_SPELL:
	default:
		return false
	}

	return true
}

// isMultiLineType returns true if the input type allows for the following type
// to be a newline without ending the statement.
func isMultiLineType(t token.TokenType) bool {
	switch t {
	case token.TT_SHEBANG:
	case token.TT_UNDEFINED:
	case token.TT_VALUE_DELIM:
	case token.TT_NEWLINE:
	case token.TT_EOS:
	case token.TT_CURVY_OPEN:
	case token.TT_SQUARE_OPEN:
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

	tk.Type = token.TT_EOS
	return tk, true
}
