package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanNumber(t *testing.T) {
	runScanTest(t, "number_test.go", scanNumber, scanNumberTests())
}

func scanNumberTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `123`,
			Expect:         dummyToken(0, 0, 3, `123`, token.LITERAL_NUMBER),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `123 + 456`,
			Expect:         dummyToken(0, 0, 3, `123`, token.LITERAL_NUMBER),
			NextUnreadRune: ' ',
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `123_456`,
			Expect:         dummyToken(0, 0, 7, `123_456`, token.LITERAL_NUMBER),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `123.456`,
			Expect:         dummyToken(0, 0, 7, `123.456`, token.LITERAL_NUMBER),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `123.456_789`,
			Expect:         dummyToken(0, 0, 11, `123.456_789`, token.LITERAL_NUMBER),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `1__2__3__.__4__5__6__`,
			Expect:         dummyToken(0, 0, 21, `1__2__3__.__4__5__6__`, token.LITERAL_NUMBER),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123.`,
			Error:    newFault(4),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123..456`,
			Error:    newFault(4),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123.___`,
			Error:    newFault(4),
		},
	}
}
