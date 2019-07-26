package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestWordLex(t *testing.T) {
	lexFuncTest(t, "lex_word_test.go", wordLex, wordLexTests())
}

func wordLexTests() []lexTest {
	return []lexTest{
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `a`,
			Expect:   symbol.Lexeme{`a`, 0, 1, 0, symbol.IDENTIFIER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `abc`,
			Expect:   symbol.Lexeme{`abc`, 0, 3, 0, symbol.IDENTIFIER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `abc_123`,
			Expect:   symbol.Lexeme{`abc_123`, 0, 7, 0, symbol.IDENTIFIER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `a__________123456789`,
			Expect:   symbol.Lexeme{`a__________123456789`, 0, 20, 0, symbol.IDENTIFIER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `sPeLL`,
			Expect:   symbol.Lexeme{`sPeLL`, 0, 5, 0, symbol.KEYWORD_SPELL},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `loop`,
			Expect:   symbol.Lexeme{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `when`,
			Expect:   symbol.Lexeme{`when`, 0, 4, 0, symbol.KEYWORD_WHEN},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `end`,
			Expect:   symbol.Lexeme{`end`, 0, 3, 0, symbol.KEYWORD_END},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `key`,
			Expect:   symbol.Lexeme{`key`, 0, 3, 0, symbol.KEYWORD_KEY},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `val`,
			Expect:   symbol.Lexeme{`val`, 0, 3, 0, symbol.KEYWORD_VAL},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `true`,
			Expect:   symbol.Lexeme{`true`, 0, 4, 0, symbol.BOOLEAN_TRUE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `false`,
			Expect:   symbol.Lexeme{`false`, 0, 5, 0, symbol.BOOLEAN_FALSE},
		},
	}
}
