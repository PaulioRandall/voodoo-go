package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanSpell(t *testing.T) {
	runScanTokenTests(t, "spell_test.go", scanSpell, scanSpellTests())
}

func scanSpellTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `@Println`,
			Expect:         dummyToken(0, 0, 8, `@Println`, token.TT_SPELL),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `@a__12__xy__`,
			Expect:         dummyToken(0, 0, 12, `@a__12__xy__`, token.TT_SPELL),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `@Println(msg)`,
			Expect:         dummyToken(0, 0, 8, `@Println`, token.TT_SPELL),
			NextUnreadRune: '(',
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `@2`,
			Expect:   errDummyToken(0, 0, 2),
		},
	}
}
