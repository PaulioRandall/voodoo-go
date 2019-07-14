package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScannerApi(t *testing.T) {
	for i, tc := range apiTests() {
		t.Log("Scanner test case: " + strconv.Itoa(i+1))
		a, err := ScanLine(tc.Input, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expects, a)
		}
	}
}

type symTest struct {
	Line      int
	Input     string
	Expects   Symbol
	ExpectErr bool
}

type symArrayTest struct {
	Line      int
	Input     string
	Expects   []Symbol
	ExpectErr bool
}

func apiTests() []symArrayTest {
	return []symArrayTest{
		symArrayTest{
			Input: `x = 1`,
			Expects: []Symbol{
				Symbol{`x`, 0, 1, 0},
				Symbol{` `, 1, 2, 0},
				Symbol{`=`, 2, 3, 0},
				Symbol{` `, 3, 4, 0},
				Symbol{`1`, 4, 5, 0},
			},
		},
		symArrayTest{
			Line:  123,
			Input: `x = true`,
			Expects: []Symbol{
				Symbol{`x`, 0, 1, 123},
				Symbol{` `, 1, 2, 123},
				Symbol{`=`, 2, 3, 123},
				Symbol{` `, 3, 4, 123},
				Symbol{`true`, 4, 8, 123},
			},
		},
		symArrayTest{
			Input: `@Println["Whelp"]`,
			Expects: []Symbol{
				Symbol{`@Println`, 0, 8, 0},
				Symbol{`[`, 8, 9, 0},
				Symbol{`"Whelp"`, 9, 16, 0},
				Symbol{`]`, 16, 17, 0},
			},
		},
	}
}
