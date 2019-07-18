package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
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
				assert.Equal(t, tc.ExpectSym, *s)
			}
		}
	}
}

func sourcerySymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `@P`,
			ExpectSym: sym.Symbol{`@P`, 0, 2, 0, sym.SOURCERY},
		},
		symTest{
			Input:     `@Println`,
			ExpectSym: sym.Symbol{`@Println`, 0, 8, 0, sym.SOURCERY},
		},
		symTest{
			Input:     `@a__12__xy__`,
			ExpectSym: sym.Symbol{`@a__12__xy__`, 0, 12, 0, sym.SOURCERY},
		},
	}
}
