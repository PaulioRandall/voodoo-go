package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanNumber(t *testing.T) {
	runScanTokenTests(t, "number_test.go", scanNumber, scanNumberTests())
}

func scanNumberTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123`,
			Expect:         dummyToken(0, 0, 3, `123`, token.TT_NUMBER),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123 + 456`,
			Expect:         dummyToken(0, 0, 3, `123`, token.TT_NUMBER),
			NextUnreadRune: ' ',
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123_456`,
			Expect:         dummyToken(0, 0, 7, `123_456`, token.TT_NUMBER),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123.456`,
			Expect:         dummyToken(0, 0, 7, `123.456`, token.TT_NUMBER),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123.456_789`,
			Expect:         dummyToken(0, 0, 11, `123.456_789`, token.TT_NUMBER),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `1__2__3__.__4__5__6__`,
			Expect:         dummyToken(0, 0, 21, `1__2__3__.__4__5__6__`, token.TT_NUMBER),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `123.`,
			Expect:   errDummyToken(0, 0, 4),
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `123..456`,
			Expect:   errDummyToken(0, 0, 4),
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `123.___`,
			Expect:   errDummyToken(0, 0, 7),
		},
	}
}
