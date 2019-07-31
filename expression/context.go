package expression

import (
	"github.com/PaulioRandall/voodoo-go/fault"
)

// Expression represents a node within a parse tree.
type Expression interface {

	// Evaluate executes the expression within the given
	// context.
	Evaluate(*Context) (Value, fault.Fault)
}

// Context represents the working environment in which
// expressions can be evaluated. It contains the existing
// identifiers and their values and provides a means to
// add and modify them.
type Context struct {
	vars map[string]Value // Map of identifiers to values
}

// Assign assigns a value to an identifier ensuring that the
// currently assigned value is compatible with the identifiers
// type, if it already exists. An assumption is made that the
// input value is never nil.
func (c *Context) Assign(id string, new Value) fault.Fault {
	old := c.vars[id]
	if old == nil {
		c.vars[id] = new
		return nil
	}

	// TODO: When the value already exists

	return nil
}
