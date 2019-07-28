package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestScanSymbol(t *testing.T) {
	runFailableScanTest(t, "symbol_test.go", scanSymbol, scanSymbolTests())
}

func scanSymbolTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `==`,
			Expect:   symbol.Token{`==`, 0, 2, 0, symbol.CMP_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `!=`,
			Expect:   symbol.Token{`!=`, 0, 2, 0, symbol.CMP_NOT_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `<`,
			Expect:   symbol.Token{`<`, 0, 1, 0, symbol.CMP_LESS_THAN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `<=`,
			Expect:   symbol.Token{`<=`, 0, 2, 0, symbol.CMP_LESS_THAN_OR_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `>`,
			Expect:   symbol.Token{`>`, 0, 1, 0, symbol.CMP_GREATER_THAN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `>=`,
			Expect:   symbol.Token{`>=`, 0, 2, 0, symbol.CMP_GREATER_THAN_OR_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `||`,
			Expect:   symbol.Token{`||`, 0, 2, 0, symbol.LOGICAL_OR},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `&&`,
			Expect:   symbol.Token{`&&`, 0, 2, 0, symbol.LOGICAL_AND},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `<-`,
			Expect:   symbol.Token{`<-`, 0, 2, 0, symbol.ASSIGNMENT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `=>`,
			Expect:   symbol.Token{`=>`, 0, 2, 0, symbol.LOGICAL_MATCH},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `_`,
			Expect:   symbol.Token{`_`, 0, 1, 0, symbol.VOID},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `!`,
			Expect:   symbol.Token{`!`, 0, 1, 0, symbol.LOGICAL_NOT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `+`,
			Expect:   symbol.Token{`+`, 0, 1, 0, symbol.CALC_ADD},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `-`,
			Expect:   symbol.Token{`-`, 0, 1, 0, symbol.CALC_SUBTRACT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `*`,
			Expect:   symbol.Token{`*`, 0, 1, 0, symbol.CALC_MULTIPLY},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `/`,
			Expect:   symbol.Token{`/`, 0, 1, 0, symbol.CALC_DIVIDE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `%`,
			Expect:   symbol.Token{`%`, 0, 1, 0, symbol.CALC_MODULO},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `(`,
			Expect:   symbol.Token{`(`, 0, 1, 0, symbol.PAREN_CURVY_OPEN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `)`,
			Expect:   symbol.Token{`)`, 0, 1, 0, symbol.PAREN_CURVY_CLOSE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `[`,
			Expect:   symbol.Token{`[`, 0, 1, 0, symbol.PAREN_SQUARE_OPEN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `]`,
			Expect:   symbol.Token{`]`, 0, 1, 0, symbol.PAREN_SQUARE_CLOSE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `,`,
			Expect:   symbol.Token{`,`, 0, 1, 0, symbol.SEPARATOR_VALUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `:`,
			Expect:   symbol.Token{`:`, 0, 1, 0, symbol.SEPARATOR_KEY_VALUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `..`,
			Expect:   symbol.Token{`..`, 0, 2, 0, symbol.RANGE},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     `=`,
			ExpectErr: fault.Dummy(fault.Symbol).Line(0).From(0).To(1),
		},
	}
}
