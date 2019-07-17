package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	"github.com/stretchr/testify/assert"
)

func TestSpaceSym(t *testing.T) {
	for i, tc := range spaceSymTests() {
		t.Log("spaceSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s, err := spaceSym(itr)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			if assert.NotNil(t, s) {
				assert.Equal(t, tc.Expects, *s)
			}
		}
	}
}

func spaceSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   " ",
			Expects: Symbol{" ", 0, 1, 0, WHITESPACE},
		},
		symTest{
			Input:   "\t",
			Expects: Symbol{"\t", 0, 1, 0, WHITESPACE},
		},
		symTest{
			Input:   "\t\n \f \v\r",
			Expects: Symbol{"\t\n \f \v\r", 0, 7, 0, WHITESPACE},
		},
	}
}
