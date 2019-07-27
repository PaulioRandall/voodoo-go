package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestSourceryLex(t *testing.T) {
	lexErrFuncTest(t, "lex_sourcery_test.go", sourceryLex, sourceryLexTests())
}

func sourceryLexTests() []lexTest {
	return []lexTest{
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `@P`,
			Expect:   symbol.Lexeme{`@P`, 0, 2, 0, symbol.SOURCERY},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `@Println`,
			Expect:   symbol.Lexeme{`@Println`, 0, 8, 0, symbol.SOURCERY},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `@a__12__xy__`,
			Expect:   symbol.Lexeme{`@a__12__xy__`, 0, 12, 0, symbol.SOURCERY},
		},
	}
}
