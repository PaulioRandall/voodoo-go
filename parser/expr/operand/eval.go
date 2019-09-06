package operand

import (
	"strconv"
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Eval satisfies the Expr interface.
func (o operand) Eval(c ctx.Context) (value.Value, perror.Perror) {
	switch o.t.Kind() {
	case token.TT_ID:
		return o.getId(c)
	case token.TT_BOOL:
		return o.parseBool()
	case token.TT_NUMBER:
		return o.parseNum()
	case token.TT_STRING:
		return o.parseStr()
	case token.TT_VOID:
		return nil, nil
	default:
		return nil, o.invalidKind()
	}
}

// parseStr parses a string token.
func (o operand) parseStr() (value.Value, perror.Perror) {
	return value.String(o.t.Text()), nil
}

// parseBool parses a boolean token.
func (o operand) parseBool() (value.Value, perror.Perror) {
	t := strings.ToLower(o.t.Text())
	v, e := strconv.ParseBool(t)
	if e != nil {
		return nil, o.badBoolFormat()
	}
	return value.Bool(v), nil
}

// parseNum parses a number token.
func (o operand) parseNum() (value.Value, perror.Perror) {
	v, e := strconv.ParseFloat(o.t.Text(), 64)
	if e != nil {
		return nil, o.badNumFormat()
	}
	return value.Number(v), nil
}

// getId returns the value identified by the operands token ID.
func (o operand) getId(c ctx.Context) (value.Value, perror.Perror) {
	v, ok := c.Vars[o.t.Text()]
	if !ok {
		return nil, o.unknownID()
	}
	return v, nil
}

// badBoolFormat returns a new Perror for a badly formatted number.
func (o operand) badBoolFormat() perror.Perror {
	return o.newPerror([]string{
		"Could not parse boolean '" + o.t.Text() + "'",
	})
}

// badNumFormat returns a new Perror for a badly formatted number.
func (o operand) badNumFormat() perror.Perror {
	return o.newPerror([]string{
		"Could not parse number '" + o.t.Text() + "'",
	})
}

// unknownID returns a new Perror for a when the token kind is an ID but the ID
// was not in the map of existing variables.
func (o operand) unknownID() perror.Perror {
	return o.newPerror([]string{
		"Unknown ID '" + o.t.Text() + "'",
	})
}

// invalidKind returns a new Perror for a when the token kind can not be
// handled by the evaluator.
func (o operand) invalidKind() perror.Perror {
	return o.newPerror([]string{
		"Can not handle operands of kind '" + token.KindName(o.t.Kind()) + "'",
	})
}

// newPerror creates a new Perror.
func (o operand) newPerror(m []string) perror.Perror {
	return perror.New(
		o.t.Line(),
		o.t.Start(),
		m,
	)
}
