package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestScanSpace(t *testing.T) {
	lexFuncTest(t, "space_test.go", scanSpace, scanSpaceTests())
}

func scanSpaceTests() []lexTest {
	return []lexTest{
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    " ",
			Expect:   symbol.Lexeme{" ", 0, 1, 0, symbol.WHITESPACE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    "\t",
			Expect:   symbol.Lexeme{"\t", 0, 1, 0, symbol.WHITESPACE},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    "\t\n \f \v\r",
			Expect:   symbol.Lexeme{"\t\n \f \v\r", 0, 7, 0, symbol.WHITESPACE},
		},
	}
}
