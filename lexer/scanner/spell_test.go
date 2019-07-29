package scanner

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
			TestLine: fault.CurrLine(),
			Input:    []rune(`@P`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `@P`,
				Type: token.SPELL,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`@Println`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `@Println`,
				Type: token.SPELL,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`@a__12__xy__`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `@a__12__xy__`,
				Type: token.SPELL,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`@Println(msg)`),
			Output:   []rune(`(msg)`),
			Expect: token.Token{
				Val:  `@Println`,
				Type: token.SPELL,
			},
		},
		scanFuncTest{
			TestLine:  fault.CurrLine(),
			Input:     []rune(`@2`),
			ExpectErr: fault.Dummy(fault.Function),
		},
	}
}
