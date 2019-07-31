package exprs

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArithmetic_Evaluate_1(t *testing.T) {
	c := ctx.Empty()
	exp_c := ctx.Empty()

	left := ctx.NumberValue(2)
	right := ctx.NumberValue(4)
	exp_v := ctx.NumberValue(6)

	a := &Arithmetic{
		Operator: typedToken(`+`, token.CALC_ADD),
		Left:     valDummy(left),
		Right:    valDummy(right),
	}

	v, err := a.Evaluate(c)

	require.NotNil(t, v)
	assert.Nil(t, err)
	assert.Equal(t, exp_v, v)
	assert.Equal(t, exp_c, c)
}

func TestArithmetic_Evaluate_2(t *testing.T) {
	c := ctx.Empty()
	exp_c := ctx.Empty()

	left := ctx.NumberValue(8)
	right := ctx.NumberValue(5)
	exp_v := ctx.NumberValue(3)

	a := &Arithmetic{
		Operator: typedToken(`-`, token.CALC_SUBTRACT),
		Left:     valDummy(left),
		Right:    valDummy(right),
	}

	v, err := a.Evaluate(c)

	require.NotNil(t, v)
	assert.Nil(t, err)
	assert.Equal(t, exp_v, v)
	assert.Equal(t, exp_c, c)
}

func TestArithmetic_Evaluate_3(t *testing.T) {
	c := ctx.Empty()
	exp_c := ctx.Empty()

	left := ctx.NumberValue(3)
	right := ctx.NumberValue(3)
	exp_v := ctx.NumberValue(9)

	a := &Arithmetic{
		Operator: typedToken(`*`, token.CALC_MULTIPLY),
		Left:     valDummy(left),
		Right:    valDummy(right),
	}

	v, err := a.Evaluate(c)

	require.NotNil(t, v)
	assert.Nil(t, err)
	assert.Equal(t, exp_v, v)
	assert.Equal(t, exp_c, c)
}

func TestArithmetic_Evaluate_4(t *testing.T) {
	c := ctx.Empty()
	exp_c := ctx.Empty()

	left := ctx.NumberValue(12)
	right := ctx.NumberValue(6)
	exp_v := ctx.NumberValue(2)

	a := &Arithmetic{
		Operator: typedToken(`/`, token.CALC_DIVIDE),
		Left:     valDummy(left),
		Right:    valDummy(right),
	}

	v, err := a.Evaluate(c)

	require.NotNil(t, v)
	assert.Nil(t, err)
	assert.Equal(t, exp_v, v)
	assert.Equal(t, exp_c, c)
}

func TestArithmetic_Evaluate_5(t *testing.T) {
	c := ctx.Empty()
	exp_c := ctx.Empty()

	left := ctx.BoolValue(true)
	right := ctx.NumberValue(6)

	a := &Arithmetic{
		Operator: typedToken(`+`, token.CALC_ADD),
		Left:     valDummy(left),
		Right:    valDummy(right),
	}

	v, err := a.Evaluate(c)

	assert.Nil(t, v)
	assert.NotNil(t, err)
	assert.Equal(t, exp_c, c)
}

func TestArithmetic_Evaluate_6(t *testing.T) {
	c := ctx.Empty()
	exp_c := ctx.Empty()

	left := ctx.NumberValue(1)
	right := ctx.NumberValue(2)

	a := &Arithmetic{
		Operator: typedToken(`<-`, token.ASSIGNMENT),
		Left:     valDummy(left),
		Right:    valDummy(right),
	}

	v, err := a.Evaluate(c)

	assert.Nil(t, v)
	assert.NotNil(t, err)
	assert.Equal(t, exp_c, c)
}
