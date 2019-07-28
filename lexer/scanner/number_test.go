package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestNumber(t *testing.T) {
	lexErrFuncTest(t, "number_test.go", scanNumber, scanNumberTests())
}

func scanNumberTests() []lexTest {
	return []lexTest{
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `2`,
			Expect:   symbol.Lexeme{`2`, 0, 1, 0, symbol.LITERAL_NUMBER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `123`,
			Expect:   symbol.Lexeme{`123`, 0, 3, 0, symbol.LITERAL_NUMBER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `123_456`,
			Expect:   symbol.Lexeme{`123_456`, 0, 7, 0, symbol.LITERAL_NUMBER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `123.456`,
			Expect:   symbol.Lexeme{`123.456`, 0, 7, 0, symbol.LITERAL_NUMBER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `123.456_789`,
			Expect:   symbol.Lexeme{`123.456_789`, 0, 11, 0, symbol.LITERAL_NUMBER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `1__2__3__.__4__5__6__`,
			Expect:   symbol.Lexeme{`1__2__3__.__4__5__6__`, 0, 21, 0, symbol.LITERAL_NUMBER},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `123..456`,
			Expect:   symbol.Lexeme{`123`, 0, 3, 0, symbol.LITERAL_NUMBER},
		},
		lexTest{
			TestLine:  fault.CurrLine(),
			Input:     `1_._2_._3`,
			ExpectErr: fault.Dummy(fault.Number).Line(0).From(6),
		},
	}
}
