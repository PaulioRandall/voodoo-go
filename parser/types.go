package parser

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
// Left and right list lengths must match otherwise invalid syntax.
type Assignment struct {
	Left     Expression // List of all identifiers on left side
	Operator Token
	Right    Expression // List created from joining results of all comma
	// separated right side expressions
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

// List represents an expression which simple evaluates
// to a list of literal values and identifiers.
type List struct {
	Tokens []Token
}

// Join represents an expression which joins the results
// of one or more other expressions into a list.
type Join struct {
	Exprs []Expression
}

// Operation represents an operation expression such
// as `addition`, `and`, `less than`, etc.
type Operation struct {
	Left     Expression // Could be any expression
	Operator Token
	Right    Expression // Value or func call
}
