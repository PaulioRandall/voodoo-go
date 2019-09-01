package stat

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// Statement is an ordered anbd executable set of tokens.
type Statement interface {

	// Kind returns the type of statement.
	Kind() Kind

	// Assign returns the array of tokens that represent the assignment target.
	// If nil then there is no assignment.
	Assign() []token.Token

	// Tokens returns the array of the tokens that represent the statement body.
	Tokens() []token.Token

	// String returns a string representation of the statment.
	String() string
}

// stat implement the Statement interface.
type stat struct {
	k Kind
	a []token.Token
	t []token.Token
}

// Empty creates a new empty statment.
func Empty() Statement {
	return stat{}
}

// New creates a new statment.
func New(k Kind, a, t []token.Token) Statement {
	return stat{k, a, t}
}

// Kind satisfies the Statement interface.
func (s stat) Kind() Kind {
	return s.k
}

// Assign satisfies the Statement interface.
func (s stat) Assign() []token.Token {
	return s.a
}

// Tokens satisfies the Statement interface.
func (s stat) Tokens() []token.Token {
	return s.t
}

// Tokens satisfies the Statement interface.
func (s stat) String() string {
	return `TODO: stat.String()`
}
