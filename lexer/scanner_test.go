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

type symScanFunc func(*runer.RuneItr) *lexeme.Lexeme

func symFuncTest(t *testing.T, fName string, f symScanFunc, tests []lexTest) {
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

type symScanErrFunc func(*runer.RuneItr) (*lexeme.Lexeme, LexError)

func symErrFuncTest(t *testing.T, fName string, f symScanErrFunc, tests []lexTest) {
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

type lexTest struct {
	Line      int
	Input     string
	ExpectSym lexeme.Lexeme
	ExpectErr LexError
}

type scanLineTest struct {
	Line       int
	Input      string
	ExpectSyms []lexeme.Lexeme
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
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`1`, 5, 6, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Input: `y <- -1.1`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`y`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`-`, 5, 6, 0, lexeme.SUBTRACT},
				lexeme.Lexeme{`1.1`, 6, 9, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Line:  123,
			Input: `x <- true`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 123, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 1, 2, 123, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 2, 4, 123, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 4, 5, 123, lexeme.WHITESPACE},
				lexeme.Lexeme{`true`, 5, 9, 123, lexeme.BOOLEAN},
			},
		},
		scanLineTest{
			Input: `@Println["Whelp"]`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`@Println`, 0, 8, 0, lexeme.SOURCERY},
				lexeme.Lexeme{`[`, 8, 9, 0, lexeme.SQUARE_BRACE_OPEN},
				lexeme.Lexeme{`"Whelp"`, 9, 16, 0, lexeme.STRING},
				lexeme.Lexeme{`]`, 16, 17, 0, lexeme.SQUARE_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: "\tresult <- spell(a, b) r, err     ",
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{"\t", 0, 1, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`result`, 1, 7, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 7, 8, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 8, 10, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 10, 11, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`spell`, 11, 16, 0, lexeme.KEYWORD_SPELL},
				lexeme.Lexeme{`(`, 16, 17, 0, lexeme.CURVED_BRACE_OPEN},
				lexeme.Lexeme{`a`, 17, 18, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`,`, 18, 19, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Lexeme{` `, 19, 20, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`b`, 20, 21, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`)`, 21, 22, 0, lexeme.CURVED_BRACE_CLOSE},
				lexeme.Lexeme{` `, 22, 23, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`r`, 23, 24, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`,`, 24, 25, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Lexeme{` `, 25, 26, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`err`, 26, 29, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`     `, 29, 34, 0, lexeme.WHITESPACE},
			},
		},
		scanLineTest{
			Input: `keyValue <- "pi": 3.1419`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`keyValue`, 0, 8, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 8, 9, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 9, 11, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 11, 12, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`"pi"`, 12, 16, 0, lexeme.STRING},
				lexeme.Lexeme{`:`, 16, 17, 0, lexeme.KEY_VALUE_SEPARATOR},
				lexeme.Lexeme{` `, 17, 18, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`3.1419`, 18, 24, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Input: `alphabet <- ["a", "b", "c"]`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`alphabet`, 0, 8, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 8, 9, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 9, 11, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 11, 12, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`[`, 12, 13, 0, lexeme.SQUARE_BRACE_OPEN},
				lexeme.Lexeme{`"a"`, 13, 16, 0, lexeme.STRING},
				lexeme.Lexeme{`,`, 16, 17, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Lexeme{` `, 17, 18, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`"b"`, 18, 21, 0, lexeme.STRING},
				lexeme.Lexeme{`,`, 21, 22, 0, lexeme.VALUE_SEPARATOR},
				lexeme.Lexeme{` `, 22, 23, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`"c"`, 23, 26, 0, lexeme.STRING},
				lexeme.Lexeme{`]`, 26, 27, 0, lexeme.SQUARE_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: `loop i <- 0..5`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`loop`, 0, 4, 0, lexeme.KEYWORD_LOOP},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`i`, 5, 6, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 6, 7, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 7, 9, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 9, 10, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`0`, 10, 11, 0, lexeme.NUMBER},
				lexeme.Lexeme{`..`, 11, 13, 0, lexeme.RANGE},
				lexeme.Lexeme{`5`, 13, 14, 0, lexeme.NUMBER},
			},
		},
		scanLineTest{
			Input: `x<-2 // The value of x is now 2`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`<-`, 1, 3, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{`2`, 3, 4, 0, lexeme.NUMBER},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`// The value of x is now 2`, 5, 31, 0, lexeme.COMMENT},
			},
		},
		scanLineTest{
			Input: `isLandscape<-length<height`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`isLandscape`, 0, 11, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`<-`, 11, 13, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{`length`, 13, 19, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`<`, 19, 20, 0, lexeme.LESS_THAN},
				lexeme.Lexeme{`height`, 20, 26, 0, lexeme.VARIABLE},
			},
		},
		scanLineTest{
			Input: `x<-3.14*(1-2+3)`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 0, lexeme.VARIABLE},
				lexeme.Lexeme{`<-`, 1, 3, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{`3.14`, 3, 7, 0, lexeme.NUMBER},
				lexeme.Lexeme{`*`, 7, 8, 0, lexeme.MULTIPLY},
				lexeme.Lexeme{`(`, 8, 9, 0, lexeme.CURVED_BRACE_OPEN},
				lexeme.Lexeme{`1`, 9, 10, 0, lexeme.NUMBER},
				lexeme.Lexeme{`-`, 10, 11, 0, lexeme.SUBTRACT},
				lexeme.Lexeme{`2`, 11, 12, 0, lexeme.NUMBER},
				lexeme.Lexeme{`+`, 12, 13, 0, lexeme.ADD},
				lexeme.Lexeme{`3`, 13, 14, 0, lexeme.NUMBER},
				lexeme.Lexeme{`)`, 14, 15, 0, lexeme.CURVED_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: `!x => y <- _`,
			ExpectSyms: []lexeme.Lexeme{
				lexeme.Lexeme{`!`, 0, 1, 0, lexeme.NEGATION},
				lexeme.Lexeme{`x`, 1, 2, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 2, 3, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`=>`, 3, 5, 0, lexeme.IF_TRUE_THEN},
				lexeme.Lexeme{` `, 5, 6, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`y`, 6, 7, 0, lexeme.VARIABLE},
				lexeme.Lexeme{` `, 7, 8, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 8, 10, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 10, 11, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`_`, 11, 12, 0, lexeme.VOID},
			},
		},
	}
}
