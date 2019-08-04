package ll_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {

	size := 9
	in_1 := `x`
	in_2 := `y`

	exp := make([]interface{}, size)
	exp[0] = in_1
	exp[1] = in_2

	s := newStack(size)
	s.Push(in_1)
	s.Push(in_2)

	assert.Equal(t, 2, s.i)
	assert.Equal(t, size, s.size)
	assert.Equal(t, exp, s.a)
}

func TestStack_Pop(t *testing.T) {

	size := 9
	exp := make([]interface{}, size)

	s := stack{
		a:    make([]interface{}, size),
		size: size,
		i:    2,
	}

	s.a[0] = `x`
	s.a[1] = `y`

	assert.Equal(t, `y`, s.Pop())
	assert.Equal(t, `x`, s.Pop())

	assert.Equal(t, 0, s.i)
	assert.Equal(t, size, s.size)
	assert.Equal(t, exp, s.a)
}
