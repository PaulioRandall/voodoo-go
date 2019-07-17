package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	"github.com/stretchr/testify/assert"
)

func TestCommentSym(t *testing.T) {
	for i, tc := range commentSymTests() {
		t.Log("commentSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		a, err := commentSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func commentSymTests() []symTest {
	return []symTest{
		symTest{
			Line:    123,
			Input:   `//`,
			Expects: Symbol{`//`, 0, 2, 123, COMMENT},
		},
		symTest{
			Input:   `// A comment`,
			Expects: Symbol{`// A comment`, 0, 12, 0, COMMENT},
		},
		symTest{
			Input:   `// Abc // 123 // xyz`,
			Expects: Symbol{`// Abc // 123 // xyz`, 0, 20, 0, COMMENT},
		},
	}
}
