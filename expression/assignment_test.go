package expression

import (
	"testing"

	//"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func newToken(val string) token.Token {
	return token.Token{
		Val: val,
	}
}

func newNumber(n string) Expression {
	num := Number{
		Number: newToken(n),
	}
	expr := Expression(num)
	return expr
}

func TestAssignmentEvaluate_1(t *testing.T) {
	c := &Context{
		vars: map[string]Value{},
	}

	exp := &Context{
		vars: map[string]Value{
			`a`: NumberValue(123.456),
		},
	}

	a := &Assignment{
		Operator:   newToken(`<-`),
		Identifier: newToken(`a`),
		Expression: newNumber(`123.456`),
	}

	v, err := a.Evaluate(c)

	assert.Nil(t, v, "Did NOT expect a value to be returned for an assignment")
	assert.Nil(t, err, "Did NOT expect an fault for a valid assignment")
	assert.Equal(t, exp, c, "Context was not in the expected state")
}
