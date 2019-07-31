package exprs

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/stretchr/testify/assert"
)

func TestAssignment_Evaluate_1(t *testing.T) {
	c := ctx.Empty()
	exp := ctx.New(map[string]ctx.Value{
		`a`: ctx.NumberValue(123.456),
	})

	a := &Assignment{
		Operator:   dummyToken(`<-`),
		Identifier: dummyToken(`a`),
		Expression: dummyNumber(`123.456`),
	}

	v, err := a.Evaluate(c)

	assert.Nil(t, v, "Did NOT expect a value to be returned for an assignment")
	assert.Nil(t, err, "Did NOT expect an fault for a valid assignment")
	assert.Equal(t, exp, c, "Context was not in the expected state")
}
