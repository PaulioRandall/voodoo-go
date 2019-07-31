package expression

import (
	"fmt"

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
func (c *Context) Assign(id string, new Value) (err fault.Fault) {
	old := c.vars[id]

	if old != nil && old.Type() != new.Type() {
		oldType := NameOfValueType(old.Type())
		newType := NameOfValueType(new.Type())

		err = EvalFault{
			ExprType: `Assignment`,
			Msgs: []string{
				fmt.Sprintf("Identifier `%s` stores values of type `%s`", id, oldType),
				fmt.Sprintf("Your trying assign a value of type `%s`", newType),
				`Implicit type casting is not allowed`,
			},
		}
		return
	}

	c.vars[id] = new
	return
}
