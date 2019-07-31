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
