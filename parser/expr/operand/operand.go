package operand

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// operand represents an operand expression.
type operand struct {
	t token.Token
}

// New returns a new operand expression.
func New(t token.Token) expr.Expr {
	return operand{t}
}

// Token satisfies the Expr interface.
func (o operand) Token() token.Token {
	return o.t
}

// String satisfies the Expr interface.
func (o operand) String() string {
	return o.t.Text()
}
