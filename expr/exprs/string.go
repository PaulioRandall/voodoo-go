package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// String represents an expression which simple evaluates
// to a literal string.
type String struct {
	String token.Token
}

// Evaluate satisfies the Expression interface.
func (s String) Evaluate(c *ctx.Context) (ctx.Value, fault.Fault) {
	val := ctx.StringValue(s.String.Val)
	return val, nil
}
