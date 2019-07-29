package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanNumber(t *testing.T) {
	runFailableScanTest(t, "number_test.go", scanNumber, scanNumberTests())
}

func scanNumberTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`2`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `2`,
				Type: token.LITERAL_NUMBER,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`123`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `123`,
				Type: token.LITERAL_NUMBER,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`123 + 456`),
			Output:   []rune(` + 456`),
			Expect: token.Token{
				Val:  `123`,
				Type: token.LITERAL_NUMBER,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`123_456`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `123_456`,
				Type: token.LITERAL_NUMBER,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`123.456`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `123.456`,
				Type: token.LITERAL_NUMBER,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`123.456_789`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `123.456_789`,
				Type: token.LITERAL_NUMBER,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`1__2__3__.__4__5__6__`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `1__2__3__.__4__5__6__`,
				Type: token.LITERAL_NUMBER,
			},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     []rune(`123.`),
			ExpectErr: fault.Dummy(fault.Number),
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     []rune(`123..456`),
			ExpectErr: fault.Dummy(fault.Number),
		},
	}
}
