package parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
)

// ParseTree represents a parse tree.
type ParseTree struct {
}

// Evaluate satisfies the Expression interface.
func (pt *ParseTree) Evaluate(c *Context) (Value, fault.Fault) {
	return Value{}, nil
}
