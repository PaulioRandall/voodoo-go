package preparser

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Statement represents a statement of strimmed tokens ready for parsing.
type Statement struct {
	Tokens   []token.Token
	prevType token.TokenType
	complete bool
}

// New creates a new statement.
func New() *Statement {
	return &Statement{
		Tokens:   []token.Token{},
		prevType: token.TT_UNDEFINED,
	}
}

// Add strims a token. If the token is not removed by the strimmer it is
// appended to the statement. If the token represents the end of the statement
// then the complete flag will be set to true and true will be returned.
// Attempting to add after the complete flag has been set will result in a
// panic.
func (s *Statement) Add(tk *token.Token) bool {
	s.check(`Can't add more tokens to a complete statement`)

	tk = s.callStrim(tk)
	if tk == nil {
		return s.complete
	}

	if tk.Type == token.TT_EOS {
		s.complete = true
	} else {
		s.Tokens = append(s.Tokens, *tk)
	}

	return s.complete
}

// callStrim calls the Strim function while keeping a track of the previous
// token types within the statement structure.
func (s *Statement) callStrim(tk *token.Token) *token.Token {
	t := tk.Type
	tk = Strim(*tk, s.prevType)
	s.prevType = t
	return tk
}

// check checks the statement is not complete, if it is an panic is generated
// using the supplied error message.
func (s *Statement) check(m string) {
	if s.complete {
		panic(m)
	}
}

// Strim normalises a token. This may involve removing the token or modifying it
// ready for parsing.
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
