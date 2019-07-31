package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Subtract represents an expression which subtracts the result
// of evaluating two different expressions. Right is subtracted
// from left.
type Subtract struct {
	Subtract token.Token // Token of the subtracts symbol
	Left     ctx.Expression
	Right    ctx.Expression
}

// Evaluate satisfies the Expression interface.
func (a Subtract) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	return nil, nil
}
