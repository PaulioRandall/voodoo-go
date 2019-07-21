package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestCommentLex(t *testing.T) {
	lexFuncTest(t, "commentLex", commentLex, commentLexTests())
}

func commentLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `//`,
			ExpectLex: symbol.Lexeme{`//`, 0, 2, 0, symbol.COMMENT},
		},
		lexTest{
			Input:     `// A comment`,
			ExpectLex: symbol.Lexeme{`// A comment`, 0, 12, 0, symbol.COMMENT},
		},
		lexTest{
			Input:     `// Abc // 123 // xyz`,
			ExpectLex: symbol.Lexeme{`// Abc // 123 // xyz`, 0, 20, 0, symbol.COMMENT},
		},
	}
}
