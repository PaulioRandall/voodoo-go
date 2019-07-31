package expression

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// ValueOf represents an expression which simple evaluates to the
// value of an identifier.
type ValueOf struct {
	Identifier token.Token // Target identifier
}

// Evaluate satisfies the Expression interface.
func (vo *ValueOf) Evaluate(c *Context) (Value, fault.Fault) {
	return c.Get(vo.Identifier.Val)
}
