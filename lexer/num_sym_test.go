package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestNumSym(t *testing.T) {
	symErrFuncTest(t, "numSym", numSym, numSymTests())
}

func numSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `2`,
			ExpectSym: symbol.Symbol{`2`, 0, 1, 0, symbol.NUMBER},
		},
		symTest{
			Input:     `123`,
			ExpectSym: symbol.Symbol{`123`, 0, 3, 0, symbol.NUMBER},
		},
		symTest{
			Input:     `123_456`,
			ExpectSym: symbol.Symbol{`123_456`, 0, 7, 0, symbol.NUMBER},
		},
		symTest{
			Input:     `123.456`,
			ExpectSym: symbol.Symbol{`123.456`, 0, 7, 0, symbol.NUMBER},
		},
		symTest{
			Input:     `123.456_789`,
			ExpectSym: symbol.Symbol{`123.456_789`, 0, 11, 0, symbol.NUMBER},
		},
		symTest{
			Input:     `1__2__3__.__4__5__6__`,
			ExpectSym: symbol.Symbol{`1__2__3__.__4__5__6__`, 0, 21, 0, symbol.NUMBER},
		},
		symTest{
			Input:     `123..456`,
			ExpectSym: symbol.Symbol{`123`, 0, 3, 0, symbol.NUMBER},
		},
		symTest{
			Input:     `1_._2_._3`,
			ExpectErr: expLexError{0, 6},
		},
	}
}
