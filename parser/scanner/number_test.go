package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanNumber(t *testing.T) {
	runScanTokenTests(t, "number_test.go", scanNumber, scanNumberTests())
}

func dummyNumToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_NUMBER)
}

func dummyNumErrToken(end int) token.Token {
	return dummyToken(0, 0, end, ``, token.TT_ERROR_UPSTREAM)
}

func scanNumberTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123`,
			Expect:         dummyNumToken(3, `123`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123 + 456`,
			Expect:         dummyNumToken(3, `123`),
			NextUnreadRune: ' ',
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123_456`,
			Expect:         dummyNumToken(7, `123_456`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123.456`,
			Expect:         dummyNumToken(7, `123.456`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `123.456_789`,
			Expect:         dummyNumToken(11, `123.456_789`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `1__2__3__.__4__5__6__`,
			Expect:         dummyNumToken(21, `1__2__3__.__4__5__6__`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `123.`,
			Expect:   dummyNumErrToken(4),
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `123..456`,
			Expect:   dummyNumErrToken(4),
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `123.___`,
			Expect:   dummyNumErrToken(7),
		},
	}
}
