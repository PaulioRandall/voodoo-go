package exprs

import (
	"strconv"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Number represents an expression which simple evaluates
// to a literal number.
type Number struct {
	Number token.Token
}

// Evaluate satisfies the Expression interface.
func (n Number) Evaluate(c *ctx.Context) (val ctx.Value, err fault.Fault) {
	v, e := strconv.ParseFloat(n.Number.Val, 64)
	if e != nil {
		err = ctx.EvalFault{
			ExprType: `Assignment`,
			Msgs: []string{
				`Could not parse this literal number`,
				e.Error(),
			},
		}
		return
	}

	val = ctx.NumberValue(v)
	return
}
