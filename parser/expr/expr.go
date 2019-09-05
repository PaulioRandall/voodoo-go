package expr

import (
	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Expr represents an expression.
type Expr interface {

	// Token returns the token that represents the operand or operation within
	// the scroll.
	Token() token.Token

	// Eval evaluates the expression within the given context.
	Eval(ctx.Context) (value.Value, perror.Perror)

	// String returns a string representation of the expression.
	String() string
}
