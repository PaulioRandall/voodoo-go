package symbol

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
	s := dummySymArray()
	itr := NewSymItr(s)
	exp := len(s)
	assert.Equal(t, exp, itr.Length())
}
