package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Assignment represents an assignment of values from the result
// of each value of a set expressions to a set of identifiers.
// Left and right set lengths must match otherwise invalid syntax.
type Assignment struct {
	Left     Value // List of all identifiers on left side
	Operator token.Token
	Right    Expression // List created from joining results of all comma
	// separated right side expressions
}

// StatName satisfies the Statement interface.
func (a Assignment) StatName() string {
	return `assignment`
}

// Evaluate satisfies the Express interface.
func (a Assignment) Evaluate(c *Context) fault.Fault {
	// TODO
	return nil
}
