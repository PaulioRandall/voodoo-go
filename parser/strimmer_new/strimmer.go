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
func Strim(in token.Token, prevType token.TokenType) *token.Token {

	var out *token.Token
	t := in.Type

	switch {
	case t == token.TT_SHEBANG:
		out = nil
	case t == token.TT_SPACE:
		out = nil
	case t == token.TT_COMMENT:
		out = nil
	case t == token.TT_NEWLINE:
		out = whenNewline(in, prevType)
	case t == token.TT_STRING:
		out = trimQuotes(in)
	case t == token.TT_NUMBER:
		out = stripUnderscores(in)
	case isAlphabeticType(t):
		out = toLower(in)
	default:
		out = &in
	}

	return out
}

// whenNewline handles newline tokens.
func whenNewline(tk token.Token, prevType token.TokenType) *token.Token {
	if isMultiLineType(prevType) {
		return nil
	}

	tk.Type = token.TT_EOS
	return &tk
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

// trimQuotes removes the leading and trailing quotes on string literals.
func trimQuotes(tk token.Token) *token.Token {
	tk.Val = tk.Val[1 : len(tk.Val)-1]
	return &tk
}

// stripUnderscores removes redudant underscores from numbers.
func stripUnderscores(tk token.Token) *token.Token {
	tk.Val = strings.ReplaceAll(tk.Val, `_`, ``)
	return &tk
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

// toLower returns the input token but with all the characters in the value
// field converted to lowercase.
func toLower(tk token.Token) *token.Token {
	tk.Val = strings.ToLower(tk.Val)
	return &tk
}
