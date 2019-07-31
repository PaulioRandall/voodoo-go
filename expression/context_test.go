package expression

import (
	"testing"

	//"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestContextAssign(t *testing.T) {
	c := &Context{
		vars: map[string]Value{},
	}
	v := NumberValue(2)
	err := c.Assign(`a`, v)

	exp := &Context{
		vars: map[string]Value{
			`a`: v,
		},
	}

	assert.Nil(t, err, "Did NOT expect an error for a valid assignment")
	assert.Equal(t, exp, c, "Context was not in the expected state")
}


