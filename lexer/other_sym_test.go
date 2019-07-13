package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOtherSym(t *testing.T) {
	for i, tc := range otherSymTests() {
		t.Log("otherSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
		a, err := otherSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func otherSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `<`,
			Expects: Symbol{`<`, 0, 1, 0},
		},
		symTest{
			Input:   `<=`,
			Expects: Symbol{`<=`, 0, 2, 0},
		},
	}
}
