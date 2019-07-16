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
			Input: `x <- 1`,
			Expects: []Symbol{
				Symbol{`x`, 0, 1, 0, VARIABLE},
				Symbol{` `, 1, 2, 0, WHITESPACE},
				Symbol{`<-`, 2, 4, 0, ASSIGNMENT},
				Symbol{` `, 4, 5, 0, WHITESPACE},
				Symbol{`1`, 5, 6, 0, NUMBER},
			},
		},
		symArrayTest{
			Input: `y <- -1.1`,
			Expects: []Symbol{
				Symbol{`y`, 0, 1, 0, VARIABLE},
				Symbol{` `, 1, 2, 0, WHITESPACE},
				Symbol{`<-`, 2, 4, 0, ASSIGNMENT},
				Symbol{` `, 4, 5, 0, WHITESPACE},
				Symbol{`-`, 5, 6, 0, SUBTRACT},
				Symbol{`1.1`, 6, 9, 0, NUMBER},
			},
		},
		symArrayTest{
			Line:  123,
			Input: `x <- true`,
			Expects: []Symbol{
				Symbol{`x`, 0, 1, 123, VARIABLE},
				Symbol{` `, 1, 2, 123, WHITESPACE},
				Symbol{`<-`, 2, 4, 123, ASSIGNMENT},
				Symbol{` `, 4, 5, 123, WHITESPACE},
				Symbol{`true`, 5, 9, 123, BOOLEAN},
			},
		},
		symArrayTest{
			Input: `@Println["Whelp"]`,
			Expects: []Symbol{
				Symbol{`@Println`, 0, 8, 0, SOURCERY},
				Symbol{`[`, 8, 9, 0, UNDEFINED},
				Symbol{`"Whelp"`, 9, 16, 0, STRING},
				Symbol{`]`, 16, 17, 0, UNDEFINED},
			},
		},
		symArrayTest{
			Input: "\tresult <- spell(a, b) r, err     ",
			Expects: []Symbol{
				Symbol{"\t", 0, 1, 0, WHITESPACE},
				Symbol{`result`, 1, 7, 0, VARIABLE},
				Symbol{` `, 7, 8, 0, WHITESPACE},
				Symbol{`<-`, 8, 10, 0, ASSIGNMENT},
				Symbol{` `, 10, 11, 0, WHITESPACE},
				Symbol{`spell`, 11, 16, 0, KEYWORD_SPELL},
				Symbol{`(`, 16, 17, 0, UNDEFINED},
				Symbol{`a`, 17, 18, 0, VARIABLE},
				Symbol{`,`, 18, 19, 0, UNDEFINED},
				Symbol{` `, 19, 20, 0, WHITESPACE},
				Symbol{`b`, 20, 21, 0, VARIABLE},
				Symbol{`)`, 21, 22, 0, UNDEFINED},
				Symbol{` `, 22, 23, 0, WHITESPACE},
				Symbol{`r`, 23, 24, 0, VARIABLE},
				Symbol{`,`, 24, 25, 0, UNDEFINED},
				Symbol{` `, 25, 26, 0, WHITESPACE},
				Symbol{`err`, 26, 29, 0, VARIABLE},
				Symbol{`     `, 29, 34, 0, WHITESPACE},
			},
		},
		symArrayTest{
			Input: `keyValue <- "pi": 3.1419`,
			Expects: []Symbol{
				Symbol{`keyValue`, 0, 8, 0, VARIABLE},
				Symbol{` `, 8, 9, 0, WHITESPACE},
				Symbol{`<-`, 9, 11, 0, ASSIGNMENT},
				Symbol{` `, 11, 12, 0, WHITESPACE},
				Symbol{`"pi"`, 12, 16, 0, STRING},
				Symbol{`:`, 16, 17, 0, UNDEFINED},
				Symbol{` `, 17, 18, 0, WHITESPACE},
				Symbol{`3.1419`, 18, 24, 0, NUMBER},
			},
		},
		symArrayTest{
			Input: `alphabet <- ["a", "b", "c"]`,
			Expects: []Symbol{
				Symbol{`alphabet`, 0, 8, 0, VARIABLE},
				Symbol{` `, 8, 9, 0, WHITESPACE},
				Symbol{`<-`, 9, 11, 0, ASSIGNMENT},
				Symbol{` `, 11, 12, 0, WHITESPACE},
				Symbol{`[`, 12, 13, 0, UNDEFINED},
				Symbol{`"a"`, 13, 16, 0, STRING},
				Symbol{`,`, 16, 17, 0, UNDEFINED},
				Symbol{` `, 17, 18, 0, WHITESPACE},
				Symbol{`"b"`, 18, 21, 0, STRING},
				Symbol{`,`, 21, 22, 0, UNDEFINED},
				Symbol{` `, 22, 23, 0, WHITESPACE},
				Symbol{`"c"`, 23, 26, 0, STRING},
				Symbol{`]`, 26, 27, 0, UNDEFINED},
			},
		},
		symArrayTest{
			Input: `loop i <- 0..5`,
			Expects: []Symbol{
				Symbol{`loop`, 0, 4, 0, KEYWORD_LOOP},
				Symbol{` `, 4, 5, 0, WHITESPACE},
				Symbol{`i`, 5, 6, 0, VARIABLE},
				Symbol{` `, 6, 7, 0, WHITESPACE},
				Symbol{`<-`, 7, 9, 0, ASSIGNMENT},
				Symbol{` `, 9, 10, 0, WHITESPACE},
				Symbol{`0`, 10, 11, 0, NUMBER},
				Symbol{`..`, 11, 13, 0, UNDEFINED},
				Symbol{`5`, 13, 14, 0, NUMBER},
			},
		},
		symArrayTest{
			Input: `x<-2 // The value of x is now 2`,
			Expects: []Symbol{
				Symbol{`x`, 0, 1, 0, VARIABLE},
				Symbol{`<-`, 1, 3, 0, ASSIGNMENT},
				Symbol{`2`, 3, 4, 0, NUMBER},
				Symbol{` `, 4, 5, 0, WHITESPACE},
				Symbol{`// The value of x is now 2`, 5, 31, 0, COMMENT},
			},
		},
		symArrayTest{
			Input: `isLandscape<-length<height`,
			Expects: []Symbol{
				Symbol{`isLandscape`, 0, 11, 0, VARIABLE},
				Symbol{`<-`, 11, 13, 0, ASSIGNMENT},
				Symbol{`length`, 13, 19, 0, VARIABLE},
				Symbol{`<`, 19, 20, 0, LESS_THAN},
				Symbol{`height`, 20, 26, 0, VARIABLE},
			},
		},
		symArrayTest{
			Input: `x<-3.14*(1-2+3)`,
			Expects: []Symbol{
				Symbol{`x`, 0, 1, 0, VARIABLE},
				Symbol{`<-`, 1, 3, 0, ASSIGNMENT},
				Symbol{`3.14`, 3, 7, 0, NUMBER},
				Symbol{`*`, 7, 8, 0, MULTIPLY},
				Symbol{`(`, 8, 9, 0, UNDEFINED},
				Symbol{`1`, 9, 10, 0, NUMBER},
				Symbol{`-`, 10, 11, 0, SUBTRACT},
				Symbol{`2`, 11, 12, 0, NUMBER},
				Symbol{`+`, 12, 13, 0, ADD},
				Symbol{`3`, 13, 14, 0, NUMBER},
				Symbol{`)`, 14, 15, 0, UNDEFINED},
			},
		},
	}
}
