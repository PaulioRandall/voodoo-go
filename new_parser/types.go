package new_parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Token
type Token token.Token

// Fault
type Fault fault.Fault

// Expression represents an evaluatable expression.
// Expressions are often built from other expressions
// which will need to be evaluated first.
type Expression interface {

	// Name returns the name of the expression.
	Name() string
}

// Operation represents an operation expression such
// as addition, subtraction, etc.
type Operation struct {
	Token Token
	Left  *Expression
	Right *Expression
}

// Value represents an expression which simple evaluates
// to a literal value or identifier.
type Value struct {
	Token Token
}

// Join represents an expression which joins the results
// of two other expressions.
type Join struct {
	Token []Token
	Left  *Expression
	Right *Expression
}
