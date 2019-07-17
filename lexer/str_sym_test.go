package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
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
			Expects: Symbol{`""`, 0, 2, 0, STRING},
		},
		symTest{
			Input:   `"From hell with love"`,
			Expects: Symbol{`"From hell with love"`, 0, 21, 0, STRING},
		},
		symTest{
			Input:   `"Bam: \"Leaders eat last!\""`,
			Expects: Symbol{`"Bam: \"Leaders eat last!\""`, 0, 28, 0, STRING},
		},
		symTest{
			Input:   `"\\\\\""`,
			Expects: Symbol{`"\\\\\""`, 0, 8, 0, STRING},
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
			Expects: Symbol{`"a"`, 0, 3, 0, STRING},
		},
		symTest{
			Input:     `"escaped \"`,
			ExpectErr: true,
		},
	}
}
