package operand

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tok(t string, k token.Kind) token.Token {
	return token.New(t, 0, 0, 0, k)
}

func TestOperand_Eval_1(t *testing.T) {
	o := operand{
		t: tok(`1`, token.TK_NUMBER),
	}

	c := ctx.New(nil)
	exp := value.Number(1)

	act, e := o.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}

func TestOperand_Eval_2(t *testing.T) {
	o := operand{
		t: tok(`3.14159`, token.TK_NUMBER),
	}

	c := ctx.New(nil)
	exp := value.Number(3.14159)

	act, e := o.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}

func TestOperand_Eval_3(t *testing.T) {
	o := operand{
		t: tok(` `, token.TK_SPACE),
	}

	c := ctx.New(nil)

	act, e := o.Eval(c)
	require.NotNil(t, e)
	assert.Nil(t, act)
}

func TestOperand_Eval_4(t *testing.T) {
	o := operand{
		t: tok(`x`, token.TK_ID),
	}

	c := ctx.New(nil)
	c.Vars[`x`] = value.Number(2)
	exp := value.Number(2)

	act, e := o.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}

func TestOperand_Eval_5(t *testing.T) {
	o := operand{
		t: tok(`x`, token.TK_ID),
	}

	c := ctx.New(nil)

	act, e := o.Eval(c)
	require.NotNil(t, e)
	assert.Nil(t, act)
}

func TestOperand_Eval_6(t *testing.T) {
	o := operand{
		t: tok(`true`, token.TK_BOOL),
	}

	c := ctx.New(nil)
	exp := value.Bool(true)

	act, e := o.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}

func TestOperand_Eval_7(t *testing.T) {
	o := operand{
		t: tok(`Dragonfly`, token.TK_STRING),
	}

	c := ctx.New(nil)
	exp := value.String(`Dragonfly`)

	act, e := o.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}
