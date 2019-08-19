package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanSpace(t *testing.T) {
	runScanTokenTests(t, "space_test.go", scanSpace, scanSpaceTests())
}

func dummySpaceToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_SPACE)
}

func scanSpaceTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          ` `,
			Expect:         dummySpaceToken(1, ` `),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "\t",
			Expect:         dummySpaceToken(1, "\t"),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "   abc",
			Expect:         dummySpaceToken(3, "   "),
			NextUnreadRune: 'a',
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "\t\f \n\v\r",
			Expect:         dummySpaceToken(3, "\t\f "),
			NextUnreadRune: '\n',
		},
	}
}
