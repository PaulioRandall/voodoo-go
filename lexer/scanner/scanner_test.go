package scanner

import (
	"strconv"
	"testing"

	fault "github.com/PaulioRandall/voodoo-go/new_fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// scanTest represents a test case for the Scan()
// function.
type scanTest struct {
	TestLine int
	Input    string
	Expect   []token.Token
	Error    fault.Fault
}

// TestScan runs the test cases for the Scan() function.
func TestScan(t *testing.T) {
	for _, tc := range scanTests() {
		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> scanner_test.go : " + testLine)
		act, err := Scan(tc.Input)

		if tc.Error != nil {
			assert.Nil(t, act, "Expected token to be nil")
			require.NotNil(t, err, "Did NOT expect error to be nil")
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
			TestLine: fault.CurrLine(),
			Input:    `x # 1`,
			Error:    newFault(3),
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `123.456.789`,
			Error:    newFault(8),
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- 1`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 1, 2, token.WHITESPACE},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `y <- -1.1`,
			Expect: []token.Token{
				token.Token{`y`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 1, 2, token.WHITESPACE},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`-`, 5, 6, token.CALC_SUBTRACT},
				token.Token{`1.1`, 6, 9, token.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- true`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 1, 2, token.WHITESPACE},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`true`, 5, 9, token.BOOLEAN_TRUE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `@Println["Whelp"]`,
			Expect: []token.Token{
				token.Token{`@Println`, 0, 8, token.SPELL},
				token.Token{`[`, 8, 9, token.PAREN_SQUARE_OPEN},
				token.Token{`"Whelp"`, 9, 16, token.LITERAL_STRING},
				token.Token{`]`, 16, 17, token.PAREN_SQUARE_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    "\tresult <- func(a, b) r, err     ",
			Expect: []token.Token{
				token.Token{"\t", 0, 1, token.WHITESPACE},
				token.Token{`result`, 1, 7, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 7, 8, token.WHITESPACE},
				token.Token{`<-`, 8, 10, token.ASSIGNMENT},
				token.Token{` `, 10, 11, token.WHITESPACE},
				token.Token{`func`, 11, 15, token.KEYWORD_FUNC},
				token.Token{`(`, 15, 16, token.PAREN_CURVY_OPEN},
				token.Token{`a`, 16, 17, token.IDENTIFIER_EXPLICIT},
				token.Token{`,`, 17, 18, token.SEPARATOR_VALUE},
				token.Token{` `, 18, 19, token.WHITESPACE},
				token.Token{`b`, 19, 20, token.IDENTIFIER_EXPLICIT},
				token.Token{`)`, 20, 21, token.PAREN_CURVY_CLOSE},
				token.Token{` `, 21, 22, token.WHITESPACE},
				token.Token{`r`, 22, 23, token.IDENTIFIER_EXPLICIT},
				token.Token{`,`, 23, 24, token.SEPARATOR_VALUE},
				token.Token{` `, 24, 25, token.WHITESPACE},
				token.Token{`err`, 25, 28, token.IDENTIFIER_EXPLICIT},
				token.Token{`     `, 28, 33, token.WHITESPACE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `alphabet <- ["a", "b", "c"]`,
			Expect: []token.Token{
				token.Token{`alphabet`, 0, 8, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 8, 9, token.WHITESPACE},
				token.Token{`<-`, 9, 11, token.ASSIGNMENT},
				token.Token{` `, 11, 12, token.WHITESPACE},
				token.Token{`[`, 12, 13, token.PAREN_SQUARE_OPEN},
				token.Token{`"a"`, 13, 16, token.LITERAL_STRING},
				token.Token{`,`, 16, 17, token.SEPARATOR_VALUE},
				token.Token{` `, 17, 18, token.WHITESPACE},
				token.Token{`"b"`, 18, 21, token.LITERAL_STRING},
				token.Token{`,`, 21, 22, token.SEPARATOR_VALUE},
				token.Token{` `, 22, 23, token.WHITESPACE},
				token.Token{`"c"`, 23, 26, token.LITERAL_STRING},
				token.Token{`]`, 26, 27, token.PAREN_SQUARE_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-2 // The value of x is now 2`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 1, 3, token.ASSIGNMENT},
				token.Token{`2`, 3, 4, token.LITERAL_NUMBER},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`// The value of x is now 2`, 5, 31, token.COMMENT},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `isLandscape<-length<height`,
			Expect: []token.Token{
				token.Token{`isLandscape`, 0, 11, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 11, 13, token.ASSIGNMENT},
				token.Token{`length`, 13, 19, token.IDENTIFIER_EXPLICIT},
				token.Token{`<`, 19, 20, token.CMP_LESS_THAN},
				token.Token{`height`, 20, 26, token.IDENTIFIER_EXPLICIT},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-3.14*(1-2+3)`,
			Expect: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 1, 3, token.ASSIGNMENT},
				token.Token{`3.14`, 3, 7, token.LITERAL_NUMBER},
				token.Token{`*`, 7, 8, token.CALC_MULTIPLY},
				token.Token{`(`, 8, 9, token.PAREN_CURVY_OPEN},
				token.Token{`1`, 9, 10, token.LITERAL_NUMBER},
				token.Token{`-`, 10, 11, token.CALC_SUBTRACT},
				token.Token{`2`, 11, 12, token.LITERAL_NUMBER},
				token.Token{`+`, 12, 13, token.CALC_ADD},
				token.Token{`3`, 13, 14, token.LITERAL_NUMBER},
				token.Token{`)`, 14, 15, token.PAREN_CURVY_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `!x => y <- _`,
			Expect: []token.Token{
				token.Token{`!`, 0, 1, token.LOGICAL_NOT},
				token.Token{`x`, 1, 2, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 2, 3, token.WHITESPACE},
				token.Token{`=>`, 3, 5, token.LOGICAL_MATCH},
				token.Token{` `, 5, 6, token.WHITESPACE},
				token.Token{`y`, 6, 7, token.IDENTIFIER_EXPLICIT},
				token.Token{` `, 7, 8, token.WHITESPACE},
				token.Token{`<-`, 8, 10, token.ASSIGNMENT},
				token.Token{` `, 10, 11, token.WHITESPACE},
				token.Token{`_`, 11, 12, token.VOID},
			},
		},
	}
}
