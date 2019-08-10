package strimmer

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/token"
)

// Strim normalises an array of tokens ready for the
// the token parsing, this involves:
// -> Removing whitespace tokens
// -> Removing comment tokens
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Removing now redundant punctuation
// -> Converting all letters to lowercase (Except string literals)
func Strim(in chan token.Token, out chan token.Token) {
	defer close(out)

	prevType := token.UNDEFINED
	var ok bool

	for tk := range in {
		switch {
		case tk.Type == token.SHEBANG:
			prevType = tk.Type
			continue
		case tk.Type == token.WHITESPACE:
			prevType = tk.Type
			continue
		case tk.Type == token.COMMENT:
			prevType = tk.Type
			continue
		case tk.Type == token.NEWLINE:
			tk, ok = whenNewline(tk, prevType)
			prevType = tk.Type
			if !ok {
				continue
			}
		case tk.Type == token.LITERAL_STRING:
			penultimate := len(tk.Val) - 1
			tk.Val = tk.Val[1:penultimate]
		case tk.Type == token.LITERAL_NUMBER:
			tk.Val = strings.ReplaceAll(tk.Val, `_`, ``)
		case isAlphabeticType(tk.Type):
			tk.Val = strings.ToLower(tk.Val)
		}

		out <- tk
		prevType = tk.Type
	}
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
