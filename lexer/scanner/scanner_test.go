package scanner

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// scanTest represents a test case for the Scan()
// function.
type scanTest struct {
	TestLine  int
	Input     string
	Expect    []token.Token
	ExpectErr fault.Fault
}

// TestScan runs the test cases for the Scan() function.
func TestScan(t *testing.T) {
	for _, tc := range scanTests() {
		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> scanner_test.go : " + testLine)
		act, err := Scan(tc.Input)

		if tc.ExpectErr != nil {
			assert.Nil(t, act, "Expected token array to be nil")
			require.NotNil(t, err, "Did NOT expect error to be nil")
			assert.NotEmpty(t, err.Error(), "Did NOT expect error message to be empty")
			fault.Assert(t, tc.ExpectErr, err)
		}

		if tc.Expect != nil {
			require.Nil(t, err, "Expected error to be nil")
			assert.Equal(t, tc.Expect, act, "Expected a different token array")
		}
	}
}

// scanTests creates a list of Scan() function tests.
func scanTests() []scanTest {
	return []scanTest{
		scanTest{
			TestLine:  fault.CurrLine(),
			Input:     `x # 1`,
			ExpectErr: fault.Dummy(fault.Symbol).SetFrom(2).SetTo(3),
		},
		scanTest{
			TestLine:  fault.CurrLine(),
			Input:     `123.456.789`,
			ExpectErr: fault.Dummy(fault.Symbol).SetFrom(7).SetTo(8),
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- 1`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 1, 2, 0, token.WHITESPACE},
				token.Token{`<-`, 2, 4, 0, token.ASSIGNMENT},
				token.Token{` `, 4, 5, 0, token.WHITESPACE},
				token.Token{`1`, 5, 6, 0, token.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `y <- -1.1`,
			Expect: []token.Token{
				token.Token{`y`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 1, 2, 0, token.WHITESPACE},
				token.Token{`<-`, 2, 4, 0, token.ASSIGNMENT},
				token.Token{` `, 4, 5, 0, token.WHITESPACE},
				token.Token{`-`, 5, 6, 0, token.CALC_SUBTRACT},
				token.Token{`1.1`, 6, 9, 0, token.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- true`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 1, 2, 0, token.WHITESPACE},
				token.Token{`<-`, 2, 4, 0, token.ASSIGNMENT},
				token.Token{` `, 4, 5, 0, token.WHITESPACE},
				token.Token{`true`, 5, 9, 0, token.BOOLEAN_TRUE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `@Println["Whelp"]`,
			Expect: []token.Token{
				token.Token{`@Println`, 0, 8, 0, token.SOURCERY},
				token.Token{`[`, 8, 9, 0, token.PAREN_SQUARE_OPEN},
				token.Token{`"Whelp"`, 9, 16, 0, token.LITERAL_STRING},
				token.Token{`]`, 16, 17, 0, token.PAREN_SQUARE_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    "\tresult <- func(a, b) r, err     ",
			Expect: []token.Token{
				token.Token{"\t", 0, 1, 0, token.WHITESPACE},
				token.Token{`result`, 1, 7, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 7, 8, 0, token.WHITESPACE},
				token.Token{`<-`, 8, 10, 0, token.ASSIGNMENT},
				token.Token{` `, 10, 11, 0, token.WHITESPACE},
				token.Token{`func`, 11, 15, 0, token.KEYWORD_FUNC},
				token.Token{`(`, 15, 16, 0, token.PAREN_CURVY_OPEN},
				token.Token{`a`, 16, 17, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`,`, 17, 18, 0, token.SEPARATOR_VALUE},
				token.Token{` `, 18, 19, 0, token.WHITESPACE},
				token.Token{`b`, 19, 20, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`)`, 20, 21, 0, token.PAREN_CURVY_CLOSE},
				token.Token{` `, 21, 22, 0, token.WHITESPACE},
				token.Token{`r`, 22, 23, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`,`, 23, 24, 0, token.SEPARATOR_VALUE},
				token.Token{` `, 24, 25, 0, token.WHITESPACE},
				token.Token{`err`, 25, 28, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`     `, 28, 33, 0, token.WHITESPACE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `alphabet <- ["a", "b", "c"]`,
			Expect: []token.Token{
				token.Token{`alphabet`, 0, 8, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 8, 9, 0, token.WHITESPACE},
				token.Token{`<-`, 9, 11, 0, token.ASSIGNMENT},
				token.Token{` `, 11, 12, 0, token.WHITESPACE},
				token.Token{`[`, 12, 13, 0, token.PAREN_SQUARE_OPEN},
				token.Token{`"a"`, 13, 16, 0, token.LITERAL_STRING},
				token.Token{`,`, 16, 17, 0, token.SEPARATOR_VALUE},
				token.Token{` `, 17, 18, 0, token.WHITESPACE},
				token.Token{`"b"`, 18, 21, 0, token.LITERAL_STRING},
				token.Token{`,`, 21, 22, 0, token.SEPARATOR_VALUE},
				token.Token{` `, 22, 23, 0, token.WHITESPACE},
				token.Token{`"c"`, 23, 26, 0, token.LITERAL_STRING},
				token.Token{`]`, 26, 27, 0, token.PAREN_SQUARE_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-2 // The value of x is now 2`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 1, 3, 0, token.ASSIGNMENT},
				token.Token{`2`, 3, 4, 0, token.LITERAL_NUMBER},
				token.Token{` `, 4, 5, 0, token.WHITESPACE},
				token.Token{`// The value of x is now 2`, 5, 31, 0, token.COMMENT},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `isLandscape<-length<height`,
			Expect: []token.Token{
				token.Token{`isLandscape`, 0, 11, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 11, 13, 0, token.ASSIGNMENT},
				token.Token{`length`, 13, 19, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`<`, 19, 20, 0, token.CMP_LESS_THAN},
				token.Token{`height`, 20, 26, 0, token.IDENTIFIER_EXPLICIT},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-3.14*(1-2+3)`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 1, 3, 0, token.ASSIGNMENT},
				token.Token{`3.14`, 3, 7, 0, token.LITERAL_NUMBER},
				token.Token{`*`, 7, 8, 0, token.CALC_MULTIPLY},
				token.Token{`(`, 8, 9, 0, token.PAREN_CURVY_OPEN},
				token.Token{`1`, 9, 10, 0, token.LITERAL_NUMBER},
				token.Token{`-`, 10, 11, 0, token.CALC_SUBTRACT},
				token.Token{`2`, 11, 12, 0, token.LITERAL_NUMBER},
				token.Token{`+`, 12, 13, 0, token.CALC_ADD},
				token.Token{`3`, 13, 14, 0, token.LITERAL_NUMBER},
				token.Token{`)`, 14, 15, 0, token.PAREN_CURVY_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `!x => y <- _`,
			Expect: []token.Token{
				token.Token{`!`, 0, 1, 0, token.LOGICAL_NOT},
				token.Token{`x`, 1, 2, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 2, 3, 0, token.WHITESPACE},
				token.Token{`=>`, 3, 5, 0, token.LOGICAL_MATCH},
				token.Token{` `, 5, 6, 0, token.WHITESPACE},
				token.Token{`y`, 6, 7, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 7, 8, 0, token.WHITESPACE},
				token.Token{`<-`, 8, 10, 0, token.ASSIGNMENT},
				token.Token{` `, 10, 11, 0, token.WHITESPACE},
				token.Token{`_`, 11, 12, 0, token.VOID},
			},
		},
	}
}
