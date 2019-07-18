package lexer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/shared"
	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestStrSym(t *testing.T) {
	for i, tc := range strSymTests() {
		t.Log("strSym() test case: " + strconv.Itoa(i+1))

		itr := shared.NewRuneItr(tc.Input)
		s, err := strSym(itr)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			if assert.NotNil(t, s) {
				assert.Equal(t, tc.ExpectSym, *s)
			}
		}
	}
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
