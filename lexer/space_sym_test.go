package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpaceSym(t *testing.T) {
	for i, tc := range spaceSymTests() {
		t.Log("spaceSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
		a, err := spaceSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func spaceSymTests() []symTest {
	return []symTest{
		symTest{
			Line:    123,
			Input:   " ",
			Expects: Symbol{" ", 0, 1, 123},
		},
		symTest{
			Input:   "\t",
			Expects: Symbol{"\t", 0, 1, 0},
		},
		symTest{
			Input:   "\t\n \f \v\r",
			Expects: Symbol{"\t\n \f \v\r", 0, 7, 0},
		},
	}
}