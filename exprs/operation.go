package exprs

import (
	"github.com/PaulioRandall/voodoo-go/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Operation represents an operation expression such
// as `addition`, `and`, `less than`, etc.
type Operation struct {
	Left     ctx.Expression
	Operator token.Token
	Right    ctx.Expression
}

// ExprType satisfies the Expression interface.
func (o Operation) ExprType() ctx.ExprType {
	return ctx.ADDITION
}

// Evaluate satisfies the Expression interface.
func (o Operation) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	// TODO
	return ctx.Value{}, nil
}
