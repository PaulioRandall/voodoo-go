package exprs

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/stretchr/testify/assert"
)

func TestNumber_Evaluate_1(t *testing.T) {
	c := ctx.Empty()
	n := dummyNumber(`123.456`)

	exp := ctx.NumberValue(123.456)
	v, err := n.Evaluate(c)

	assert.Nil(t, err, "Did NOT expect a fault for a parsing a valid literal number")
	assert.Equal(t, exp, v, "Expected a different value")
	assert.Equal(t, ctx.Empty(), c, "Did NOT expect context to have changed state")
}
