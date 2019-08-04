package exprs

import (
	"github.com/PaulioRandall/voodoo-go/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
)

// Join represents an expression which simple joins the
// results of all held expressions into a list.
type Join struct {
	Exprs []ctx.Expression
}

// ExprName satisfies the Expression interface.
func (j Join) ExprName() string {
	return `join`
}

// Evaluate satisfies the Expression interface.
func (j Join) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {

	size := len(j.Exprs)
	vals := make([]ctx.Value, size)

	for i, e := range j.Exprs {
		v, err := e.Evaluate(c)
		if err != nil {
			return ctx.Value{}, err
		}
		vals[i] = v
	}

	out := ctx.Value{
		Val:  vals,
		Type: ctx.LIST_TYPE,
	}

	return out, nil
}
