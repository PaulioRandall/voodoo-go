package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestCommentLex(t *testing.T) {
	lexFuncTest(t, "commentLex", commentLex, commentLexTests())
}

func commentLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `//`,
			ExpectLex: lexeme.Lexeme{`//`, 0, 2, 0, lexeme.COMMENT},
		},
		lexTest{
			Input:     `// A comment`,
			ExpectLex: lexeme.Lexeme{`// A comment`, 0, 12, 0, lexeme.COMMENT},
		},
		lexTest{
			Input:     `// Abc // 123 // xyz`,
			ExpectLex: lexeme.Lexeme{`// Abc // 123 // xyz`, 0, 20, 0, lexeme.COMMENT},
		},
	}
}