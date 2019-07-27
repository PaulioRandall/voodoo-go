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

func TestScannerApi(t *testing.T) {
	for _, tc := range apiTests() {
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

type lexScanFunc func(*runer.RuneItr) *symbol.Lexeme

func lexFuncTest(t *testing.T, fName string, f lexScanFunc, tests []lexTest) {
	for _, tc := range tests {
		require.NotNil(t, tc.Expect)
		require.Nil(t, tc.ExpectErr)

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fName + " : " + testLine)

		itr := runer.NewRuneItr(tc.Input)
		act := f(itr)

		require.NotNil(t, act)
		assert.Equal(t, tc.Expect, *act)
	}
}

type lexScanErrFunc func(*runer.RuneItr) (*symbol.Lexeme, fault.Fault)

func lexErrFuncTest(t *testing.T, fName string, f lexScanErrFunc, tests []lexTest) {
	for _, tc := range tests {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fName + " : " + testLine)

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

type lexTest struct {
	TestLine  int
	Line      int
	Input     string
	Expect    symbol.Lexeme
	ExpectErr fault.Fault
}

type scanLineTest struct {
	TestLine  int
	Input     string
	Expect    []symbol.Lexeme
	ExpectErr fault.Fault
}

func apiTests() []scanLineTest {
	return []scanLineTest{
		scanLineTest{
			TestLine:  fault.CurrLine(),
			Input:     `x # 1`,
			ExpectErr: fault.Dummy(fault.Symbol).Line(0).From(2).To(3),
		},
		scanLineTest{
			TestLine:  fault.CurrLine(),
			Input:     `123.456.789`,
			ExpectErr: fault.Dummy(fault.Number).Line(0).From(7),
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- 1`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`1`, 5, 6, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `y <- -1.1`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`y`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`-`, 5, 6, 0, symbol.CALC_SUBTRACT},
				symbol.Lexeme{`1.1`, 6, 9, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- true`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`true`, 5, 9, 0, symbol.BOOLEAN_TRUE},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `@Println["Whelp"]`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`@Println`, 0, 8, 0, symbol.SOURCERY},
				symbol.Lexeme{`[`, 8, 9, 0, symbol.PAREN_SQUARE_OPEN},
				symbol.Lexeme{`"Whelp"`, 9, 16, 0, symbol.LITERAL_STRING},
				symbol.Lexeme{`]`, 16, 17, 0, symbol.PAREN_SQUARE_CLOSE},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    "\tresult <- func(a, b) r, err     ",
			Expect: []symbol.Lexeme{
				symbol.Lexeme{"\t", 0, 1, 0, symbol.WHITESPACE},
				symbol.Lexeme{`result`, 1, 7, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Lexeme{`func`, 11, 15, 0, symbol.KEYWORD_FUNC},
				symbol.Lexeme{`(`, 15, 16, 0, symbol.PAREN_CURVY_OPEN},
				symbol.Lexeme{`a`, 16, 17, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`,`, 17, 18, 0, symbol.SEPARATOR_VALUE},
				symbol.Lexeme{` `, 18, 19, 0, symbol.WHITESPACE},
				symbol.Lexeme{`b`, 19, 20, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`)`, 20, 21, 0, symbol.PAREN_CURVY_CLOSE},
				symbol.Lexeme{` `, 21, 22, 0, symbol.WHITESPACE},
				symbol.Lexeme{`r`, 22, 23, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`,`, 23, 24, 0, symbol.SEPARATOR_VALUE},
				symbol.Lexeme{` `, 24, 25, 0, symbol.WHITESPACE},
				symbol.Lexeme{`err`, 25, 28, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`     `, 28, 33, 0, symbol.WHITESPACE},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `keyValue <- "pi": 3.1419`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`keyValue`, 0, 8, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Lexeme{`"pi"`, 12, 16, 0, symbol.LITERAL_STRING},
				symbol.Lexeme{`:`, 16, 17, 0, symbol.SEPARATOR_KEY_VALUE},
				symbol.Lexeme{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Lexeme{`3.1419`, 18, 24, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `alphabet <- ["a", "b", "c"]`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`alphabet`, 0, 8, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Lexeme{`[`, 12, 13, 0, symbol.PAREN_SQUARE_OPEN},
				symbol.Lexeme{`"a"`, 13, 16, 0, symbol.LITERAL_STRING},
				symbol.Lexeme{`,`, 16, 17, 0, symbol.SEPARATOR_VALUE},
				symbol.Lexeme{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Lexeme{`"b"`, 18, 21, 0, symbol.LITERAL_STRING},
				symbol.Lexeme{`,`, 21, 22, 0, symbol.SEPARATOR_VALUE},
				symbol.Lexeme{` `, 22, 23, 0, symbol.WHITESPACE},
				symbol.Lexeme{`"c"`, 23, 26, 0, symbol.LITERAL_STRING},
				symbol.Lexeme{`]`, 26, 27, 0, symbol.PAREN_SQUARE_CLOSE},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `loop i <- 0..5`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`i`, 5, 6, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 7, 9, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 9, 10, 0, symbol.WHITESPACE},
				symbol.Lexeme{`0`, 10, 11, 0, symbol.LITERAL_NUMBER},
				symbol.Lexeme{`..`, 11, 13, 0, symbol.RANGE},
				symbol.Lexeme{`5`, 13, 14, 0, symbol.LITERAL_NUMBER},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-2 // The value of x is now 2`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{`2`, 3, 4, 0, symbol.LITERAL_NUMBER},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`// The value of x is now 2`, 5, 31, 0, symbol.COMMENT},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `isLandscape<-length<height`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`isLandscape`, 0, 11, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`<-`, 11, 13, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{`length`, 13, 19, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`<`, 19, 20, 0, symbol.CMP_LESS_THAN},
				symbol.Lexeme{`height`, 20, 26, 0, symbol.IDENTIFIER_EXPLICIT},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-3.14*(1-2+3)`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{`3.14`, 3, 7, 0, symbol.LITERAL_NUMBER},
				symbol.Lexeme{`*`, 7, 8, 0, symbol.CALC_MULTIPLY},
				symbol.Lexeme{`(`, 8, 9, 0, symbol.PAREN_CURVY_OPEN},
				symbol.Lexeme{`1`, 9, 10, 0, symbol.LITERAL_NUMBER},
				symbol.Lexeme{`-`, 10, 11, 0, symbol.CALC_SUBTRACT},
				symbol.Lexeme{`2`, 11, 12, 0, symbol.LITERAL_NUMBER},
				symbol.Lexeme{`+`, 12, 13, 0, symbol.CALC_ADD},
				symbol.Lexeme{`3`, 13, 14, 0, symbol.LITERAL_NUMBER},
				symbol.Lexeme{`)`, 14, 15, 0, symbol.PAREN_CURVY_CLOSE},
			},
		},
		scanLineTest{
			TestLine: fault.CurrLine(),
			Input:    `!x => y <- _`,
			Expect: []symbol.Lexeme{
				symbol.Lexeme{`!`, 0, 1, 0, symbol.LOGICAL_NOT},
				symbol.Lexeme{`x`, 1, 2, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 2, 3, 0, symbol.WHITESPACE},
				symbol.Lexeme{`=>`, 3, 5, 0, symbol.LOGICAL_MATCH},
				symbol.Lexeme{` `, 5, 6, 0, symbol.WHITESPACE},
				symbol.Lexeme{`y`, 6, 7, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Lexeme{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Lexeme{`_`, 11, 12, 0, symbol.VOID},
			},
		},
	}
}
