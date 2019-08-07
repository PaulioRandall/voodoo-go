package new_scanner

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
			TestLine:       fault.CurrLine(),
			Input:          `""`,
			Expect:         dummyToken(0, 0, 2, `""`, token.LITERAL_STRING),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `"From hell with love"`,
			Expect:         dummyToken(0, 0, 21, `"From hell with love"`, token.LITERAL_STRING),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `"From hell with love", 123.456`,
			Expect:         dummyToken(0, 0, 21, `"From hell with love"`, token.LITERAL_STRING),
			NextUnreadRune: ',',
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `"Simon: \"Leaders eat last!\""`,
			Expect:         dummyToken(0, 0, 30, `"Simon: \"Leaders eat last!\""`, token.LITERAL_STRING),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `"\\\\\""`,
			Expect:         dummyToken(0, 0, 8, `"\\\\\""`, token.LITERAL_STRING),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `"`,
			NextUnreadRune: '"',
			Error:          newFault(1),
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `"escaped \"`,
			NextUnreadRune: '"',
			Error:          newFault(11),
		},
	}
}
