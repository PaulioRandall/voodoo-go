package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
)

// ExprType represents a type of expression.
type ExprType int

const (
	UNDEFINED ExprType = iota
	ASSIGNMENT
	ADDITION
)

// Expression represents an expression that results in a value.
type Expression interface {

	// ExprType returns the type of the expression.
	ExprType() ExprType

	// Evaluate evaluates the expression returning the resultant
	// value if there is one.
	Evaluate(*Context) (Value, fault.Fault)
}
