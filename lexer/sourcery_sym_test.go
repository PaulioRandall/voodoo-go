package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestSourcerySym(t *testing.T) {
	symErrFuncTest(t, "sourcerySym", sourcerySym, sourcerySymTests())
}

func sourcerySymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `@P`,
			ExpectSym: lexeme.Lexeme{`@P`, 0, 2, 0, lexeme.SOURCERY},
		},
		symTest{
			Input:     `@Println`,
			ExpectSym: lexeme.Lexeme{`@Println`, 0, 8, 0, lexeme.SOURCERY},
		},
		symTest{
			Input:     `@a__12__xy__`,
			ExpectSym: lexeme.Lexeme{`@a__12__xy__`, 0, 12, 0, lexeme.SOURCERY},
		},
	}
}
