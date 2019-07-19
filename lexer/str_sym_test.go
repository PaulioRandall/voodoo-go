package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestStrSym(t *testing.T) {
	symErrFuncTest(t, "strSym", strSym, strSymTests())
}

func strSymTests() []lexTest {
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
			Input:     `"Bam: \"Leaders eat last!\""`,
			ExpectLex: lexeme.Lexeme{`"Bam: \"Leaders eat last!\""`, 0, 28, 0, lexeme.STRING},
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
