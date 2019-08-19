package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanString(t *testing.T) {
	runScanTokenTests(t, "string_test.go", scanString, scanStringTests())
}

func dummyStrToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_STRING)
}

func scanStringTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `""`,
			Expect:         dummyStrToken(2, `""`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"From hell with love"`,
			Expect:         dummyStrToken(21, `"From hell with love"`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"From hell with love", 123.456`,
			Expect:         dummyStrToken(21, `"From hell with love"`),
			NextUnreadRune: ',',
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"Simon: \"Leaders eat last!\""`,
			Expect:         dummyStrToken(30, `"Simon: \"Leaders eat last!\""`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `"\\\\\""`,
			Expect:         dummyStrToken(8, `"\\\\\""`),
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
