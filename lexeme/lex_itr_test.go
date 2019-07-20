package lexeme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var NIL_LEX *Lexeme = nil

func dummyLexArray(ss ...string) []Lexeme {
	r := []Lexeme{}
	for _, s := range ss {
		l := Lexeme{
			Val: s,
		}
		r = append(r, l)
	}
	return r
}

func TestLexItr_Length(t *testing.T) {
	ls := dummyLexArray(`a`, `b`, `語`)
	itr := NewLexItr(ls)
	exp := len(ls)
	assert.Equal(t, exp, itr.Length())
}

func TestLexItr_HasNext(t *testing.T) {
	ls := dummyLexArray(`a`, `b`, `語`)
	itr := NewLexItr(ls)

	assert.True(t, itr.HasNext())
	itr.index += 1
	assert.True(t, itr.HasNext())
	itr.index += 1
	assert.True(t, itr.HasNext())
	itr.index += 1
	assert.False(t, itr.HasNext())
	itr.index += 1
	assert.False(t, itr.HasNext())
}

func TestLexItr_NextLex(t *testing.T) {
	ls := dummyLexArray(`a`, `b`, `語`)
	itr := NewLexItr(ls)

	assert.Equal(t, 0, itr.index)

	assert.Equal(t, &ls[0], itr.NextLex())
	assert.Equal(t, 1, itr.index)
	assert.Equal(t, &ls[1], itr.NextLex())
	assert.Equal(t, 2, itr.index)
	assert.Equal(t, &ls[2], itr.NextLex())
	assert.Equal(t, 3, itr.index)

	assert.Equal(t, NIL_LEX, itr.NextLex())
	assert.Equal(t, itr.length, itr.index)
}
