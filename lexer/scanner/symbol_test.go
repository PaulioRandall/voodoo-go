package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanSymbol(t *testing.T) {
	runFailableScanTest(t, "symbol_test.go", scanSymbol, scanSymbolTests())
}

func newToken(s string, t token.TokenType) token.Token {
	return token.Token{
		Val:  s,
		Type: t,
	}
}

func scanSymbolTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`==`),
			Output:   []rune{},
			Expect:   newToken(`==`, token.CMP_EQUAL),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`!=`),
			Output:   []rune{},
			Expect:   newToken(`!=`, token.CMP_NOT_EQUAL),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`<`),
			Output:   []rune{},
			Expect:   newToken(`<`, token.CMP_LESS_THAN),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`<=`),
			Output:   []rune{},
			Expect:   newToken(`<=`, token.CMP_LESS_THAN_OR_EQUAL),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`>`),
			Output:   []rune{},
			Expect:   newToken(`>`, token.CMP_GREATER_THAN),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`>=`),
			Output:   []rune{},
			Expect:   newToken(`>=`, token.CMP_GREATER_THAN_OR_EQUAL),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`||`),
			Output:   []rune{},
			Expect:   newToken(`||`, token.LOGICAL_OR),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`&&`),
			Output:   []rune{},
			Expect:   newToken(`&&`, token.LOGICAL_AND),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`<-`),
			Output:   []rune{},
			Expect:   newToken(`<-`, token.ASSIGNMENT),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`=>`),
			Output:   []rune{},
			Expect:   newToken(`=>`, token.LOGICAL_MATCH),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`_`),
			Output:   []rune{},
			Expect:   newToken(`_`, token.VOID),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`!`),
			Output:   []rune{},
			Expect:   newToken(`!`, token.LOGICAL_NOT),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`+`),
			Output:   []rune{},
			Expect:   newToken(`+`, token.CALC_ADD),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`-`),
			Output:   []rune{},
			Expect:   newToken(`-`, token.CALC_SUBTRACT),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`*`),
			Output:   []rune{},
			Expect:   newToken(`*`, token.CALC_MULTIPLY),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`/`),
			Output:   []rune{},
			Expect:   newToken(`/`, token.CALC_DIVIDE),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`%`),
			Output:   []rune{},
			Expect:   newToken(`%`, token.CALC_MODULO),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`(`),
			Output:   []rune{},
			Expect:   newToken(`(`, token.PAREN_CURVY_OPEN),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`)`),
			Output:   []rune{},
			Expect:   newToken(`)`, token.PAREN_CURVY_CLOSE),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`[`),
			Output:   []rune{},
			Expect:   newToken(`[`, token.PAREN_SQUARE_OPEN),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`]`),
			Output:   []rune{},
			Expect:   newToken(`]`, token.PAREN_SQUARE_CLOSE),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`,`),
			Output:   []rune{},
			Expect:   newToken(`,`, token.SEPARATOR_VALUE),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`+ 69`),
			Output:   []rune(` 69`),
			Expect:   newToken(`+`, token.CALC_ADD),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`=`),
			Error:    newFault(1),
		},
	}
}
