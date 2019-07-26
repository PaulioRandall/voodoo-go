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
			Expect:   symbol.Lexeme{`+`, 0, 1, 0, symbol.ADD},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `-`,
			Expect:   symbol.Lexeme{`-`, 0, 1, 0, symbol.SUBTRACT},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `*`,
			Expect:   symbol.Lexeme{`*`, 0, 1, 0, symbol.MULTIPLY},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `/`,
			Expect:   symbol.Lexeme{`/`, 0, 1, 0, symbol.DIVIDE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `%`,
			Expect:   symbol.Lexeme{`%`, 0, 1, 0, symbol.MODULO},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `(`,
			Expect:   symbol.Lexeme{`(`, 0, 1, 0, symbol.CURVED_BRACE_OPEN},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `)`,
			Expect:   symbol.Lexeme{`)`, 0, 1, 0, symbol.CURVED_BRACE_CLOSE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `[`,
			Expect:   symbol.Lexeme{`[`, 0, 1, 0, symbol.SQUARE_BRACE_OPEN},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `]`,
			Expect:   symbol.Lexeme{`]`, 0, 1, 0, symbol.SQUARE_BRACE_CLOSE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `,`,
			Expect:   symbol.Lexeme{`,`, 0, 1, 0, symbol.VALUE_SEPARATOR},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `:`,
			Expect:   symbol.Lexeme{`:`, 0, 1, 0, symbol.KEY_VALUE_SEPARATOR},
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
