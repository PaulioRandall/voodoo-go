package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestStrSym(t *testing.T) {
	symErrFuncTest(t, "strSym", strSym, strSymTests())
}

func strSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `""`,
			ExpectSym: lexeme.Symbol{`""`, 0, 2, 0, lexeme.STRING},
		},
		symTest{
			Input:     `"From hell with love"`,
			ExpectSym: lexeme.Symbol{`"From hell with love"`, 0, 21, 0, lexeme.STRING},
		},
		symTest{
			Input:     `"Bam: \"Leaders eat last!\""`,
			ExpectSym: lexeme.Symbol{`"Bam: \"Leaders eat last!\""`, 0, 28, 0, lexeme.STRING},
		},
		symTest{
			Input:     `"\\\\\""`,
			ExpectSym: lexeme.Symbol{`"\\\\\""`, 0, 8, 0, lexeme.STRING},
		},
		symTest{
			Input:     `"`,
			ExpectErr: expLexError{0, 1},
		},
		symTest{
			Input:     `"a"x`,
			ExpectSym: lexeme.Symbol{`"a"`, 0, 3, 0, lexeme.STRING},
		},
		symTest{
			Input:     `"escaped \"`,
			ExpectErr: expLexError{0, 11},
		},
	}
}
