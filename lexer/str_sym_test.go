package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestStrSym(t *testing.T) {
	symErrFuncTest(t, "strSym", strSym, strSymTests())
}

func strSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `""`,
			ExpectSym: symbol.Symbol{`""`, 0, 2, 0, symbol.STRING},
		},
		symTest{
			Input:     `"From hell with love"`,
			ExpectSym: symbol.Symbol{`"From hell with love"`, 0, 21, 0, symbol.STRING},
		},
		symTest{
			Input:     `"Bam: \"Leaders eat last!\""`,
			ExpectSym: symbol.Symbol{`"Bam: \"Leaders eat last!\""`, 0, 28, 0, symbol.STRING},
		},
		symTest{
			Input:     `"\\\\\""`,
			ExpectSym: symbol.Symbol{`"\\\\\""`, 0, 8, 0, symbol.STRING},
		},
		symTest{
			Input:     ``,
			ExpectErr: true,
		},
		symTest{
			Input:     `:(`,
			ExpectErr: true,
		},
		symTest{
			Input:     `"`,
			ExpectErr: true,
		},
		symTest{
			Input:     `"a"x`,
			ExpectSym: symbol.Symbol{`"a"`, 0, 3, 0, symbol.STRING},
		},
		symTest{
			Input:     `"escaped \"`,
			ExpectErr: true,
		},
	}
}
