package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestScanSpell(t *testing.T) {
	runFailableScanTest(t, "spell_test.go", scanSpell, scanSpellTests())
}

func scanSpellTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `@P`,
			Expect:   symbol.Token{`@P`, 0, 2, 0, symbol.SOURCERY},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `@Println`,
			Expect:   symbol.Token{`@Println`, 0, 8, 0, symbol.SOURCERY},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `@a__12__xy__`,
			Expect:   symbol.Token{`@a__12__xy__`, 0, 12, 0, symbol.SOURCERY},
		},
	}
}
