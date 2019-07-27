package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestString(t *testing.T) {
	lexErrFuncTest(t, "string_test.go", strLex, strLexTests())
}

func strLexTests() []lexTest {
	return []lexTest{
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `""`,
			Expect:   symbol.Lexeme{`""`, 0, 2, 0, symbol.LITERAL_STRING},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `"From hell with love"`,
			Expect:   symbol.Lexeme{`"From hell with love"`, 0, 21, 0, symbol.LITERAL_STRING},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `"Simon: \"Leaders eat last!\""`,
			Expect:   symbol.Lexeme{`"Simon: \"Leaders eat last!\""`, 0, 30, 0, symbol.LITERAL_STRING},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `"\\\\\""`,
			Expect:   symbol.Lexeme{`"\\\\\""`, 0, 8, 0, symbol.LITERAL_STRING},
		},
		lexTest{
			TestLine:  fault.CurrLine(),
			Input:     `"`,
			ExpectErr: fault.Dummy(fault.String).Line(0).From(0).To(1),
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `"a"x`,
			Expect:   symbol.Lexeme{`"a"`, 0, 3, 0, symbol.LITERAL_STRING},
		},
		lexTest{
			TestLine:  fault.CurrLine(),
			Input:     `"escaped \"`,
			ExpectErr: fault.Dummy(fault.String).Line(0).From(0).To(11),
		},
	}
}
