package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurseSym(t *testing.T) {
	for i, tc := range curseSymTests() {
		t.Log("curseSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
		a, err := curseSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func curseSymTests() []symTest {
	return []symTest{
		symTest{
			Line:    123,
			Input:   `@P`,
			Expects: Symbol{`@P`, 0, 2, 123, UNDEFINED},
		},
		symTest{
			Input:   `@Println`,
			Expects: Symbol{`@Println`, 0, 8, 0, UNDEFINED},
		},
		symTest{
			Input:   `@a__12__xy__`,
			Expects: Symbol{`@a__12__xy__`, 0, 12, 0, UNDEFINED},
		},
	}
}
