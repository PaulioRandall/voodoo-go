package operand

import (
	"strconv"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Eval satisfies the Expr interface.
func (o operand) Eval(ctx.Context) (value.Value, perror.Perror) {
	v, e := strconv.ParseFloat(o.t.Text(), 64)
	if e != nil {
		return nil, o.badFormat()
	}
	return value.Number(v), nil
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
