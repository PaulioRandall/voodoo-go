package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestCommentSym(t *testing.T) {
	symFuncTest(t, "commentSym", commentSym, commentSymTests())
}

func commentSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `//`,
			ExpectSym: symbol.Symbol{`//`, 0, 2, 0, symbol.COMMENT},
		},
		symTest{
			Input:     `// A comment`,
			ExpectSym: symbol.Symbol{`// A comment`, 0, 12, 0, symbol.COMMENT},
		},
		symTest{
			Input:     `// Abc // 123 // xyz`,
			ExpectSym: symbol.Symbol{`// Abc // 123 // xyz`, 0, 20, 0, symbol.COMMENT},
		},
	}
}
