package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
)

// Expression represents an expression that results in a value.
type Expression interface {

	// ExprName returns the name of the expression type.
	ExprName() string

	// Evaluate evaluates the expression returning the resultant
	// value if there is one.
	Evaluate(*Context) (Value, fault.Fault)
}
