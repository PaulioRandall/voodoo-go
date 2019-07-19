package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestWordSym(t *testing.T) {
	symFuncTest(t, "wordSym", wordSym, wordSymTests())
}

func wordSymTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `a`,
			ExpectLex: lexeme.Lexeme{`a`, 0, 1, 0, lexeme.VARIABLE},
		},
		lexTest{
			Input:     `abc`,
			ExpectLex: lexeme.Lexeme{`abc`, 0, 3, 0, lexeme.VARIABLE},
		},
		lexTest{
			Input:     `abc_123`,
			ExpectLex: lexeme.Lexeme{`abc_123`, 0, 7, 0, lexeme.VARIABLE},
		},
		lexTest{
			Input:     `a__________123456789`,
			ExpectLex: lexeme.Lexeme{`a__________123456789`, 0, 20, 0, lexeme.VARIABLE},
		},
		lexTest{
			Input:     `SCROLL`,
			ExpectLex: lexeme.Lexeme{`SCROLL`, 0, 6, 0, lexeme.KEYWORD_SCROLL},
		},
		lexTest{
			Input:     `sPeLL`,
			ExpectLex: lexeme.Lexeme{`sPeLL`, 0, 5, 0, lexeme.KEYWORD_SPELL},
		},
		lexTest{
			Input:     `loop`,
			ExpectLex: lexeme.Lexeme{`loop`, 0, 4, 0, lexeme.KEYWORD_LOOP},
		},
		lexTest{
			Input:     `when`,
			ExpectLex: lexeme.Lexeme{`when`, 0, 4, 0, lexeme.KEYWORD_WHEN},
		},
		lexTest{
			Input:     `end`,
			ExpectLex: lexeme.Lexeme{`end`, 0, 3, 0, lexeme.KEYWORD_END},
		},
		lexTest{
			Input:     `key`,
			ExpectLex: lexeme.Lexeme{`key`, 0, 3, 0, lexeme.KEYWORD_KEY},
		},
		lexTest{
			Input:     `val`,
			ExpectLex: lexeme.Lexeme{`val`, 0, 3, 0, lexeme.KEYWORD_VAL},
		},
		lexTest{
			Input:     `true`,
			ExpectLex: lexeme.Lexeme{`true`, 0, 4, 0, lexeme.BOOLEAN},
		},
		lexTest{
			Input:     `false`,
			ExpectLex: lexeme.Lexeme{`false`, 0, 5, 0, lexeme.BOOLEAN},
		},
	}
}
