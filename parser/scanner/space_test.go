package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanSpace(t *testing.T) {
	runScanTokenTests(t, "space_test.go", scanSpace, scanSpaceTests())
}

func scanSpaceTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          ` `,
			Expect:         dummyToken(0, 0, 1, ` `, token.TT_SPACE),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "\t",
			Expect:         dummyToken(0, 0, 1, "\t", token.TT_SPACE),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "   abc",
			Expect:         dummyToken(0, 0, 3, "   ", token.TT_SPACE),
			NextUnreadRune: 'a',
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "\t\f \n\v\r",
			Expect:         dummyToken(0, 0, 3, "\t\f ", token.TT_SPACE),
			NextUnreadRune: '\n',
		},
	}
}
