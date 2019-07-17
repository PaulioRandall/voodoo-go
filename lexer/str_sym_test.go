package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestStrSym(t *testing.T) {
	for i, tc := range strSymTests() {
		t.Log("strSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s, err := strSym(itr)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			if assert.NotNil(t, s) {
				assert.Equal(t, tc.Expects, *s)
			}
		}
	}
}

func strSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `""`,
			Expects: sym.Symbol{`""`, 0, 2, 0, sym.STRING},
		},
		symTest{
			Input:   `"From hell with love"`,
			Expects: sym.Symbol{`"From hell with love"`, 0, 21, 0, sym.STRING},
		},
		symTest{
			Input:   `"Bam: \"Leaders eat last!\""`,
			Expects: sym.Symbol{`"Bam: \"Leaders eat last!\""`, 0, 28, 0, sym.STRING},
		},
		symTest{
			Input:   `"\\\\\""`,
			Expects: sym.Symbol{`"\\\\\""`, 0, 8, 0, sym.STRING},
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
			Input:   `"a"x`,
			Expects: sym.Symbol{`"a"`, 0, 3, 0, sym.STRING},
		},
		symTest{
			Input:     `"escaped \"`,
			ExpectErr: true,
		},
	}
}
