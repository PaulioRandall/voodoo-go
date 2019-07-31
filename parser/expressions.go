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
func (pt *Assignment) Evaluate(c *Context) fault.Fault {
	return nil
}
