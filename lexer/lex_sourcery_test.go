package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestSourceryLex(t *testing.T) {
	lexErrFuncTest(t, "sourceryLex", sourceryLex, sourceryLexTests())
}

func sourceryLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `@P`,
			ExpectLex: symbol.Lexeme{`@P`, 0, 2, 0, symbol.SOURCERY},
		},
		lexTest{
			Input:     `@Println`,
			ExpectLex: symbol.Lexeme{`@Println`, 0, 8, 0, symbol.SOURCERY},
		},
		lexTest{
			Input:     `@a__12__xy__`,
			ExpectLex: symbol.Lexeme{`@a__12__xy__`, 0, 12, 0, symbol.SOURCERY},
		},
	}
}
