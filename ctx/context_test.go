package ctx

import (
	"testing"

	//"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestContextAssign_1(t *testing.T) {
	c := &Context{
		vars: map[string]Value{},
	}

	v := NumberValue(2)
	exp := &Context{
		vars: map[string]Value{
			`a`: v,
		},
	}

	err := c.Assign(`a`, v)

	assert.Nil(t, err, "Did NOT expect an fault for a valid assignment")
	assert.Equal(t, exp, c, "Context was not in the expected state")
}

func TestContextAssign_2(t *testing.T) {
	c := &Context{
		vars: map[string]Value{
			`a`: NumberValue(2),
		},
	}

	v := NumberValue(3)
	exp := &Context{
		vars: map[string]Value{
			`a`: v,
		},
	}

	err := c.Assign(`a`, v)

	assert.Nil(t, err, "Did NOT expect an fault for a valid assignment")
	assert.Equal(t, exp, c, "Context was not in the expected state")
}

func TestContextAssign_3(t *testing.T) {
	c := &Context{
		vars: map[string]Value{
			`a`: NumberValue(2),
		},
	}

	v := BoolValue(true)
	err := c.Assign(`a`, v)
	assert.NotNil(t, err, "Expected a fault for this invalid assignment")
}

func TestContextGet_1(t *testing.T) {
	exp := NumberValue(2)

	c := &Context{
		vars: map[string]Value{
			`a`: exp,
		},
	}

	v, err := c.Get(`a`)
	assert.Nil(t, err, "Expected the fault to be nil")
	assert.Equal(t, exp, v, "Expected a different value")
}

func TestContextGet_2(t *testing.T) {
	c := &Context{
		vars: map[string]Value{},
	}

	v, err := c.Get(`a`)
	assert.Nil(t, v, "Expected the value to be nil")
	assert.NotNil(t, err, "Expected a fault when attempting to read a non existent variable")
}
