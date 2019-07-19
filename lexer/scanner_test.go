package lexer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScannerApi(t *testing.T) {
	for i, tc := range apiTests() {
		t.Log("Scanner test case: " + strconv.Itoa(i+1))
		s, err := ScanLine(tc.Input, tc.Line)

		if tc.ExpectErr != nil {
			require.Nil(t, s)
			assert.NotEmpty(t, err.Error())
			assert.Equal(t, tc.ExpectErr.Line(), err.Line())
			assert.Equal(t, tc.ExpectErr.Col(), err.Col())
		}

		if tc.ExpectSyms != nil {
			require.Nil(t, err)
			assert.Equal(t, tc.ExpectSyms, s)
		}
	}
}

type symScanFunc func(*runer.RuneItr) *lexeme.Symbol

func symFuncTest(t *testing.T, fName string, f symScanFunc, tests []symTest) {
	for i, tc := range tests {
		require.NotNil(t, tc.ExpectSym)
		require.Nil(t, tc.ExpectErr)

		t.Log(fName + "() test case: " + strconv.Itoa(i+1))

		itr := runer.NewRuneItr(tc.Input)
		s := f(itr)

		require.NotNil(t, s)
		assert.Equal(t, tc.ExpectSym, *s)
	}
}

type symScanErrFunc func(*runer.RuneItr) (*lexeme.Symbol, LexError)

func symErrFuncTest(t *testing.T, fName string, f symScanErrFunc, tests []symTest) {
	for i, tc := range tests {
		t.Log(fName + "() test case: " + strconv.Itoa(i+1))

		itr := runer.NewRuneItr(tc.Input)
		s, err := f(itr)

		if tc.ExpectErr != nil {
			assert.NotNil(t, err)
			require.Nil(t, s)
			assert.NotEmpty(t, err.Error())
			assert.Equal(t, tc.ExpectErr.Line(), err.Line())
			assert.Equal(t, tc.ExpectErr.Col(), err.Col())
		} else {
			assert.Nil(t, err)
			require.NotNil(t, s)
			assert.Equal(t, tc.ExpectSym, *s)
		}
	}
}

type symTest struct {
	Line      int
	Input     string
	ExpectSym lexeme.Symbol
	ExpectErr LexError
}

type scanLineTest struct {
	Line       int
	Input      string
	ExpectSyms []lexeme.Symbol
	ExpectErr  LexError
}

type expLexError struct {
	line int // Line number
	col  int // Column number
}

func (e expLexError) Error() string {
	// Error messages should be semantically validated
	// so this is not required for testing.
	return ""
}

func (e expLexError) Line() int {
	return e.line
}

func (e expLexError) Col() int {
	return e.col
}

func apiTests() []scanLineTest {
	return []scanLineTest{
		scanLineTest{
			Input:     `x # 1`,
			ExpectErr: expLexError{0, 3},
		},
		scanLineTest{
			Input:     `123.456.789`,
			ExpectErr: expLexError{0, 7},
		},
		scanLineTest{
			Input: `x <- 1`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`x`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`1`, 5, 6, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Input: `y <- -1.1`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`y`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`-`, 5, 6, 0, lexeme.SUBTRACT},
				lexeme.Symbol{`1.1`, 6, 9, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Line:  123,
			Input: `x <- true`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`x`, 0, 1, 123, lexeme.VARIABLE},
				lexeme.Symbol{` `, 1, 2, 123, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 2, 4, 123, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 4, 5, 123, lexeme.WHITESPACE},
				lexeme.Symbol{`true`, 5, 9, 123, lexeme.BOOLEAN},
			},
		},
		scanLineTest{
			Input: `@Println["Whelp"]`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`@Println`, 0, 8, 0, lexeme.SOURCERY},
				lexeme.Symbol{`[`, 8, 9, 0, lexeme.SQUARE_BRACE_OPEN},
				lexeme.Symbol{`"Whelp"`, 9, 16, 0, lexeme.STRING},
				lexeme.Symbol{`]`, 16, 17, 0, lexeme.SQUARE_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: "\tresult <- spell(a, b) r, err     ",
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{"\t", 0, 1, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`result`, 1, 7, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 7, 8, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 8, 10, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 10, 11, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`spell`, 11, 16, 0, lexeme.KEYWORD_SPELL},
				lexeme.Symbol{`(`, 16, 17, 0, lexeme.CURVED_BRACE_OPEN},
				lexeme.Symbol{`a`, 17, 18, 0, lexeme.VARIABLE},
				lexeme.Symbol{`,`, 18, 19, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Symbol{` `, 19, 20, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`b`, 20, 21, 0, lexeme.VARIABLE},
				lexeme.Symbol{`)`, 21, 22, 0, lexeme.CURVED_BRACE_CLOSE},
				lexeme.Symbol{` `, 22, 23, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`r`, 23, 24, 0, lexeme.VARIABLE},
				lexeme.Symbol{`,`, 24, 25, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Symbol{` `, 25, 26, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`err`, 26, 29, 0, lexeme.VARIABLE},
				lexeme.Symbol{`     `, 29, 34, 0, lexeme.WHITESPACE},
			},
		},
		scanLineTest{
			Input: `keyValue <- "pi": 3.1419`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`keyValue`, 0, 8, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 8, 9, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 9, 11, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 11, 12, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`"pi"`, 12, 16, 0, lexeme.STRING},
				lexeme.Symbol{`:`, 16, 17, 0, lexeme.KEY_VALUE_SEPARATOR},
				lexeme.Symbol{` `, 17, 18, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`3.1419`, 18, 24, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Input: `alphabet <- ["a", "b", "c"]`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`alphabet`, 0, 8, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 8, 9, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 9, 11, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 11, 12, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`[`, 12, 13, 0, lexeme.SQUARE_BRACE_OPEN},
				lexeme.Symbol{`"a"`, 13, 16, 0, lexeme.STRING},
				lexeme.Symbol{`,`, 16, 17, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Symbol{` `, 17, 18, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`"b"`, 18, 21, 0, lexeme.STRING},
				lexeme.Symbol{`,`, 21, 22, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Symbol{` `, 22, 23, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`"c"`, 23, 26, 0, lexeme.STRING},
				lexeme.Symbol{`]`, 26, 27, 0, lexeme.SQUARE_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: `loop i <- 0..5`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`loop`, 0, 4, 0, lexeme.KEYWORD_LOOP},
				lexeme.Symbol{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`i`, 5, 6, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 6, 7, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 7, 9, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 9, 10, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`0`, 10, 11, 0, lexeme.NUMBER},
				lexeme.Symbol{`..`, 11, 13, 0, lexeme.RANGE},
				lexeme.Symbol{`5`, 13, 14, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Input: `x<-2 // The value of x is now 2`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`x`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Symbol{`<-`, 1, 3, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{`2`, 3, 4, 0, lexeme.NUMBER},
				lexeme.Symbol{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`// The value of x is now 2`, 5, 31, 0, lexeme.COMMENT},
			},
		},
		scanLineTest{
			Input: `isLandscape<-length<height`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`isLandscape`, 0, 11, 0, lexeme.VARIABLE},
				lexeme.Symbol{`<-`, 11, 13, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{`length`, 13, 19, 0, lexeme.VARIABLE},
				lexeme.Symbol{`<`, 19, 20, 0, lexeme.LESS_THAN},
				lexeme.Symbol{`height`, 20, 26, 0, lexeme.VARIABLE},
			},
		},
		scanLineTest{
			Input: `x<-3.14*(1-2+3)`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`x`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Symbol{`<-`, 1, 3, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{`3.14`, 3, 7, 0, lexeme.NUMBER},
				lexeme.Symbol{`*`, 7, 8, 0, lexeme.MULTIPLY},
				lexeme.Symbol{`(`, 8, 9, 0, lexeme.CURVED_BRACE_OPEN},
				lexeme.Symbol{`1`, 9, 10, 0, lexeme.NUMBER},
				lexeme.Symbol{`-`, 10, 11, 0, lexeme.SUBTRACT},
				lexeme.Symbol{`2`, 11, 12, 0, lexeme.NUMBER},
				lexeme.Symbol{`+`, 12, 13, 0, lexeme.ADD},
				lexeme.Symbol{`3`, 13, 14, 0, lexeme.NUMBER},
				lexeme.Symbol{`)`, 14, 15, 0, lexeme.CURVED_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: `!x => y <- _`,
			ExpectSyms: []lexeme.Symbol{
				lexeme.Symbol{`!`, 0, 1, 0, lexeme.NEGATION},
				lexeme.Symbol{`x`, 1, 2, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 2, 3, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`=>`, 3, 5, 0, lexeme.IF_TRUE_THEN},
				lexeme.Symbol{` `, 5, 6, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`y`, 6, 7, 0, lexeme.VARIABLE},
				lexeme.Symbol{` `, 7, 8, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`<-`, 8, 10, 0, lexeme.ASSIGNMENT},
				lexeme.Symbol{` `, 10, 11, 0, lexeme.WHITESPACE},
				lexeme.Symbol{`_`, 11, 12, 0, lexeme.VOID},
			},
		},
	}
}
