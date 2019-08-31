package stat

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// Statement is an ordered anbd executable set of tokens.
type Statement interface {

	// Tokens returns the array of the tokens that represent the statement.
	Tokens() []token.Token

	// Kind returns the type of statement.
	Kind() Kind

	// String returns a string representation of the statment.
	String() string
}

// stat implement the Statement interface.
type stat struct {
	t []token.Token
	k Kind
}

// Tokens satisfies the Statement interface.
func (f stat) Tokens() []token.Token {
	return f.t
}

// Tokens satisfies the Statement interface.
func (f stat) Kind() Kind {
	return f.k
}

// Tokens satisfies the Statement interface.
func (f stat) String() string {
	return `TODO: farmStat.String()`
}
