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
			Input:    `2`,
			Expect:   token.Token{`2`, 0, 1, 0, token.LITERAL_NUMBER},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123`,
			Expect:   token.Token{`123`, 0, 3, 0, token.LITERAL_NUMBER},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123_456`,
			Expect:   token.Token{`123_456`, 0, 7, 0, token.LITERAL_NUMBER},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123.456`,
			Expect:   token.Token{`123.456`, 0, 7, 0, token.LITERAL_NUMBER},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123.456_789`,
			Expect:   token.Token{`123.456_789`, 0, 11, 0, token.LITERAL_NUMBER},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `1__2__3__.__4__5__6__`,
			Expect:   token.Token{`1__2__3__.__4__5__6__`, 0, 21, 0, token.LITERAL_NUMBER},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `123..456`,
			Expect:   token.Token{`123`, 0, 3, 0, token.LITERAL_NUMBER},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     `1_._2_._3`,
			ExpectErr: fault.Dummy(fault.Number).SetLine(0).SetFrom(6),
		},
	}
}
