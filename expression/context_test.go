package expression

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
