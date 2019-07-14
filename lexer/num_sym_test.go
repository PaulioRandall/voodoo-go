package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSym(t *testing.T) {
	for i, tc := range numSymTests() {
		t.Log("numSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
		a, err := numSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func numSymTests() []symTest {
	return []symTest{
		symTest{
			Line:    123,
			Input:   `2`,
			Expects: Symbol{`2`, 0, 1, 123},
		},
		symTest{
			Input:   `123`,
			Expects: Symbol{`123`, 0, 3, 0},
		},
		symTest{
			Input:   `123_456`,
			Expects: Symbol{`123_456`, 0, 7, 0},
		},
		symTest{
			Input:   `123.456`,
			Expects: Symbol{`123.456`, 0, 7, 0},
		},
		symTest{
			Input:   `123.456_789`,
			Expects: Symbol{`123.456_789`, 0, 11, 0},
		},
		symTest{
			Input:   `1__2__3__.__4__5__6__`,
			Expects: Symbol{`1__2__3__.__4__5__6__`, 0, 21, 0},
		},
		symTest{
			Input:     `123..456`,
			ExpectErr: true,
		},
		symTest{
			Input:     `1_._2_._3`,
			ExpectErr: true,
		},
	}
}
