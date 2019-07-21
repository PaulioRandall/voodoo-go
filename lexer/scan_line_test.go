package lexer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScannerApi(t *testing.T) {
	for i, tc := range apiTests() {
		t.Log("ScanLine() test case: " + strconv.Itoa(i+1))
		ls, err := ScanLine(tc.Input, tc.Line)

		if tc.ExpectErr != nil {
			require.Nil(t, ls)
			assert.NotEmpty(t, err.Error())
			assert.Equal(t, tc.ExpectErr.Line(), err.Line())
			assert.Equal(t, tc.ExpectErr.Col(), err.Col())
		}

		if tc.ExpectLexs != nil {
			require.Nil(t, err)
			assert.Equal(t, tc.ExpectLexs, ls)
		}
	}
}

type lexScanFunc func(*runer.RuneItr) *symbol.Lexeme

func lexFuncTest(t *testing.T, fName string, f lexScanFunc, tests []lexTest) {
	for i, tc := range tests {
		require.NotNil(t, tc.ExpectLex)
		require.Nil(t, tc.ExpectErr)

		t.Log(fName + "() test case: " + strconv.Itoa(i+1))

		itr := runer.NewRuneItr(tc.Input)
		l := f(itr)

		require.NotNil(t, l)
		assert.Equal(t, tc.ExpectLex, *l)
	}
}

type lexScanErrFunc func(*runer.RuneItr) (*symbol.Lexeme, LexError)

func lexErrFuncTest(t *testing.T, fName string, f lexScanErrFunc, tests []lexTest) {
	for i, tc := range tests {
		t.Log(fName + "() test case: " + strconv.Itoa(i+1))

		itr := runer.NewRuneItr(tc.Input)
		l, err := f(itr)

		if tc.ExpectErr != nil {
			assert.NotNil(t, err)
			require.Nil(t, l)
			assert.NotEmpty(t, err.Error())
			assert.Equal(t, tc.ExpectErr.Line(), err.Line())
			assert.Equal(t, tc.ExpectErr.Col(), err.Col())
		} else {
			assert.Nil(t, err)
			require.NotNil(t, l)
			assert.Equal(t, tc.ExpectLex, *l)
		}
	}
}

type lexTest struct {
	Line      int
	Input     string
	ExpectLex symbol.Lexeme
	ExpectErr LexError
}

type scanLineTest struct {
	Line       int
	Input      string
	ExpectLexs []symbol.Lexeme
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
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`1`, 5, 6, 0, symbol.NUMBER},
			},
		},
		scanLineTest{
			Input: `y <- -1.1`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`y`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`-`, 5, 6, 0, symbol.SUBTRACT},
				symbol.Lexeme{`1.1`, 6, 9, 0, symbol.NUMBER},
			},
		},
		scanLineTest{
			Line:  123,
			Input: `x <- true`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 123, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 1, 2, 123, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 123, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 123, symbol.WHITESPACE},
				symbol.Lexeme{`true`, 5, 9, 123, symbol.BOOLEAN_TRUE},
			},
		},
		scanLineTest{
			Input: `@Println["Whelp"]`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`@Println`, 0, 8, 0, symbol.SOURCERY},
				symbol.Lexeme{`[`, 8, 9, 0, symbol.SQUARE_BRACE_OPEN},
				symbol.Lexeme{`"Whelp"`, 9, 16, 0, symbol.STRING},
				symbol.Lexeme{`]`, 16, 17, 0, symbol.SQUARE_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: "\tresult <- spell(a, b) r, err     ",
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{"\t", 0, 1, 0, symbol.WHITESPACE},
				symbol.Lexeme{`result`, 1, 7, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Lexeme{`spell`, 11, 16, 0, symbol.KEYWORD_SPELL},
				symbol.Lexeme{`(`, 16, 17, 0, symbol.CURVED_BRACE_OPEN},
				symbol.Lexeme{`a`, 17, 18, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`,`, 18, 19, 0, symbol.VALUE_SEPARATOR},
				symbol.Lexeme{` `, 19, 20, 0, symbol.WHITESPACE},
				symbol.Lexeme{`b`, 20, 21, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`)`, 21, 22, 0, symbol.CURVED_BRACE_CLOSE},
				symbol.Lexeme{` `, 22, 23, 0, symbol.WHITESPACE},
				symbol.Lexeme{`r`, 23, 24, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`,`, 24, 25, 0, symbol.VALUE_SEPARATOR},
				symbol.Lexeme{` `, 25, 26, 0, symbol.WHITESPACE},
				symbol.Lexeme{`err`, 26, 29, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`     `, 29, 34, 0, symbol.WHITESPACE},
			},
		},
		scanLineTest{
			Input: `keyValue <- "pi": 3.1419`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`keyValue`, 0, 8, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Lexeme{`"pi"`, 12, 16, 0, symbol.STRING},
				symbol.Lexeme{`:`, 16, 17, 0, symbol.KEY_VALUE_SEPARATOR},
				symbol.Lexeme{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Lexeme{`3.1419`, 18, 24, 0, symbol.NUMBER},
			},
		},
		scanLineTest{
			Input: `alphabet <- ["a", "b", "c"]`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`alphabet`, 0, 8, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Lexeme{`[`, 12, 13, 0, symbol.SQUARE_BRACE_OPEN},
				symbol.Lexeme{`"a"`, 13, 16, 0, symbol.STRING},
				symbol.Lexeme{`,`, 16, 17, 0, symbol.VALUE_SEPARATOR},
				symbol.Lexeme{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Lexeme{`"b"`, 18, 21, 0, symbol.STRING},
				symbol.Lexeme{`,`, 21, 22, 0, symbol.VALUE_SEPARATOR},
				symbol.Lexeme{` `, 22, 23, 0, symbol.WHITESPACE},
				symbol.Lexeme{`"c"`, 23, 26, 0, symbol.STRING},
				symbol.Lexeme{`]`, 26, 27, 0, symbol.SQUARE_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: `loop i <- 0..5`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`i`, 5, 6, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 7, 9, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 9, 10, 0, symbol.WHITESPACE},
				symbol.Lexeme{`0`, 10, 11, 0, symbol.NUMBER},
				symbol.Lexeme{`..`, 11, 13, 0, symbol.RANGE},
				symbol.Lexeme{`5`, 13, 14, 0, symbol.NUMBER},
			},
		},
		scanLineTest{
			Input: `x<-2 // The value of x is now 2`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{`2`, 3, 4, 0, symbol.NUMBER},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`// The value of x is now 2`, 5, 31, 0, symbol.COMMENT},
			},
		},
		scanLineTest{
			Input: `isLandscape<-length<height`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`isLandscape`, 0, 11, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`<-`, 11, 13, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{`length`, 13, 19, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`<`, 19, 20, 0, symbol.LESS_THAN},
				symbol.Lexeme{`height`, 20, 26, 0, symbol.IDENTIFIER},
			},
		},
		scanLineTest{
			Input: `x<-3.14*(1-2+3)`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Lexeme{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{`3.14`, 3, 7, 0, symbol.NUMBER},
				symbol.Lexeme{`*`, 7, 8, 0, symbol.MULTIPLY},
				symbol.Lexeme{`(`, 8, 9, 0, symbol.CURVED_BRACE_OPEN},
				symbol.Lexeme{`1`, 9, 10, 0, symbol.NUMBER},
				symbol.Lexeme{`-`, 10, 11, 0, symbol.SUBTRACT},
				symbol.Lexeme{`2`, 11, 12, 0, symbol.NUMBER},
				symbol.Lexeme{`+`, 12, 13, 0, symbol.ADD},
				symbol.Lexeme{`3`, 13, 14, 0, symbol.NUMBER},
				symbol.Lexeme{`)`, 14, 15, 0, symbol.CURVED_BRACE_CLOSE},
			},
		},
		scanLineTest{
			Input: `!x => y <- _`,
			ExpectLexs: []symbol.Lexeme{
				symbol.Lexeme{`!`, 0, 1, 0, symbol.NEGATION},
				symbol.Lexeme{`x`, 1, 2, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 2, 3, 0, symbol.WHITESPACE},
				symbol.Lexeme{`=>`, 3, 5, 0, symbol.IF_MATCH_THEN},
				symbol.Lexeme{` `, 5, 6, 0, symbol.WHITESPACE},
				symbol.Lexeme{`y`, 6, 7, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Lexeme{`_`, 11, 12, 0, symbol.VOID},
			},
		},
	}
}
