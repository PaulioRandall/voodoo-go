package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanSymbol(t *testing.T) {
	runFailableScanTest(t, "symbol_test.go", scanSymbol, scanSymbolTests())
}

func scanSymbolTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `==`,
			Expect:   token.Token{`==`, 0, 2, 0, token.CMP_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `!=`,
			Expect:   token.Token{`!=`, 0, 2, 0, token.CMP_NOT_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `<`,
			Expect:   token.Token{`<`, 0, 1, 0, token.CMP_LESS_THAN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `<=`,
			Expect:   token.Token{`<=`, 0, 2, 0, token.CMP_LESS_THAN_OR_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `>`,
			Expect:   token.Token{`>`, 0, 1, 0, token.CMP_GREATER_THAN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `>=`,
			Expect:   token.Token{`>=`, 0, 2, 0, token.CMP_GREATER_THAN_OR_EQUAL},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `||`,
			Expect:   token.Token{`||`, 0, 2, 0, token.LOGICAL_OR},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `&&`,
			Expect:   token.Token{`&&`, 0, 2, 0, token.LOGICAL_AND},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `<-`,
			Expect:   token.Token{`<-`, 0, 2, 0, token.ASSIGNMENT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `=>`,
			Expect:   token.Token{`=>`, 0, 2, 0, token.LOGICAL_MATCH},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `_`,
			Expect:   token.Token{`_`, 0, 1, 0, token.VOID},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `!`,
			Expect:   token.Token{`!`, 0, 1, 0, token.LOGICAL_NOT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `+`,
			Expect:   token.Token{`+`, 0, 1, 0, token.CALC_ADD},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `-`,
			Expect:   token.Token{`-`, 0, 1, 0, token.CALC_SUBTRACT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `*`,
			Expect:   token.Token{`*`, 0, 1, 0, token.CALC_MULTIPLY},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `/`,
			Expect:   token.Token{`/`, 0, 1, 0, token.CALC_DIVIDE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `%`,
			Expect:   token.Token{`%`, 0, 1, 0, token.CALC_MODULO},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `(`,
			Expect:   token.Token{`(`, 0, 1, 0, token.PAREN_CURVY_OPEN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `)`,
			Expect:   token.Token{`)`, 0, 1, 0, token.PAREN_CURVY_CLOSE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `[`,
			Expect:   token.Token{`[`, 0, 1, 0, token.PAREN_SQUARE_OPEN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `]`,
			Expect:   token.Token{`]`, 0, 1, 0, token.PAREN_SQUARE_CLOSE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `,`,
			Expect:   token.Token{`,`, 0, 1, 0, token.SEPARATOR_VALUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `:`,
			Expect:   token.Token{`:`, 0, 1, 0, token.SEPARATOR_KEY_VALUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `..`,
			Expect:   token.Token{`..`, 0, 2, 0, token.RANGE},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     `=`,
			ExpectErr: fault.Dummy(fault.Symbol).SetLine(0).SetFrom(0).SetTo(1),
		},
	}
}
