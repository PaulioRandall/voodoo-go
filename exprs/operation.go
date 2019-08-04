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

// ExprName satisfies the Expression interface.
func (o Operation) ExprName() string {
	return `operation`
}

// Evaluate satisfies the Expression interface.
func (o Operation) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	// TODO
	return ctx.Value{}, nil
}
