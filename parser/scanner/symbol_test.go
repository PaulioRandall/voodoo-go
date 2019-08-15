package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanSymbol(t *testing.T) {
	runScanTest(t, "symbol_test.go", scanSymbol, scanSymbolTests())
}

func scanSymbolTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `==`,
			Expect:         dummyToken(0, 0, 2, `==`, token.TT_CMP_EQ),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `!=`,
			Expect:         dummyToken(0, 0, 2, `!=`, token.TT_CMP_NOT_EQ),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `<`,
			Expect:         dummyToken(0, 0, 1, `<`, token.TT_CMP_LT),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `<=`,
			Expect:         dummyToken(0, 0, 2, `<=`, token.TT_CMP_LT_OR_EQ),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `>`,
			Expect:         dummyToken(0, 0, 1, `>`, token.TT_CMP_MT),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `>=`,
			Expect:         dummyToken(0, 0, 2, `>=`, token.TT_CMP_MT_OR_EQ),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `||`,
			Expect:         dummyToken(0, 0, 2, `||`, token.TT_LOGIC_OR),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `&&`,
			Expect:         dummyToken(0, 0, 2, `&&`, token.TT_LOGIC_AND),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `<-`,
			Expect:         dummyToken(0, 0, 2, `<-`, token.TT_ASSIGN),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `=>`,
			Expect:         dummyToken(0, 0, 2, `=>`, token.TT_LOGIC_MATCH),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `_`,
			Expect:         dummyToken(0, 0, 1, `_`, token.VOID),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `!`,
			Expect:         dummyToken(0, 0, 1, `!`, token.TT_LOGIC_NOT),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `+`,
			Expect:         dummyToken(0, 0, 1, `+`, token.CALC_ADD),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `-`,
			Expect:         dummyToken(0, 0, 1, `-`, token.CALC_SUBTRACT),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `*`,
			Expect:         dummyToken(0, 0, 1, `*`, token.CALC_MULTIPLY),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `/`,
			Expect:         dummyToken(0, 0, 1, `/`, token.CALC_DIVIDE),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `%`,
			Expect:         dummyToken(0, 0, 1, `%`, token.CALC_MODULO),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `(`,
			Expect:         dummyToken(0, 0, 1, `(`, token.PAREN_CURVY_OPEN),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `)`,
			Expect:         dummyToken(0, 0, 1, `)`, token.PAREN_CURVY_CLOSE),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `[`,
			Expect:         dummyToken(0, 0, 1, `[`, token.PAREN_SQUARE_OPEN),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `]`,
			Expect:         dummyToken(0, 0, 1, `]`, token.PAREN_SQUARE_CLOSE),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `,`,
			Expect:         dummyToken(0, 0, 1, `,`, token.VALUE_DELIM),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `+ 69`,
			Expect:         dummyToken(0, 0, 1, `+`, token.CALC_ADD),
			NextUnreadRune: ' ',
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `=`,
			NextUnreadRune: '=',
			Error:          newFault(1),
		},
	}
}
