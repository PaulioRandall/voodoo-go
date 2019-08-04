package ctx

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/fault"
)

// Context represents the working environment in which
// expressions can be evaluated. It contains the existing
// identifiers and their values and provides a means to
// add and modify them.
type Context struct {
	vars map[string]Value // Map of identifiers to actual values
}

// New returns a new initialised context.
func New(m map[string]Value) *Context {
	return &Context{
		vars: m,
	}
}

// Empty returns a new initialised context with an empty
// variable map.
func Empty() *Context {
	return &Context{
		vars: map[string]Value{},
	}
}

// Assign assigns a value to an identifier ensuring that the
// currently assigned value is compatible with the identifiers
// type, if it already exists. An assumption is made that the
// input value is never nil.
func (c *Context) Assign(id string, new Value) fault.Fault {
	old, haveEntry := c.vars[id]

	if !haveEntry || old.IsSameType(new) {
		c.vars[id] = new
		return nil
	}

	return EvalFault{
		Msgs: []string{
			fmt.Sprintf("Identifier `%s` stores values of type `%v`", id, old.Type),
			fmt.Sprintf("Your trying assign a value of type `%s`", new.Type),
			`Implicit type casting is not allowed`,
		},
	}
}

// Get returns the value of the specified identifier. If the identifier
// does not exist a fault is returned.
func (c *Context) Get(id string) (Value, fault.Fault) {

	v, ok := c.vars[id]

	if ok {
		return v, nil
	}

	return Value{}, EvalFault{
		Msgs: []string{
			fmt.Sprintf("Variable `%s` doesn't exist", id),
			`Or at least not within the current context`,
		},
	}
}
