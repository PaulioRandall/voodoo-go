package parser

import (
	"github.com/PaulioRandall/voodoo-go/expression"
	"github.com/PaulioRandall/voodoo-go/fault"
)

// ParseTree represents a parse tree.
type ParseTree struct {
}

// Evaluate satisfies the Expression interface.
func (pt *ParseTree) Evaluate(c *expression.Context) (expression.Value, fault.Fault) {
	return nil, nil
}
