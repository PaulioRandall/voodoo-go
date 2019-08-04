package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Operation represents an operation expression such
// as `addition`, `and`, `less than`, etc.
type Operation struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

// ExprType satisfies the Expression interface.
func (o Operation) ExprType() ExprType {
	return ADDITION
}

// Evaluate satisfies the Expression interface.
func (o Operation) Evaluate(c *Context) (Value, fault.Fault) {
	// TODO
	return Value{}, nil
}
