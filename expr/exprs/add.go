package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Add represents an expression which adds the result of
// evaluating two different expressions.
type Add struct {
	Add   token.Token // Token of the add symbol
	Left  ctx.Expression
	Right ctx.Expression
}

// Evaluate satisfies the Expression interface.
func (a Add) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	return nil, nil
}
