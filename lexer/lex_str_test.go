package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestStrLex(t *testing.T) {
	lexErrFuncTest(t, "strLex", strLex, strLexTests())
}

func strLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `""`,
			ExpectLex: lexeme.Lexeme{`""`, 0, 2, 0, lexeme.STRING},
		},
		lexTest{
			Input:     `"From hell with love"`,
			ExpectLex: lexeme.Lexeme{`"From hell with love"`, 0, 21, 0, lexeme.STRING},
		},
		lexTest{
			Input:     `"Simon: \"Leaders eat last!\""`,
			ExpectLex: lexeme.Lexeme{`"Simon: \"Leaders eat last!\""`, 0, 30, 0, lexeme.STRING},
		},
		lexTest{
			Input:     `"\\\\\""`,
			ExpectLex: lexeme.Lexeme{`"\\\\\""`, 0, 8, 0, lexeme.STRING},
		},
		lexTest{
			Input:     `"`,
			ExpectErr: expLexError{0, 1},
		},
		lexTest{
			Input:     `"a"x`,
			ExpectLex: lexeme.Lexeme{`"a"`, 0, 3, 0, lexeme.STRING},
		},
		lexTest{
			Input:     `"escaped \"`,
			ExpectErr: expLexError{0, 11},
		},
	}
}
