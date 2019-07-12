package lexer

/*
import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSym(t *testing.T) {
	for i, tc := range wordSymTests() {
		t.Log("wordSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
		a, err := strSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, tc.Expects, a)
		}
	}
}

// TODO: NEXT <<<<<<<<----------------------------
func wordSymTests() []symTest {
	return []symTest{
		symTest{
			Line:    0,
			Input:   `abc`,
			Expects: Symbol{`""`, 0, 2, 0},
		},
		symTest{
			Line:    123,
			Input:   `"From hell with love"`,
			Expects: Symbol{`"From hell with love"`, 0, 21, 123},
		},
		symTest{
			Line:    0,
			Input:   `"Bam: \"Leaders eat last!\""`,
			Expects: Symbol{`"Bam: \"Leaders eat last!\""`, 0, 28, 0},
		},
		symTest{
			Line:    0,
			Input:   `"\\\\\""`,
			Expects: Symbol{`"\\\\\""`, 0, 8, 0},
		},
	}
}
*/
