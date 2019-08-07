package new_scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanSpell(t *testing.T) {
	runFailableScanTest(t, "spell_test.go", scanSpell, scanSpellTests())
}

func scanSpellTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `@Println`,
			Expect:         dummyToken(0, 0, 8, `@Println`, token.SPELL),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `@a__12__xy__`,
			Expect:         dummyToken(0, 0, 12, `@a__12__xy__`, token.SPELL),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `@Println(msg)`,
			Expect:         dummyToken(0, 0, 8, `@Println`, token.SPELL),
			NextUnreadRune: '(',
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `@2`,
			Error:    newFault(1),
		},
	}
}
