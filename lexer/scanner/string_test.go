package scanner

import (
	"testing"

	fault "github.com/PaulioRandall/voodoo-go/new_fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanString(t *testing.T) {
	runFailableScanTest(t, "string_test.go", scanString, scanStringTests())
}

func scanStringTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`""`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `""`,
				Type: token.LITERAL_STRING,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`"From hell with love"`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `"From hell with love"`,
				Type: token.LITERAL_STRING,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`"From hell with love", 123.456`),
			Output:   []rune(`, 123.456`),
			Expect: token.Token{
				Val:  `"From hell with love"`,
				Type: token.LITERAL_STRING,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`"Simon: \"Leaders eat last!\""`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `"Simon: \"Leaders eat last!\""`,
				Type: token.LITERAL_STRING,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`"\\\\\""`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `"\\\\\""`,
				Type: token.LITERAL_STRING,
			},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     []rune(`"`),
			ExpectErr: newFault(1),
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     []rune(`"escaped \"`),
			ExpectErr: newFault(11),
		},
	}
}
