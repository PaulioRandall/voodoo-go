package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanString(t *testing.T) {
	runFailableScanTest(t, "string_test.go", scanString, scanStringTests())
}

func scanStringTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `""`,
			Expect:   token.Token{`""`, 0, 2, 0, token.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"From hell with love"`,
			Expect:   token.Token{`"From hell with love"`, 0, 21, 0, token.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"Simon: \"Leaders eat last!\""`,
			Expect:   token.Token{`"Simon: \"Leaders eat last!\""`, 0, 30, 0, token.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"\\\\\""`,
			Expect:   token.Token{`"\\\\\""`, 0, 8, 0, token.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     `"`,
			ExpectErr: fault.Dummy(fault.String).SetLine(0).SetFrom(0).SetTo(1),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"a"x`,
			Expect:   token.Token{`"a"`, 0, 3, 0, token.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     `"escaped \"`,
			ExpectErr: fault.Dummy(fault.String).SetLine(0).SetFrom(0).SetTo(11),
		},
	}
}
