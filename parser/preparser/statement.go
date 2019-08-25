package preparser

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"strconv"
)

// Statement represents a statement of strimmed tokens ready for parsing.
type Statement struct {
	tokens   []token.Token // Tokens
	complete bool          // completelete
}

// NewStatement creates a new statement.
func NewStatement() *Statement {
	return &Statement{
		tokens: []token.Token{},
	}
}

// IsEmpty returns true if the statement is empty.
func (s *Statement) IsEmpty() bool {
	return len(s.tokens) < 1
}

// Len returns the number of tokens within the statement.
func (s *Statement) Len() int {
	return len(s.tokens)
}

// Iscompletelete returns the completeletion flag.
func (s *Statement) IsComplete() bool {
	return s.complete
}

// Setcompletelete sets the completelete flag to true preventing any more tokens being
// added and flagging the statement as ready to parse.
func (s *Statement) SetComplete() {
	s.complete = true
}

// Append appends a token to the statement. Attempting to append after the
// completelete flag has been set will result in a panic.
func (s *Statement) Append(tk token.Token) {
	if s.complete {
		panic(`Can't add more tokens to a completelete statement`)
	}
	s.tokens = append(s.tokens, tk)
}

// Tokens returns a completelete statement.
func (s *Statement) Tokens() []token.Token {
	if !s.complete {
		panic(`A statement must be completelete before it can be built`)
	}
	return s.tokens
}

// Reset clears the tokens and the completelete flag from the statement.
func (s *Statement) Reset() {
	s.tokens = []token.Token{}
	s.complete = false
}

// token returns the token at the specified index. Used for testing only.
func (s *Statement) token(i int) *token.Token {
	if i >= s.Len() {
		m := "Token index out of bounds; "
		m += "(index)" + strconv.Itoa(i) + " >= "
		m += strconv.Itoa(s.Len()) + "(len(tokens))"
		panic(m)
	}
	return &s.tokens[i]
}
