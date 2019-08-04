package ctx

import (
	"testing"

	//"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func boolValue(b bool) Value {
	return Value{
		Val:  b,
		Type: BOOL_TYPE,
	}
}

func numValue(n float64) Value {
	return Value{
		Val:  n,
		Type: NUMBER_TYPE,
	}
}

func TestContextAssign_1(t *testing.T) {
	c := &Context{
		vars: map[string]Value{},
	}

	v := numValue(2)
	exp := &Context{
		vars: map[string]Value{
			`a`: v,
		},
	}

	err := c.Assign(`a`, v)

	assert.Nil(t, err)
	assert.Equal(t, exp, c)
}

func TestContextAssign_2(t *testing.T) {
	c := &Context{
		vars: map[string]Value{
			`a`: numValue(2),
		},
	}

	v := numValue(3)
	exp := &Context{
		vars: map[string]Value{
			`a`: v,
		},
	}

	err := c.Assign(`a`, v)

	assert.Nil(t, err)
	assert.Equal(t, exp, c)
}

func TestContextAssign_3(t *testing.T) {
	c := &Context{
		vars: map[string]Value{
			`a`: numValue(2),
		},
	}

	v := boolValue(true)
	err := c.Assign(`a`, v)
	assert.NotNil(t, err)
}

func TestContextGet_1(t *testing.T) {
	exp := numValue(2)

	c := &Context{
		vars: map[string]Value{
			`a`: exp,
		},
	}

	v, err := c.Get(`a`)
	assert.Nil(t, err)
	assert.Equal(t, exp, v)
}

func TestContextGet_2(t *testing.T) {
	c := &Context{
		vars: map[string]Value{},
	}

	v, err := c.Get(`a`)
	assert.Empty(t, v)
	assert.NotNil(t, err)
}
