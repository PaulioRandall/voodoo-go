package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestNumLex(t *testing.T) {
	lexErrFuncTest(t, "numLex", numLex, numLexTests())
}

func numLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `2`,
			ExpectLex: symbol.Lexeme{`2`, 0, 1, 0, symbol.NUMBER},
		},
		lexTest{
			Input:     `123`,
			ExpectLex: symbol.Lexeme{`123`, 0, 3, 0, symbol.NUMBER},
		},
		lexTest{
			Input:     `123_456`,
			ExpectLex: symbol.Lexeme{`123_456`, 0, 7, 0, symbol.NUMBER},
		},
		lexTest{
			Input:     `123.456`,
			ExpectLex: symbol.Lexeme{`123.456`, 0, 7, 0, symbol.NUMBER},
		},
		lexTest{
			Input:     `123.456_789`,
			ExpectLex: symbol.Lexeme{`123.456_789`, 0, 11, 0, symbol.NUMBER},
		},
		lexTest{
			Input:     `1__2__3__.__4__5__6__`,
			ExpectLex: symbol.Lexeme{`1__2__3__.__4__5__6__`, 0, 21, 0, symbol.NUMBER},
		},
		lexTest{
			Input:     `123..456`,
			ExpectLex: symbol.Lexeme{`123`, 0, 3, 0, symbol.NUMBER},
		},
		lexTest{
			Input:     `1_._2_._3`,
			ExpectErr: expLexError{0, 6},
		},
	}
}
