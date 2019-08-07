package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanSpace(t *testing.T) {
	runScanTest(t, "space_test.go", scanSpace, scanSpaceTests())
}

func scanSpaceTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          ` `,
			Expect:         dummyToken(0, 0, 1, ` `, token.WHITESPACE),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          "\t",
			Expect:         dummyToken(0, 0, 1, "\t", token.WHITESPACE),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          "   abc",
			Expect:         dummyToken(0, 0, 3, "   ", token.WHITESPACE),
			NextUnreadRune: 'a',
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          "\t\f \n\v\r",
			Expect:         dummyToken(0, 0, 3, "\t\f ", token.WHITESPACE),
			NextUnreadRune: '\n',
		},
	}
}
