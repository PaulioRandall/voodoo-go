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
			Input:    []rune(" "),
			Output:   []rune{},
			Expect: token.Token{
				Val:  " ",
				Type: token.WHITESPACE,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune("\t"),
			Output:   []rune{},
			Expect: token.Token{
				Val:  "\t",
				Type: token.WHITESPACE,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune("\t\n \f \v\r"),
			Output:   []rune{},
			Expect: token.Token{
				Val:  "\t\n \f \v\r",
				Type: token.WHITESPACE,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune("   == 9"),
			Output:   []rune("== 9"),
			Expect: token.Token{
				Val:  "   ",
				Type: token.WHITESPACE,
			},
		},
	}
}
