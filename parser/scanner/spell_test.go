package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanSpell(t *testing.T) {
	runScanTokenTests(t, "spell_test.go", scanSpell, scanSpellTests())
}

func dummySpellToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_SPELL)
}

func scanSpellTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `@Println`,
			Expect:         dummySpellToken(8, `@Println`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `@a__12__xy__`,
			Expect:         dummySpellToken(12, `@a__12__xy__`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `@Println(msg)`,
			Expect:         dummySpellToken(8, `@Println`),
			NextUnreadRune: '(',
		},
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `@2`,
			Expect:   errDummyToken(0, 0, 2),
		},
	}
}
