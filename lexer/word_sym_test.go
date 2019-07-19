package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestWordSym(t *testing.T) {
	symFuncTest(t, "wordSym", wordSym, wordSymTests())
}

func wordSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `a`,
			ExpectSym: lexeme.Lexeme{`a`, 0, 1, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `abc`,
			ExpectSym: lexeme.Lexeme{`abc`, 0, 3, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `abc_123`,
			ExpectSym: lexeme.Lexeme{`abc_123`, 0, 7, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `a__________123456789`,
			ExpectSym: lexeme.Lexeme{`a__________123456789`, 0, 20, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `SCROLL`,
			ExpectSym: lexeme.Lexeme{`SCROLL`, 0, 6, 0, lexeme.KEYWORD_SCROLL},
		},
		symTest{
			Input:     `sPeLL`,
			ExpectSym: lexeme.Lexeme{`sPeLL`, 0, 5, 0, lexeme.KEYWORD_SPELL},
		},
		symTest{
			Input:     `loop`,
			ExpectSym: lexeme.Lexeme{`loop`, 0, 4, 0, lexeme.KEYWORD_LOOP},
		},
		symTest{
			Input:     `when`,
			ExpectSym: lexeme.Lexeme{`when`, 0, 4, 0, lexeme.KEYWORD_WHEN},
		},
		symTest{
			Input:     `end`,
			ExpectSym: lexeme.Lexeme{`end`, 0, 3, 0, lexeme.KEYWORD_END},
		},
		symTest{
			Input:     `key`,
			ExpectSym: lexeme.Lexeme{`key`, 0, 3, 0, lexeme.KEYWORD_KEY},
		},
		symTest{
			Input:     `val`,
			ExpectSym: lexeme.Lexeme{`val`, 0, 3, 0, lexeme.KEYWORD_VAL},
		},
		symTest{
			Input:     `true`,
			ExpectSym: lexeme.Lexeme{`true`, 0, 4, 0, lexeme.BOOLEAN},
		},
		symTest{
			Input:     `false`,
			ExpectSym: lexeme.Lexeme{`false`, 0, 5, 0, lexeme.BOOLEAN},
		},
	}
}
