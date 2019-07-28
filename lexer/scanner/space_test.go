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
			TestLine: fault.CurrLine(),
			Input:    " ",
			Expect:   token.Token{" ", 0, 1, 0, token.WHITESPACE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    "\t",
			Expect:   token.Token{"\t", 0, 1, 0, token.WHITESPACE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    "\t\n \f \v\r",
			Expect:   token.Token{"\t\n \f \v\r", 0, 7, 0, token.WHITESPACE},
		},
	}
}
