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
			ExpectSym: lexeme.Symbol{`a`, 0, 1, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `abc`,
			ExpectSym: lexeme.Symbol{`abc`, 0, 3, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `abc_123`,
			ExpectSym: lexeme.Symbol{`abc_123`, 0, 7, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `a__________123456789`,
			ExpectSym: lexeme.Symbol{`a__________123456789`, 0, 20, 0, lexeme.VARIABLE},
		},
		symTest{
			Input:     `SCROLL`,
			ExpectSym: lexeme.Symbol{`SCROLL`, 0, 6, 0, lexeme.KEYWORD_SCROLL},
		},
		symTest{
			Input:     `sPeLL`,
			ExpectSym: lexeme.Symbol{`sPeLL`, 0, 5, 0, lexeme.KEYWORD_SPELL},
		},
		symTest{
			Input:     `loop`,
			ExpectSym: lexeme.Symbol{`loop`, 0, 4, 0, lexeme.KEYWORD_LOOP},
		},
		symTest{
			Input:     `when`,
			ExpectSym: lexeme.Symbol{`when`, 0, 4, 0, lexeme.KEYWORD_WHEN},
		},
		symTest{
			Input:     `end`,
			ExpectSym: lexeme.Symbol{`end`, 0, 3, 0, lexeme.KEYWORD_END},
		},
		symTest{
			Input:     `key`,
			ExpectSym: lexeme.Symbol{`key`, 0, 3, 0, lexeme.KEYWORD_KEY},
		},
		symTest{
			Input:     `val`,
			ExpectSym: lexeme.Symbol{`val`, 0, 3, 0, lexeme.KEYWORD_VAL},
		},
		symTest{
			Input:     `true`,
			ExpectSym: lexeme.Symbol{`true`, 0, 4, 0, lexeme.BOOLEAN},
		},
		symTest{
			Input:     `false`,
			ExpectSym: lexeme.Symbol{`false`, 0, 5, 0, lexeme.BOOLEAN},
		},
	}
}
