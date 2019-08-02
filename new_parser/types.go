package new_parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Token
type Token token.Token

// Fault
type Fault fault.Fault

// Statement represents an executable statement.
// Statements are built from expressions.
type Statement interface {

	// StatName returns the name of the statment.
	StatName() string
}

// Assignment represents an assignment statement.
type Assignment struct {
	Token Token
	Left  *Join
	Right *Join
}

// Expression represents an evaluatable expression.
// Expressions are often built from other expressions.
type Expression interface {

	// ExprName returns the name of the expression.
	ExprName() string
}

// Value represents an expression which simple evaluates
// to a literal value or identifier.
type Value struct {
	Token Token
}

// Join represents an expression which joins the results
// of two other expressions.
type Join struct {
	Left  *Expression
	Right *Expression
}

// Operation represents an operation expression such
// as addition, subtraction, etc.
type Operation struct {
	Token Token
	Left  *Expression
	Right *Expression
}
