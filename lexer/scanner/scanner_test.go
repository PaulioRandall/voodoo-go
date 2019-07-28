package scanner

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// scanFuncTest represents a test case for any of the
// delegate scanning functions.
type scanFuncTest struct {
	TestLine  int
	Input     string
	Expect    symbol.Token
	ExpectErr fault.Fault
}

// runScanTest runs the input test cases on the input
// function.
func runScanTest(
	t *testing.T,
	fileName string,
	f func(*runer.RuneItr) *symbol.Token,
	tests []scanFuncTest) {

	for _, tc := range tests {
		require.NotNil(t, tc.Expect)
		require.Nil(t, tc.ExpectErr)

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		itr := runer.NewRuneItr(tc.Input)
		act := f(itr)

		require.NotNil(t, act)
		assert.Equal(t, tc.Expect, *act)
	}
}

// runFailableScanTest runs the input test cases on the
// input function.
func runFailableScanTest(
	t *testing.T,
	fileName string,
	f func(*runer.RuneItr) (*symbol.Token, fault.Fault),
	tests []scanFuncTest) {

	for _, tc := range tests {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		itr := runer.NewRuneItr(tc.Input)
		act, err := f(itr)

		if tc.ExpectErr != nil {
			assert.Nil(t, act)
			require.NotNil(t, err)
			assert.NotEmpty(t, err.Error())
			fault.Assert(t, tc.ExpectErr, err)

		} else {
			assert.Nil(t, err)
			require.NotNil(t, act)
			assert.Equal(t, tc.Expect, *act)
		}
	}
}

// scanTest represents a test case for the Scan()
// function.
type scanTest struct {
	TestLine  int
	Input     string
	Expect    []symbol.Token
	ExpectErr fault.Fault
}

// TestScan runs the test cases for the Scan() function.
func TestScan(t *testing.T) {
	for _, tc := range scanTests() {
		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> scanner_test.go : " + testLine)
		act, err := Scan(tc.Input)

		if tc.ExpectErr != nil {
			assert.Nil(t, act)
			require.NotNil(t, err)
			assert.NotEmpty(t, err.Error())
			fault.Assert(t, tc.ExpectErr, err)
		}

		if tc.Expect != nil {
			require.Nil(t, err)
			assert.Equal(t, tc.Expect, act)
		}
	}
}

// scanTests creates a list of Scan() function tests.
func scanTests() []scanTest {
	return []scanTest{
		scanTest{
			TestLine:  fault.CurrLine(),
			Input:     `x # 1`,
			ExpectErr: fault.Dummy(fault.Symbol).Line(0).From(2).To(3),
		},
		scanTest{
			TestLine:  fault.CurrLine(),
			Input:     `123.456.789`,
			ExpectErr: fault.Dummy(fault.Number).Line(0).From(7),
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- 1`,
			Expect: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Token{`1`, 5, 6, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `y <- -1.1`,
			Expect: []symbol.Token{
				symbol.Token{`y`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Token{`-`, 5, 6, 0, symbol.CALC_SUBTRACT},
				symbol.Token{`1.1`, 6, 9, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- true`,
			Expect: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Token{`true`, 5, 9, 0, symbol.BOOLEAN_TRUE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `@Println["Whelp"]`,
			Expect: []symbol.Token{
				symbol.Token{`@Println`, 0, 8, 0, symbol.SOURCERY},
				symbol.Token{`[`, 8, 9, 0, symbol.PAREN_SQUARE_OPEN},
				symbol.Token{`"Whelp"`, 9, 16, 0, symbol.LITERAL_STRING},
				symbol.Token{`]`, 16, 17, 0, symbol.PAREN_SQUARE_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    "\tresult <- func(a, b) r, err     ",
			Expect: []symbol.Token{
				symbol.Token{"\t", 0, 1, 0, symbol.WHITESPACE},
				symbol.Token{`result`, 1, 7, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Token{`func`, 11, 15, 0, symbol.KEYWORD_FUNC},
				symbol.Token{`(`, 15, 16, 0, symbol.PAREN_CURVY_OPEN},
				symbol.Token{`a`, 16, 17, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`,`, 17, 18, 0, symbol.SEPARATOR_VALUE},
				symbol.Token{` `, 18, 19, 0, symbol.WHITESPACE},
				symbol.Token{`b`, 19, 20, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`)`, 20, 21, 0, symbol.PAREN_CURVY_CLOSE},
				symbol.Token{` `, 21, 22, 0, symbol.WHITESPACE},
				symbol.Token{`r`, 22, 23, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`,`, 23, 24, 0, symbol.SEPARATOR_VALUE},
				symbol.Token{` `, 24, 25, 0, symbol.WHITESPACE},
				symbol.Token{`err`, 25, 28, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`     `, 28, 33, 0, symbol.WHITESPACE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `keyValue <- "pi": 3.1419`,
			Expect: []symbol.Token{
				symbol.Token{`keyValue`, 0, 8, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Token{`"pi"`, 12, 16, 0, symbol.LITERAL_STRING},
				symbol.Token{`:`, 16, 17, 0, symbol.SEPARATOR_KEY_VALUE},
				symbol.Token{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Token{`3.1419`, 18, 24, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `alphabet <- ["a", "b", "c"]`,
			Expect: []symbol.Token{
				symbol.Token{`alphabet`, 0, 8, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Token{`[`, 12, 13, 0, symbol.PAREN_SQUARE_OPEN},
				symbol.Token{`"a"`, 13, 16, 0, symbol.LITERAL_STRING},
				symbol.Token{`,`, 16, 17, 0, symbol.SEPARATOR_VALUE},
				symbol.Token{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Token{`"b"`, 18, 21, 0, symbol.LITERAL_STRING},
				symbol.Token{`,`, 21, 22, 0, symbol.SEPARATOR_VALUE},
				symbol.Token{` `, 22, 23, 0, symbol.WHITESPACE},
				symbol.Token{`"c"`, 23, 26, 0, symbol.LITERAL_STRING},
				symbol.Token{`]`, 26, 27, 0, symbol.PAREN_SQUARE_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `loop i <- 0..5`,
			Expect: []symbol.Token{
				symbol.Token{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
				symbol.Token{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Token{`i`, 5, 6, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 7, 9, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 9, 10, 0, symbol.WHITESPACE},
				symbol.Token{`0`, 10, 11, 0, symbol.LITERAL_NUMBER},
				symbol.Token{`..`, 11, 13, 0, symbol.RANGE},
				symbol.Token{`5`, 13, 14, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-2 // The value of x is now 2`,
			Expect: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Token{`2`, 3, 4, 0, symbol.LITERAL_NUMBER},
				symbol.Token{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Token{`// The value of x is now 2`, 5, 31, 0, symbol.COMMENT},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `isLandscape<-length<height`,
			Expect: []symbol.Token{
				symbol.Token{`isLandscape`, 0, 11, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`<-`, 11, 13, 0, symbol.ASSIGNMENT},
				symbol.Token{`length`, 13, 19, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`<`, 19, 20, 0, symbol.CMP_LESS_THAN},
				symbol.Token{`height`, 20, 26, 0, symbol.IDENTIFIER_EXPLICIT},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-3.14*(1-2+3)`,
			Expect: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Token{`3.14`, 3, 7, 0, symbol.LITERAL_NUMBER},
				symbol.Token{`*`, 7, 8, 0, symbol.CALC_MULTIPLY},
				symbol.Token{`(`, 8, 9, 0, symbol.PAREN_CURVY_OPEN},
				symbol.Token{`1`, 9, 10, 0, symbol.LITERAL_NUMBER},
				symbol.Token{`-`, 10, 11, 0, symbol.CALC_SUBTRACT},
				symbol.Token{`2`, 11, 12, 0, symbol.LITERAL_NUMBER},
				symbol.Token{`+`, 12, 13, 0, symbol.CALC_ADD},
				symbol.Token{`3`, 13, 14, 0, symbol.LITERAL_NUMBER},
				symbol.Token{`)`, 14, 15, 0, symbol.PAREN_CURVY_CLOSE},
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `!x => y <- _`,
			Expect: []symbol.Token{
				symbol.Token{`!`, 0, 1, 0, symbol.LOGICAL_NOT},
				symbol.Token{`x`, 1, 2, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 2, 3, 0, symbol.WHITESPACE},
				symbol.Token{`=>`, 3, 5, 0, symbol.LOGICAL_MATCH},
				symbol.Token{` `, 5, 6, 0, symbol.WHITESPACE},
				symbol.Token{`y`, 6, 7, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Token{`_`, 11, 12, 0, symbol.VOID},
			},
		},
	}
}
