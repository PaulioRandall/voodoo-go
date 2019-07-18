package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestSourcerySym(t *testing.T) {
	symErrFuncTest(t, "sourcerySym", sourcerySym, sourcerySymTests())
}

func sourcerySymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `@P`,
			ExpectSym: symbol.Symbol{`@P`, 0, 2, 0, symbol.SOURCERY},
		},
		symTest{
			Input:     `@Println`,
			ExpectSym: symbol.Symbol{`@Println`, 0, 8, 0, symbol.SOURCERY},
		},
		symTest{
			Input:     `@a__12__xy__`,
			ExpectSym: symbol.Symbol{`@a__12__xy__`, 0, 12, 0, symbol.SOURCERY},
		},
	}
}
