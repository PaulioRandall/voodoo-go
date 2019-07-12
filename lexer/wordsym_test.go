package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSym(t *testing.T) {
	for i, tc := range wordSymTests() {
		t.Log("wordSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
		a, err := wordSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func wordSymTests() []symTest {
	return []symTest{
		symTest{
			Line:    123,
			Input:   `a`,
			Expects: Symbol{`a`, 0, 1, 123},
		},
		symTest{
			Input:   `abc`,
			Expects: Symbol{`abc`, 0, 3, 0},
		},
		symTest{
			Input:   `abc_123`,
			Expects: Symbol{`abc_123`, 0, 7, 0},
		},
		symTest{
			Input:   `a__________123456789`,
			Expects: Symbol{`a__________123456789`, 0, 20, 0},
		},
	}
}
