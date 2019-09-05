package operand

import (
	"strconv"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Eval satisfies the Expr interface.
func (o operand) Eval(ctx.Context) (value.Value, perror.Perror) {
	if o.t.Kind() != token.TT_NUMBER {
		return nil, o.invalidKind()
	}

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

// invalidKind returns a new Perror for a when the token kind can not be
// handled by the evaluator.
func (o operand) invalidKind() perror.Perror {
	return perror.New(
		o.t.Line(),
		o.t.Start(),
		[]string{
			"Can not handle operands of kind '" + token.KindName(o.t.Kind()) + "'",
		},
	)
}
