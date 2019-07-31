package expression

import (
	"strconv"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// ValueOf represents an expression which simple evaluates to the
// value of an identifier.
type ValueOf struct {
	Identifier token.Token // Target identifier
}

// Evaluate satisfies the Expression interface.
func (vo ValueOf) Evaluate(c *Context) (Value, fault.Fault) {
	return c.Get(vo.Identifier.Val)
}

// Number represents an expression which simple evaluates
// to a literal number.
type Number struct {
	Number token.Token // Target identifier
}

// Evaluate satisfies the Expression interface.
func (n Number) Evaluate(c *Context) (val Value, err fault.Fault) {
	v, e := strconv.ParseFloat(n.Number.Val, 64)
	if e != nil {
		err = EvalFault{
			ExprType: `Assignment`,
			Msgs: []string{
				`Could not parse this literal number`,
				e.Error(),
			},
		}
		return
	}

	val = NumberValue(v)
	return
}
