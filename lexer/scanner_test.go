package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ETestScannerApi(t *testing.T) {
	for i, tc := range apiTests() {
		t.Log("Scanner test case: " + strconv.Itoa(i+1))
		a := ScanLine(tc.Input, tc.Line)
		assert.Equal(t, tc.Expects, a)
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
			Line:  0,
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
			Line:  9,
			Input: `x = "Whelp!"`,
			Expects: []Symbol{
				Symbol{`x`, 0, 1, 9},
				Symbol{` `, 1, 2, 9},
				Symbol{`=`, 2, 3, 9},
				Symbol{` `, 3, 4, 9},
				Symbol{`"`, 4, 5, 9},
				Symbol{`Whelp`, 5, 10, 9},
				Symbol{`!`, 10, 11, 9},
				Symbol{`"`, 11, 12, 9},
			},
		},
		symArrayTest{
			Line:  0,
			Input: `@Println("First line\nsecond line")`,
			Expects: []Symbol{
				Symbol{`@`, 0, 1, 0},
				Symbol{`Println`, 1, 8, 0},
				Symbol{`(`, 8, 9, 0},
				Symbol{`"`, 9, 10, 0},
				Symbol{`First`, 10, 15, 0},
				Symbol{` `, 15, 16, 0},
				Symbol{`line`, 16, 20, 0},
				Symbol{`\`, 20, 21, 0},
				Symbol{`n`, 21, 22, 0},
				Symbol{`second`, 22, 28, 0},
				Symbol{` `, 28, 29, 0},
				Symbol{`line`, 29, 33, 0},
				Symbol{`"`, 33, 34, 0},
				Symbol{`)`, 34, 35, 0},
			},
		},
	}
}
