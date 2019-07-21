package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestStrLex(t *testing.T) {
	lexErrFuncTest(t, "strLex", strLex, strLexTests())
}

func strLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `""`,
			ExpectLex: symbol.Lexeme{`""`, 0, 2, 0, symbol.STRING},
		},
		lexTest{
			Input:     `"From hell with love"`,
			ExpectLex: symbol.Lexeme{`"From hell with love"`, 0, 21, 0, symbol.STRING},
		},
		lexTest{
			Input:     `"Simon: \"Leaders eat last!\""`,
			ExpectLex: symbol.Lexeme{`"Simon: \"Leaders eat last!\""`, 0, 30, 0, symbol.STRING},
		},
		lexTest{
			Input:     `"\\\\\""`,
			ExpectLex: symbol.Lexeme{`"\\\\\""`, 0, 8, 0, symbol.STRING},
		},
		lexTest{
			Input:     `"`,
			ExpectErr: expLexError{0, 1},
		},
		lexTest{
			Input:     `"a"x`,
			ExpectLex: symbol.Lexeme{`"a"`, 0, 3, 0, symbol.STRING},
		},
		lexTest{
			Input:     `"escaped \"`,
			ExpectErr: expLexError{0, 11},
		},
	}
}
