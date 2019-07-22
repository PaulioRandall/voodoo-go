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
