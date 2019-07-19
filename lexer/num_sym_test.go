package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestNumSym(t *testing.T) {
	symErrFuncTest(t, "numSym", numSym, numSymTests())
}

func numSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `2`,
			ExpectSym: lexeme.Lexeme{`2`, 0, 1, 0, lexeme.NUMBER},
		},
		symTest{
			Input:     `123`,
			ExpectSym: lexeme.Lexeme{`123`, 0, 3, 0, lexeme.NUMBER},
		},
		symTest{
			Input:     `123_456`,
			ExpectSym: lexeme.Lexeme{`123_456`, 0, 7, 0, lexeme.NUMBER},
		},
		symTest{
			Input:     `123.456`,
			ExpectSym: lexeme.Lexeme{`123.456`, 0, 7, 0, lexeme.NUMBER},
		},
		symTest{
			Input:     `123.456_789`,
			ExpectSym: lexeme.Lexeme{`123.456_789`, 0, 11, 0, lexeme.NUMBER},
		},
		symTest{
			Input:     `1__2__3__.__4__5__6__`,
			ExpectSym: lexeme.Lexeme{`1__2__3__.__4__5__6__`, 0, 21, 0, lexeme.NUMBER},
		},
		symTest{
			Input:     `123..456`,
			ExpectSym: lexeme.Lexeme{`123`, 0, 3, 0, lexeme.NUMBER},
		},
		symTest{
			Input:     `1_._2_._3`,
			ExpectErr: expLexError{0, 6},
		},
	}
}
