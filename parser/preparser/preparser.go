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

	tk = s.strim(tk)
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

// IsEmpty returns true if the statement is empty.
func (s *Statement) IsEmpty() bool {
	return len(s.Tokens) < 1
}

// Len returns the number of tokens within the statement.
func (s *Statement) Len() int {
	return len(s.Tokens)
}

// Complete sets the complete flag to true preventing any more tokens being
// added and flagging the statement as ready to parse.
func (s *Statement) Complete() {
	s.complete = true
}

// check checks the statement is not complete, if it is an panic is generated
// using the supplied error message.
func (s *Statement) check(m string) {
	if s.complete {
		panic(m)
	}
}

// strim normalises a token. This may involve removing the token or modifying it
// ready for parsing. Sometimes an extra token needs to be inserted before or
// after the normal one.
func (s *Statement) strim(tk *token.Token) *token.Token {

	t := tk.Type

	switch {
	case t == token.TT_SHEBANG:
		tk = nil
	case t == token.TT_SPACE:
		tk = nil
	case t == token.TT_COMMENT:
		tk = nil
	case t == token.TT_NEWLINE:
		tk = whenNewline(tk, s.prevType)
	case t == token.TT_STRING:
		tk = trimQuotes(tk)
	case isAlphabeticType(t):
		tk = toLower(tk)
	}

	s.prevType = t
	return tk
}

// whenNewline handles newline tokens.
func whenNewline(tk *token.Token, prevType token.TokenType) *token.Token {
	if isMultiLineType(prevType) {
		return nil
	}

	tk.Type = token.TT_EOS
	return tk
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
	case token.TT_CURVED_OPEN:
	case token.TT_SQUARE_OPEN:
	default:
		return false
	}

	return true
}

// trimQuotes removes the leading and trailing quotes on string literals.
func trimQuotes(tk *token.Token) *token.Token {
	tk.Val = tk.Val[1 : len(tk.Val)-1]
	return tk
}

// isAlphabeticType returns true if input token type is for
// tokens that may have alphabetic values.
func isAlphabeticType(t token.TokenType) bool {
	switch t {
	case token.TT_FUNC:
	case token.TT_LOOP:
	case token.TT_MATCH:
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
func toLower(tk *token.Token) *token.Token {
	tk.Val = strings.ToLower(tk.Val)
	return tk
}
