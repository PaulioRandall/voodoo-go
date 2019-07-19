package lexeme

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

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
	ls := dummyLexArray(`a`, `b`, `c`)
	itr := NewLexItr(ls)
	exp := len(ls)
	assert.Equal(t, exp, itr.Length())
}

func TestLexItr_HasNext(t *testing.T) {
	ls := dummyLexArray(`a`, `b`, `c`)
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
