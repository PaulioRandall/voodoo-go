package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestSymbolLex(t *testing.T) {
	lexErrFuncTest(t, "lex_symbol_test.go", symbolLex, symbolLexTests())
}

func symbolLexTests() []lexTest {
	return []lexTest{
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `==`,
			Expect:   symbol.Lexeme{`==`, 0, 2, 0, symbol.CMP_EQUAL},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `!=`,
			Expect:   symbol.Lexeme{`!=`, 0, 2, 0, symbol.CMP_NOT_EQUAL},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `<`,
			Expect:   symbol.Lexeme{`<`, 0, 1, 0, symbol.CMP_LESS_THAN},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `<=`,
			Expect:   symbol.Lexeme{`<=`, 0, 2, 0, symbol.CMP_LESS_THAN_OR_EQUAL},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `>`,
			Expect:   symbol.Lexeme{`>`, 0, 1, 0, symbol.CMP_GREATER_THAN},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `>=`,
			Expect:   symbol.Lexeme{`>=`, 0, 2, 0, symbol.CMP_GREATER_THAN_OR_EQUAL},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `||`,
			Expect:   symbol.Lexeme{`||`, 0, 2, 0, symbol.LOGICAL_OR},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `&&`,
			Expect:   symbol.Lexeme{`&&`, 0, 2, 0, symbol.LOGICAL_AND},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `<-`,
			Expect:   symbol.Lexeme{`<-`, 0, 2, 0, symbol.ASSIGNMENT},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `=>`,
			Expect:   symbol.Lexeme{`=>`, 0, 2, 0, symbol.LOGICAL_MATCH},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `_`,
			Expect:   symbol.Lexeme{`_`, 0, 1, 0, symbol.VOID},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `!`,
			Expect:   symbol.Lexeme{`!`, 0, 1, 0, symbol.LOGICAL_NOT},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `+`,
			Expect:   symbol.Lexeme{`+`, 0, 1, 0, symbol.CALC_ADD},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `-`,
			Expect:   symbol.Lexeme{`-`, 0, 1, 0, symbol.CALC_SUBTRACT},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `*`,
			Expect:   symbol.Lexeme{`*`, 0, 1, 0, symbol.CALC_MULTIPLY},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `/`,
			Expect:   symbol.Lexeme{`/`, 0, 1, 0, symbol.CALC_DIVIDE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `%`,
			Expect:   symbol.Lexeme{`%`, 0, 1, 0, symbol.CALC_MODULO},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `(`,
			Expect:   symbol.Lexeme{`(`, 0, 1, 0, symbol.PAREN_CURVY_OPEN},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `)`,
			Expect:   symbol.Lexeme{`)`, 0, 1, 0, symbol.PAREN_CURVY_CLOSE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `[`,
			Expect:   symbol.Lexeme{`[`, 0, 1, 0, symbol.PAREN_SQUARE_OPEN},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `]`,
			Expect:   symbol.Lexeme{`]`, 0, 1, 0, symbol.PAREN_SQUARE_CLOSE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `,`,
			Expect:   symbol.Lexeme{`,`, 0, 1, 0, symbol.SEPARATOR_VALUE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `:`,
			Expect:   symbol.Lexeme{`:`, 0, 1, 0, symbol.SEPARATOR_KEY_VALUE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `..`,
			Expect:   symbol.Lexeme{`..`, 0, 2, 0, symbol.RANGE},
		},
		lexTest{
			TestLine:  fault.CurrLine(),
			Input:     `=`,
			ExpectErr: fault.Dummy(fault.Symbol).Line(0).From(0).To(1),
		},
	}
}
