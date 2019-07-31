package exprs

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummyNumber(n string) ctx.Expression {
	num := Number{
		Number: dummyToken(n),
	}
	expr := ctx.Expression(num)
	return expr
}

func TestNumber_Evaluate_1(t *testing.T) {
	c := ctx.Empty()
	n := dummyNumber(`123.456`)

	exp := ctx.NumberValue(123.456)
	v, err := n.Evaluate(c)

	assert.Nil(t, err)
	assert.Equal(t, exp, v)
	assert.Equal(t, ctx.Empty(), c)
}

func TestNumber_Evaluate_2(t *testing.T) {
	c := ctx.Empty()
	n := dummyNumber(`abc`)

	v, err := n.Evaluate(c)

	assert.NotNil(t, err)
	require.Nil(t, v)
}
