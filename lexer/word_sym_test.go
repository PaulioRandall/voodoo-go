package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	"github.com/stretchr/testify/assert"
)

func TestWordSym(t *testing.T) {
	for i, tc := range wordSymTests() {
		t.Log("wordSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		a, err := wordSym(itr)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func wordSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `a`,
			Expects: Symbol{`a`, 0, 1, 0, VARIABLE},
		},
		symTest{
			Input:   `abc`,
			Expects: Symbol{`abc`, 0, 3, 0, VARIABLE},
		},
		symTest{
			Input:   `abc_123`,
			Expects: Symbol{`abc_123`, 0, 7, 0, VARIABLE},
		},
		symTest{
			Input:   `a__________123456789`,
			Expects: Symbol{`a__________123456789`, 0, 20, 0, VARIABLE},
		},
		symTest{
			Input:   `SCROLL`,
			Expects: Symbol{`SCROLL`, 0, 6, 0, KEYWORD_SCROLL},
		},
		symTest{
			Input:   `sPeLL`,
			Expects: Symbol{`sPeLL`, 0, 5, 0, KEYWORD_SPELL},
		},
		symTest{
			Input:   `loop`,
			Expects: Symbol{`loop`, 0, 4, 0, KEYWORD_LOOP},
		},
		symTest{
			Input:   `when`,
			Expects: Symbol{`when`, 0, 4, 0, KEYWORD_WHEN},
		},
		symTest{
			Input:   `end`,
			Expects: Symbol{`end`, 0, 3, 0, KEYWORD_END},
		},
		symTest{
			Input:   `key`,
			Expects: Symbol{`key`, 0, 3, 0, KEYWORD_KEY},
		},
		symTest{
			Input:   `val`,
			Expects: Symbol{`val`, 0, 3, 0, KEYWORD_VAL},
		},
		symTest{
			Input:   `true`,
			Expects: Symbol{`true`, 0, 4, 0, BOOLEAN},
		},
		symTest{
			Input:   `false`,
			Expects: Symbol{`false`, 0, 5, 0, BOOLEAN},
		},
	}
}
