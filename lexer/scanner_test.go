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
		symArrayTest{
			Input: "\tresult <- spell(a, b) r, err     ",
			Expects: []Symbol{
				Symbol{"\t", 0, 1, 0},
				Symbol{`result`, 1, 7, 0},
				Symbol{` `, 7, 8, 0},
				Symbol{`<-`, 8, 10, 0},
				Symbol{` `, 10, 11, 0},
				Symbol{`spell`, 11, 16, 0},
				Symbol{`(`, 16, 17, 0},
				Symbol{`a`, 17, 18, 0},
				Symbol{`,`, 18, 19, 0},
				Symbol{` `, 19, 20, 0},
				Symbol{`b`, 20, 21, 0},
				Symbol{`)`, 21, 22, 0},
				Symbol{` `, 22, 23, 0},
				Symbol{`r`, 23, 24, 0},
				Symbol{`,`, 24, 25, 0},
				Symbol{` `, 25, 26, 0},
				Symbol{`err`, 26, 29, 0},
				Symbol{`     `, 29, 34, 0},
			},
		},
		symArrayTest{
			Input: `keyValue <- "pi": 3.1419`,
			Expects: []Symbol{
				Symbol{`keyValue`, 0, 8, 0},
				Symbol{` `, 8, 9, 0},
				Symbol{`<-`, 9, 11, 0},
				Symbol{` `, 11, 12, 0},
				Symbol{`"pi"`, 12, 16, 0},
				Symbol{`:`, 16, 17, 0},
				Symbol{` `, 17, 18, 0},
				Symbol{`3.1419`, 18, 24, 0},
			},
		},
		// alphabet <- ["a", "b", "c"]
	}
}
