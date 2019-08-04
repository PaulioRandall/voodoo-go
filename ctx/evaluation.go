package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Evaluation represents an expression which performs some
// logic when it evaluates such as addition, calling a func, etc.
type Evaluation struct {
	Tokens []token.Token
	Type   ExprType
}

// Expr satisfies the Expression interface.
func (e Evaluation) Expr() ExprType {
	return e.Type
}

// Evaluate satisfies the Expression interface.
func (e Evaluation) Evaluate(c *Context) (Value, fault.Fault) {
	// TODO
	return Value{}, nil
}
