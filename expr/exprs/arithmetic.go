package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Arithmetic represents an expression which performs
// one of the four basic arithmetic operations. Addition,
// subtraction, multiplication, or division and returns
// the result.
type Arithmetic struct {
	Operator token.Token    // Arithmetic token defined in the scroll
	Left     ctx.Expression // Left value or numerator
	Right    ctx.Expression // Right value or denominator
}

// Evaluate satisfies the Expression interface.
func (a Arithmetic) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	return nil, nil
}
