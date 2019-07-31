package exprs

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummyBool(n string) ctx.Expression {
	b := Bool{
		Bool: dummyToken(n),
	}
	expr := ctx.Expression(b)
	return expr
}

func TestBool_Evaluate_1(t *testing.T) {
	c := ctx.Empty()
	n := dummyBool(`true`)

	exp := ctx.BoolValue(true)
	v, err := n.Evaluate(c)

	assert.Nil(t, err)
	assert.Equal(t, exp, v)
	assert.Equal(t, ctx.Empty(), c)
}

func TestBool_Evaluate_2(t *testing.T) {
	c := ctx.Empty()
	n := dummyBool(`false`)

	exp := ctx.BoolValue(false)
	v, err := n.Evaluate(c)

	assert.Nil(t, err)
	assert.Equal(t, exp, v)
	assert.Equal(t, ctx.Empty(), c)
}

func TestBool_Evaluate_3(t *testing.T) {
	c := ctx.Empty()
	n := dummyBool(`abc`)

	v, err := n.Evaluate(c)

	assert.NotNil(t, err)
	require.Nil(t, v)
}
