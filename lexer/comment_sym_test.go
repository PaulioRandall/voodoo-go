package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestCommentSym(t *testing.T) {
	symFuncTest(t, "commentSym", commentSym, commentSymTests())
}

func commentSymTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `//`,
			ExpectSym: lexeme.Lexeme{`//`, 0, 2, 0, lexeme.COMMENT},
		},
		lexTest{
			Input:     `// A comment`,
			ExpectSym: lexeme.Lexeme{`// A comment`, 0, 12, 0, lexeme.COMMENT},
		},
		lexTest{
			Input:     `// Abc // 123 // xyz`,
			ExpectSym: lexeme.Lexeme{`// Abc // 123 // xyz`, 0, 20, 0, lexeme.COMMENT},
		},
	}
}
