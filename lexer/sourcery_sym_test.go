package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	"github.com/stretchr/testify/assert"
)

func TestSourcerySym(t *testing.T) {
	for i, tc := range sourcerySymTests() {
		t.Log("sourcerySym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s, err := sourcerySym(itr)

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

func sourcerySymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `@P`,
			Expects: Symbol{`@P`, 0, 2, 0, SOURCERY},
		},
		symTest{
			Input:   `@Println`,
			Expects: Symbol{`@Println`, 0, 8, 0, SOURCERY},
		},
		symTest{
			Input:   `@a__12__xy__`,
			Expects: Symbol{`@a__12__xy__`, 0, 12, 0, SOURCERY},
		},
	}
}
