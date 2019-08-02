package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummyExe(p, r int) Exe {
	return Exe{
		Params:  p,
		Returns: r,
	}
}

func TestExeStack(t *testing.T) {
	s := ExeStack{}
	e1 := dummyExe(2, 1)
	e2 := dummyExe(4, 2)

	s.Push(e1)
	require.Len(t, s.stack, 1)
	assert.Equal(t, e1, s.stack[0])

	s.Push(e2)
	require.Len(t, s.stack, 2)
	assert.Equal(t, e2, s.stack[1])

	a2, ok := s.Pop()
	require.True(t, ok)
	assert.Equal(t, e2, a2)
	require.Len(t, s.stack, 1)
	assert.Equal(t, e1, s.stack[0])

	a1, ok := s.Peek()
	require.True(t, ok)
	assert.Equal(t, e1, a1)

	s.Push(e1)
	s.Push(e2)
	s.Push(e2)
	s.Sink(2)

	a, _ := s.Pop()
	assert.Equal(t, e1, a)
	a, _ = s.Pop()
	assert.Equal(t, e1, a)
	a, _ = s.Pop()
	assert.Equal(t, e2, a)
	a, _ = s.Pop()
	assert.Equal(t, e2, a)

	a0, ok := s.Pop()
	assert.False(t, ok)
	assert.Equal(t, Exe{}, a0)
}

func dummyToken(t token.TokenType) Token {
	return Token{
		Type: t,
	}
}

func TestValStack(t *testing.T) {
	s := ValStack{}
	e1 := dummyToken(token.CALC_ADD)
	e2 := dummyToken(token.CALC_SUBTRACT)

	s.Push(e1)
	require.Len(t, s.stack, 1)
	assert.Equal(t, e1, s.stack[0])

	s.Push(e2)
	require.Len(t, s.stack, 2)
	assert.Equal(t, e2, s.stack[1])

	a2, ok := s.Pop()
	require.True(t, ok)
	assert.Equal(t, e2, a2)
	require.Len(t, s.stack, 1)
	assert.Equal(t, e1, s.stack[0])

	a1, ok := s.Peek()
	require.True(t, ok)
	assert.Equal(t, e1, a1)

	s.Push(e1)
	s.Push(e2)
	s.Push(e2)
	s.Sink(2)

	a, _ := s.Pop()
	assert.Equal(t, e1, a)
	a, _ = s.Pop()
	assert.Equal(t, e1, a)
	a, _ = s.Pop()
	assert.Equal(t, e2, a)
	a, _ = s.Pop()
	assert.Equal(t, e2, a)

	a0, ok := s.Pop()
	assert.False(t, ok)
	assert.Equal(t, Token{}, a0)
}
