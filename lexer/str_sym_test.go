package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrSym(t *testing.T) {
	for i, tc := range strSymTests() {
		t.Log("strSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
		a, err := strSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func strSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `""`,
			Expects: Symbol{`""`, 0, 2, 0},
		},
		symTest{
			Line:    123,
			Input:   `"From hell with love"`,
			Expects: Symbol{`"From hell with love"`, 0, 21, 123},
		},
		symTest{
			Input:   `"Bam: \"Leaders eat last!\""`,
			Expects: Symbol{`"Bam: \"Leaders eat last!\""`, 0, 28, 0},
		},
		symTest{
			Input:   `"\\\\\""`,
			Expects: Symbol{`"\\\\\""`, 0, 8, 0},
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
			Expects: Symbol{`"a"`, 0, 3, 0},
		},
		symTest{
			Input:     `"escaped \"`,
			ExpectErr: true,
		},
	}
}
