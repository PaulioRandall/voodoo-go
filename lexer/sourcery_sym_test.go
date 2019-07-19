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
			ExpectSym: lexeme.Symbol{`@P`, 0, 2, 0, lexeme.SOURCERY},
		},
		symTest{
			Input:     `@Println`,
			ExpectSym: lexeme.Symbol{`@Println`, 0, 8, 0, lexeme.SOURCERY},
		},
		symTest{
			Input:     `@a__12__xy__`,
			ExpectSym: lexeme.Symbol{`@a__12__xy__`, 0, 12, 0, lexeme.SOURCERY},
		},
	}
}
