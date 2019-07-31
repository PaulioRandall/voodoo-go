package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Bool represents an expression which simple evaluates
// to a literal boolean.
type Bool struct {
	Bool token.Token
}

// Evaluate satisfies the Expression interface.
func (b Bool) Evaluate(c *ctx.Context) (val ctx.Value, err fault.Fault) {

	switch b.Bool.Val {
	case `true`:
		val = ctx.BoolValue(true)
	case `false`:
		val = ctx.BoolValue(false)
	default:
		err = ctx.EvalFault{
			ExprType: `Bool`,
			Msgs: []string{
				`Could not parse this literal boolean`,
				"Expected wither `true` or `false`",
			},
		}
	}

	return
}
