package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanString(t *testing.T) {
	runScanTokenTests(t, "string_test.go", scanString, scanStringTests())
}

func scanStringTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `""`,
			Expect:         dummyToken(0, 0, 2, `""`, token.TT_STRING),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"From hell with love"`,
			Expect:         dummyToken(0, 0, 21, `"From hell with love"`, token.TT_STRING),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"From hell with love", 123.456`,
			Expect:         dummyToken(0, 0, 21, `"From hell with love"`, token.TT_STRING),
			NextUnreadRune: ',',
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"Simon: \"Leaders eat last!\""`,
			Expect:         dummyToken(0, 0, 30, `"Simon: \"Leaders eat last!\""`, token.TT_STRING),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"\\\\\""`,
			Expect:         dummyToken(0, 0, 8, `"\\\\\""`, token.TT_STRING),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `"`,
			Expect:   errDummyToken(0, 0, 1),
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `"escaped \"`,
			Expect:   errDummyToken(0, 0, 11),
		},
	}
}
