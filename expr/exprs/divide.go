package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Divide represents an expression which divides the result
// of evaluating two different expressions. Left is the
// numerator, right is the denominator.
type Divide struct {
	Divide token.Token // Token of the divide symbol
	Left   ctx.Expression
	Right  ctx.Expression
}

// Evaluate satisfies the Expression interface.
func (a Divide) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	return nil, nil
}
