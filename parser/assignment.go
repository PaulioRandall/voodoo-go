package parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Assignment represents an assignment expression.
type Assignment struct {
	Operator   token.Token // Scroll token representing the operator used
	Identifier token.Token // Identifier that will reference the result
	Expression Expression  // Expression to evaluate that produces the result
}

// Evaluate satisfies the Expression interface.
func (a *Assignment) Evaluate(c *Context) (v Value, err fault.Fault) {
	val, err := a.Expression.Evaluate(c)
	if err != nil {
		return
	}

	id := a.Identifier.Val
	c.IDs[id] = val
	return
}
