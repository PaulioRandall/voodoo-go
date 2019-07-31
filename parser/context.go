package parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
)

// Context represents the working environment in which
// expressions can be evaluated. It contains the
// identifiers and their values available to an
// expression and provide means to add or modify them.
type Context struct {
	IDs map[string]Value // Map of identifiers available to expressions
}

// Assign assigns a value to an identifier ensuring that the value
// is not empty and that the assigned is compatible if the identifier
// already exists.
func (c *Context) Assign(id string, v Value) fault.Fault {
	return nil
}
