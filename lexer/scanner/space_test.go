package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestScanSpace(t *testing.T) {
	runScanTest(t, "space_test.go", scanSpace, scanSpaceTests())
}

func scanSpaceTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    " ",
			Expect:   symbol.Token{" ", 0, 1, 0, symbol.WHITESPACE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    "\t",
			Expect:   symbol.Token{"\t", 0, 1, 0, symbol.WHITESPACE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    "\t\n \f \v\r",
			Expect:   symbol.Token{"\t\n \f \v\r", 0, 7, 0, symbol.WHITESPACE},
		},
	}
}
