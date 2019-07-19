package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestNumLex(t *testing.T) {
	lexErrFuncTest(t, "numLex", numLex, numLexTests())
}

func numLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `2`,
			ExpectLex: lexeme.Lexeme{`2`, 0, 1, 0, lexeme.NUMBER},
		},
		lexTest{
			Input:     `123`,
			ExpectLex: lexeme.Lexeme{`123`, 0, 3, 0, lexeme.NUMBER},
		},
		lexTest{
			Input:     `123_456`,
			ExpectLex: lexeme.Lexeme{`123_456`, 0, 7, 0, lexeme.NUMBER},
		},
		lexTest{
			Input:     `123.456`,
			ExpectLex: lexeme.Lexeme{`123.456`, 0, 7, 0, lexeme.NUMBER},
		},
		lexTest{
			Input:     `123.456_789`,
			ExpectLex: lexeme.Lexeme{`123.456_789`, 0, 11, 0, lexeme.NUMBER},
		},
		lexTest{
			Input:     `1__2__3__.__4__5__6__`,
			ExpectLex: lexeme.Lexeme{`1__2__3__.__4__5__6__`, 0, 21, 0, lexeme.NUMBER},
		},
		lexTest{
			Input:     `123..456`,
			ExpectLex: lexeme.Lexeme{`123`, 0, 3, 0, lexeme.NUMBER},
		},
		lexTest{
			Input:     `1_._2_._3`,
			ExpectErr: expLexError{0, 6},
		},
	}
}
