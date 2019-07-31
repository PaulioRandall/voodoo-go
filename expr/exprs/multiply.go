package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Multiply represents an expression which multiplies the result
// of evaluating two different expressions.
type Multiply struct {
	Multiply token.Token // Token of the multiply symbol
	Left     ctx.Expression
	Right    ctx.Expression
}

// Evaluate satisfies the Expression interface.
func (a Multiply) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	return nil, nil
}
