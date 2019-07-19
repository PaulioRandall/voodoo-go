package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestSourceryLex(t *testing.T) {
	lexErrFuncTest(t, "sourceryLex", sourceryLex, sourceryLexTests())
}

func sourceryLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `@P`,
			ExpectLex: lexeme.Lexeme{`@P`, 0, 2, 0, lexeme.SOURCERY},
		},
		lexTest{
			Input:     `@Println`,
			ExpectLex: lexeme.Lexeme{`@Println`, 0, 8, 0, lexeme.SOURCERY},
		},
		lexTest{
			Input:     `@a__12__xy__`,
			ExpectLex: lexeme.Lexeme{`@a__12__xy__`, 0, 12, 0, lexeme.SOURCERY},
		},
	}
}
