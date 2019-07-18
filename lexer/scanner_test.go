package lexer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScannerApi(t *testing.T) {
	for i, tc := range apiTests() {
		t.Log("Scanner test case: " + strconv.Itoa(i+1))
		s, err := ScanLine(tc.Input, tc.Line)

		if tc.ExpectErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
			assert.Equal(t, tc.ExpectSyms, s)
		}
	}
}

type symTest struct {
	Line      int
	Input     string
	ExpectSym symbol.Symbol
	ExpectErr bool
}

type scanLinesTest struct {
	Line       int
	Input      string
	ExpectSyms []symbol.Symbol
	ExpectErr  bool
}

func apiTests() []scanLinesTest {
	return []scanLinesTest{
		scanLinesTest{
			Input:     `x # 1`,
			ExpectErr: true,
		},
		scanLinesTest{
			Input: `x <- 1`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`x`, 0, 1, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Symbol{`1`, 5, 6, 0, symbol.NUMBER},
			},
		},
		scanLinesTest{
			Input: `y <- -1.1`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`y`, 0, 1, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Symbol{`-`, 5, 6, 0, symbol.SUBTRACT},
				symbol.Symbol{`1.1`, 6, 9, 0, symbol.NUMBER},
			},
		},
		scanLinesTest{
			Line:  123,
			Input: `x <- true`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`x`, 0, 1, 123, symbol.VARIABLE},
				symbol.Symbol{` `, 1, 2, 123, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 2, 4, 123, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 4, 5, 123, symbol.WHITESPACE},
				symbol.Symbol{`true`, 5, 9, 123, symbol.BOOLEAN},
			},
		},
		scanLinesTest{
			Input: `@Println["Whelp"]`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`@Println`, 0, 8, 0, symbol.SOURCERY},
				symbol.Symbol{`[`, 8, 9, 0, symbol.SQUARE_BRACE_OPEN},
				symbol.Symbol{`"Whelp"`, 9, 16, 0, symbol.STRING},
				symbol.Symbol{`]`, 16, 17, 0, symbol.SQUARE_BRACE_CLOSE},
			},
		},
		scanLinesTest{
			Input: "\tresult <- spell(a, b) r, err     ",
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{"\t", 0, 1, 0, symbol.WHITESPACE},
				symbol.Symbol{`result`, 1, 7, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Symbol{`spell`, 11, 16, 0, symbol.KEYWORD_SPELL},
				symbol.Symbol{`(`, 16, 17, 0, symbol.CURVED_BRACE_OPEN},
				symbol.Symbol{`a`, 17, 18, 0, symbol.VARIABLE},
				symbol.Symbol{`,`, 18, 19, 0, symbol.VALUE_SEPARATOR},
				symbol.Symbol{` `, 19, 20, 0, symbol.WHITESPACE},
				symbol.Symbol{`b`, 20, 21, 0, symbol.VARIABLE},
				symbol.Symbol{`)`, 21, 22, 0, symbol.CURVED_BRACE_CLOSE},
				symbol.Symbol{` `, 22, 23, 0, symbol.WHITESPACE},
				symbol.Symbol{`r`, 23, 24, 0, symbol.VARIABLE},
				symbol.Symbol{`,`, 24, 25, 0, symbol.VALUE_SEPARATOR},
				symbol.Symbol{` `, 25, 26, 0, symbol.WHITESPACE},
				symbol.Symbol{`err`, 26, 29, 0, symbol.VARIABLE},
				symbol.Symbol{`     `, 29, 34, 0, symbol.WHITESPACE},
			},
		},
		scanLinesTest{
			Input: `keyValue <- "pi": 3.1419`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`keyValue`, 0, 8, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Symbol{`"pi"`, 12, 16, 0, symbol.STRING},
				symbol.Symbol{`:`, 16, 17, 0, symbol.KEY_VALUE_SEPARATOR},
				symbol.Symbol{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Symbol{`3.1419`, 18, 24, 0, symbol.NUMBER},
			},
		},
		scanLinesTest{
			Input: `alphabet <- ["a", "b", "c"]`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`alphabet`, 0, 8, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 8, 9, 0, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 9, 11, 0, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 11, 12, 0, symbol.WHITESPACE},
				symbol.Symbol{`[`, 12, 13, 0, symbol.SQUARE_BRACE_OPEN},
				symbol.Symbol{`"a"`, 13, 16, 0, symbol.STRING},
				symbol.Symbol{`,`, 16, 17, 0, symbol.VALUE_SEPARATOR},
				symbol.Symbol{` `, 17, 18, 0, symbol.WHITESPACE},
				symbol.Symbol{`"b"`, 18, 21, 0, symbol.STRING},
				symbol.Symbol{`,`, 21, 22, 0, symbol.VALUE_SEPARATOR},
				symbol.Symbol{` `, 22, 23, 0, symbol.WHITESPACE},
				symbol.Symbol{`"c"`, 23, 26, 0, symbol.STRING},
				symbol.Symbol{`]`, 26, 27, 0, symbol.SQUARE_BRACE_CLOSE},
			},
		},
		scanLinesTest{
			Input: `loop i <- 0..5`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
				symbol.Symbol{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Symbol{`i`, 5, 6, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 7, 9, 0, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 9, 10, 0, symbol.WHITESPACE},
				symbol.Symbol{`0`, 10, 11, 0, symbol.NUMBER},
				symbol.Symbol{`..`, 11, 13, 0, symbol.RANGE},
				symbol.Symbol{`5`, 13, 14, 0, symbol.NUMBER},
			},
		},
		scanLinesTest{
			Input: `x<-2 // The value of x is now 2`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`x`, 0, 1, 0, symbol.VARIABLE},
				symbol.Symbol{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Symbol{`2`, 3, 4, 0, symbol.NUMBER},
				symbol.Symbol{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Symbol{`// The value of x is now 2`, 5, 31, 0, symbol.COMMENT},
			},
		},
		scanLinesTest{
			Input: `isLandscape<-length<height`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`isLandscape`, 0, 11, 0, symbol.VARIABLE},
				symbol.Symbol{`<-`, 11, 13, 0, symbol.ASSIGNMENT},
				symbol.Symbol{`length`, 13, 19, 0, symbol.VARIABLE},
				symbol.Symbol{`<`, 19, 20, 0, symbol.LESS_THAN},
				symbol.Symbol{`height`, 20, 26, 0, symbol.VARIABLE},
			},
		},
		scanLinesTest{
			Input: `x<-3.14*(1-2+3)`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`x`, 0, 1, 0, symbol.VARIABLE},
				symbol.Symbol{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
				symbol.Symbol{`3.14`, 3, 7, 0, symbol.NUMBER},
				symbol.Symbol{`*`, 7, 8, 0, symbol.MULTIPLY},
				symbol.Symbol{`(`, 8, 9, 0, symbol.CURVED_BRACE_OPEN},
				symbol.Symbol{`1`, 9, 10, 0, symbol.NUMBER},
				symbol.Symbol{`-`, 10, 11, 0, symbol.SUBTRACT},
				symbol.Symbol{`2`, 11, 12, 0, symbol.NUMBER},
				symbol.Symbol{`+`, 12, 13, 0, symbol.ADD},
				symbol.Symbol{`3`, 13, 14, 0, symbol.NUMBER},
				symbol.Symbol{`)`, 14, 15, 0, symbol.CURVED_BRACE_CLOSE},
			},
		},
		scanLinesTest{
			Input: `!x => y <- _`,
			ExpectSyms: []symbol.Symbol{
				symbol.Symbol{`!`, 0, 1, 0, symbol.NEGATION},
				symbol.Symbol{`x`, 1, 2, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 2, 3, 0, symbol.WHITESPACE},
				symbol.Symbol{`=>`, 3, 5, 0, symbol.IF_TRUE_THEN},
				symbol.Symbol{` `, 5, 6, 0, symbol.WHITESPACE},
				symbol.Symbol{`y`, 6, 7, 0, symbol.VARIABLE},
				symbol.Symbol{` `, 7, 8, 0, symbol.WHITESPACE},
				symbol.Symbol{`<-`, 8, 10, 0, symbol.ASSIGNMENT},
				symbol.Symbol{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Symbol{`_`, 11, 12, 0, symbol.VOID},
			},
		},
	}
}
