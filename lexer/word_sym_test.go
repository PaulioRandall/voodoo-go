package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestWordSym(t *testing.T) {
	symFuncTest(t, "wordSym", wordSym, wordSymTests())
}

func wordSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `a`,
			ExpectSym: symbol.Symbol{`a`, 0, 1, 0, symbol.VARIABLE},
		},
		symTest{
			Input:     `abc`,
			ExpectSym: symbol.Symbol{`abc`, 0, 3, 0, symbol.VARIABLE},
		},
		symTest{
			Input:     `abc_123`,
			ExpectSym: symbol.Symbol{`abc_123`, 0, 7, 0, symbol.VARIABLE},
		},
		symTest{
			Input:     `a__________123456789`,
			ExpectSym: symbol.Symbol{`a__________123456789`, 0, 20, 0, symbol.VARIABLE},
		},
		symTest{
			Input:     `SCROLL`,
			ExpectSym: symbol.Symbol{`SCROLL`, 0, 6, 0, symbol.KEYWORD_SCROLL},
		},
		symTest{
			Input:     `sPeLL`,
			ExpectSym: symbol.Symbol{`sPeLL`, 0, 5, 0, symbol.KEYWORD_SPELL},
		},
		symTest{
			Input:     `loop`,
			ExpectSym: symbol.Symbol{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
		},
		symTest{
			Input:     `when`,
			ExpectSym: symbol.Symbol{`when`, 0, 4, 0, symbol.KEYWORD_WHEN},
		},
		symTest{
			Input:     `end`,
			ExpectSym: symbol.Symbol{`end`, 0, 3, 0, symbol.KEYWORD_END},
		},
		symTest{
			Input:     `key`,
			ExpectSym: symbol.Symbol{`key`, 0, 3, 0, symbol.KEYWORD_KEY},
		},
		symTest{
			Input:     `val`,
			ExpectSym: symbol.Symbol{`val`, 0, 3, 0, symbol.KEYWORD_VAL},
		},
		symTest{
			Input:     `true`,
			ExpectSym: symbol.Symbol{`true`, 0, 4, 0, symbol.BOOLEAN},
		},
		symTest{
			Input:     `false`,
			ExpectSym: symbol.Symbol{`false`, 0, 5, 0, symbol.BOOLEAN},
		},
	}
}
