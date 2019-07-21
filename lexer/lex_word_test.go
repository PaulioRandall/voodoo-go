package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestWordLex(t *testing.T) {
	lexFuncTest(t, "wordLex", wordLex, wordLexTests())
}

func wordLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `a`,
			ExpectLex: symbol.Lexeme{`a`, 0, 1, 0, symbol.IDENTIFIER},
		},
		lexTest{
			Input:     `abc`,
			ExpectLex: symbol.Lexeme{`abc`, 0, 3, 0, symbol.IDENTIFIER},
		},
		lexTest{
			Input:     `abc_123`,
			ExpectLex: symbol.Lexeme{`abc_123`, 0, 7, 0, symbol.IDENTIFIER},
		},
		lexTest{
			Input:     `a__________123456789`,
			ExpectLex: symbol.Lexeme{`a__________123456789`, 0, 20, 0, symbol.IDENTIFIER},
		},
		lexTest{
			Input:     `SCROLL`,
			ExpectLex: symbol.Lexeme{`SCROLL`, 0, 6, 0, symbol.KEYWORD_SCROLL},
		},
		lexTest{
			Input:     `sPeLL`,
			ExpectLex: symbol.Lexeme{`sPeLL`, 0, 5, 0, symbol.KEYWORD_SPELL},
		},
		lexTest{
			Input:     `loop`,
			ExpectLex: symbol.Lexeme{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
		},
		lexTest{
			Input:     `when`,
			ExpectLex: symbol.Lexeme{`when`, 0, 4, 0, symbol.KEYWORD_WHEN},
		},
		lexTest{
			Input:     `end`,
			ExpectLex: symbol.Lexeme{`end`, 0, 3, 0, symbol.KEYWORD_END},
		},
		lexTest{
			Input:     `key`,
			ExpectLex: symbol.Lexeme{`key`, 0, 3, 0, symbol.KEYWORD_KEY},
		},
		lexTest{
			Input:     `val`,
			ExpectLex: symbol.Lexeme{`val`, 0, 3, 0, symbol.KEYWORD_VAL},
		},
		lexTest{
			Input:     `true`,
			ExpectLex: symbol.Lexeme{`true`, 0, 4, 0, symbol.BOOLEAN_TRUE},
		},
		lexTest{
			Input:     `false`,
			ExpectLex: symbol.Lexeme{`false`, 0, 5, 0, symbol.BOOLEAN_FALSE},
		},
	}
}
