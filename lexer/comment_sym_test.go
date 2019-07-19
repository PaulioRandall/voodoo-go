package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestCommentSym(t *testing.T) {
	symFuncTest(t, "commentSym", commentSym, commentSymTests())
}

func commentSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `//`,
			ExpectSym: lexeme.Lexeme{`//`, 0, 2, 0, lexeme.COMMENT},
		},
		symTest{
			Input:     `// A comment`,
			ExpectSym: lexeme.Lexeme{`// A comment`, 0, 12, 0, lexeme.COMMENT},
		},
		symTest{
			Input:     `// Abc // 123 // xyz`,
			ExpectSym: lexeme.Lexeme{`// Abc // 123 // xyz`, 0, 20, 0, lexeme.COMMENT},
		},
	}
}
