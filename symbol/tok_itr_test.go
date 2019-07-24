package symbol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var NIL_TOK *Token = nil

func dummyTokArray(ss ...string) []Token {
	r := []Token{}
	for _, s := range ss {
		t := Token{
			Val: s,
		}
		r = append(r, t)
	}
	return r
}

func TestTokItr_Length(t *testing.T) {
	ts := dummyTokArray(`a`, `b`, `語`)
	itr := NewTokItr(ts)
	exp := len(ts)
	assert.Equal(t, exp, itr.Length())
}

func TestTokItr_HasNext(t *testing.T) {
	a := dummyTokArray(`a`, `b`, `語`)
	itr := NewTokItr(a)

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

func TestTokItr_NextTok(t *testing.T) {
	ls := dummyTokArray(`a`, `b`, `語`)
	itr := NewTokItr(ls)

	assert.Equal(t, 0, itr.index)

	assert.Equal(t, &ls[0], itr.NextTok())
	assert.Equal(t, 1, itr.index)
	assert.Equal(t, &ls[1], itr.NextTok())
	assert.Equal(t, 2, itr.index)
	assert.Equal(t, &ls[2], itr.NextTok())
	assert.Equal(t, 3, itr.index)

	assert.Equal(t, NIL_TOK, itr.NextTok())
	assert.Equal(t, itr.length, itr.index)
}
