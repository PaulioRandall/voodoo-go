package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestScanString(t *testing.T) {
	runFailableScanTest(t, "string_test.go", scanString, scanStringTests())
}

func scanStringTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `""`,
			Expect:   symbol.Token{`""`, 0, 2, 0, symbol.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"From hell with love"`,
			Expect:   symbol.Token{`"From hell with love"`, 0, 21, 0, symbol.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"Simon: \"Leaders eat last!\""`,
			Expect:   symbol.Token{`"Simon: \"Leaders eat last!\""`, 0, 30, 0, symbol.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"\\\\\""`,
			Expect:   symbol.Token{`"\\\\\""`, 0, 8, 0, symbol.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     `"`,
			ExpectErr: fault.Dummy(fault.String).Line(0).From(0).To(1),
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `"a"x`,
			Expect:   symbol.Token{`"a"`, 0, 3, 0, symbol.LITERAL_STRING},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     `"escaped \"`,
			ExpectErr: fault.Dummy(fault.String).Line(0).From(0).To(11),
		},
	}
}
