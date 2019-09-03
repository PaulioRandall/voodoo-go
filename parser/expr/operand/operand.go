package operand

import (
	"strconv"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
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

// Exe satisfies the Expr interface.
func (o operand) Exe(ctx.Context) (value.Value, perror.Perror) {
	v, e := strconv.ParseFloat(o.t.Text(), 64)
	if e != nil {
		return nil, o.badFormat()
	}
	return value.NewNumber(v), nil
}

// badFormat returns a new Perror for a badly formatted number.
func (o operand) badFormat() perror.Perror {
	return perror.New(
		o.t.Line(),
		o.t.Start(),
		[]string{
			"Could not parse number '" + o.t.Text() + "'",
		},
	)
}

// String satisfies the Expr interface.
func (o operand) String() string {
	return o.t.Text()
}
