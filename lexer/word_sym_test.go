package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestWordSym(t *testing.T) {
	for i, tc := range wordSymTests() {
		t.Log("wordSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s := wordSym(itr)
		assert.Equal(t, tc.Expects, *s)
	}
}

func wordSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `a`,
			Expects: sym.Symbol{`a`, 0, 1, 0, sym.VARIABLE},
		},
		symTest{
			Input:   `abc`,
			Expects: sym.Symbol{`abc`, 0, 3, 0, sym.VARIABLE},
		},
		symTest{
			Input:   `abc_123`,
			Expects: sym.Symbol{`abc_123`, 0, 7, 0, sym.VARIABLE},
		},
		symTest{
			Input:   `a__________123456789`,
			Expects: sym.Symbol{`a__________123456789`, 0, 20, 0, sym.VARIABLE},
		},
		symTest{
			Input:   `SCROLL`,
			Expects: sym.Symbol{`SCROLL`, 0, 6, 0, sym.KEYWORD_SCROLL},
		},
		symTest{
			Input:   `sPeLL`,
			Expects: sym.Symbol{`sPeLL`, 0, 5, 0, sym.KEYWORD_SPELL},
		},
		symTest{
			Input:   `loop`,
			Expects: sym.Symbol{`loop`, 0, 4, 0, sym.KEYWORD_LOOP},
		},
		symTest{
			Input:   `when`,
			Expects: sym.Symbol{`when`, 0, 4, 0, sym.KEYWORD_WHEN},
		},
		symTest{
			Input:   `end`,
			Expects: sym.Symbol{`end`, 0, 3, 0, sym.KEYWORD_END},
		},
		symTest{
			Input:   `key`,
			Expects: sym.Symbol{`key`, 0, 3, 0, sym.KEYWORD_KEY},
		},
		symTest{
			Input:   `val`,
			Expects: sym.Symbol{`val`, 0, 3, 0, sym.KEYWORD_VAL},
		},
		symTest{
			Input:   `true`,
			Expects: sym.Symbol{`true`, 0, 4, 0, sym.BOOLEAN},
		},
		symTest{
			Input:   `false`,
			Expects: sym.Symbol{`false`, 0, 5, 0, sym.BOOLEAN},
		},
	}
}
