package lexeme

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func dummySymArray(s ...string) []Symbol {
	r := []Symbol{}
	for _, v := range s {
		sym := Symbol{
			Val: v,
		}
		r = append(r, sym)
	}
	return r
}

func TestSymItr_Length(t *testing.T) {
	s := dummySymArray(`a`, `b`, `c`)
	itr := NewSymItr(s)
	exp := len(s)
	assert.Equal(t, exp, itr.Length())
}

func TestSymItr_HasNext(t *testing.T) {
	s := dummySymArray(`a`, `b`, `c`)
	itr := NewSymItr(s)

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
