package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
)

// Statement represents an executable statement.
// Statements are built from expressions.
type Statement interface {

	// StatName returns the name of the statment.
	StatName() string

	// Evaluate evaluates the expression returning the resultant
	// value if there is one.
	Evaluate(*Context) fault.Fault
}
