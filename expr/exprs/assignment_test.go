package exprs

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/stretchr/testify/assert"
)

func TestAssignment_Evaluate_1(t *testing.T) {
	c := ctx.Empty()
	n := ctx.NumberValue(123.456)

	exp := ctx.New(map[string]ctx.Value{
		`a`: n,
	})

	a := &Assignment{
		Operator:   dummyToken(`<-`),
		Identifier: dummyToken(`a`),
		Expression: dummy{
			Val: n,
		},
	}

	v, err := a.Evaluate(c)

	assert.Nil(t, v)
	assert.Nil(t, err)
	assert.Equal(t, exp, c)
}

func TestAssignment_Evaluate_2(t *testing.T) {
	c := ctx.Empty()
	exp := ctx.Empty()

	a := &Assignment{
		Operator:   dummyToken(`<-`),
		Identifier: dummyToken(`a`),
		Expression: dummy{
			Err: ctx.EvalFault{},
		},
	}

	v, err := a.Evaluate(c)

	assert.Nil(t, v)
	assert.NotNil(t, err)
	assert.Equal(t, exp, c)
}
