package lexer

import (
	"strconv"
	"testing"

	sym "github.com/PaulioRandall/voodoo-go/symbol"
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
			assert.Equal(t, tc.Expects, s)
		}
	}
}

type symTest struct {
	Line      int
	Input     string
	Expects   sym.Symbol
	ExpectErr bool
}

type symArrayTest struct {
	Line      int
	Input     string
	Expects   []sym.Symbol
	ExpectErr bool
}

func apiTests() []symArrayTest {
	return []symArrayTest{
		symArrayTest{
			Input:     `x # 1`,
			ExpectErr: true,
		},
		symArrayTest{
			Input: `x <- 1`,
			Expects: []sym.Symbol{
				sym.Symbol{`x`, 0, 1, 0, sym.VARIABLE},
				sym.Symbol{` `, 1, 2, 0, sym.WHITESPACE},
				sym.Symbol{`<-`, 2, 4, 0, sym.ASSIGNMENT},
				sym.Symbol{` `, 4, 5, 0, sym.WHITESPACE},
				sym.Symbol{`1`, 5, 6, 0, sym.NUMBER},
			},
		},
		symArrayTest{
			Input: `y <- -1.1`,
			Expects: []sym.Symbol{
				sym.Symbol{`y`, 0, 1, 0, sym.VARIABLE},
				sym.Symbol{` `, 1, 2, 0, sym.WHITESPACE},
				sym.Symbol{`<-`, 2, 4, 0, sym.ASSIGNMENT},
				sym.Symbol{` `, 4, 5, 0, sym.WHITESPACE},
				sym.Symbol{`-`, 5, 6, 0, sym.SUBTRACT},
				sym.Symbol{`1.1`, 6, 9, 0, sym.NUMBER},
			},
		},
		symArrayTest{
			Line:  123,
			Input: `x <- true`,
			Expects: []sym.Symbol{
				sym.Symbol{`x`, 0, 1, 123, sym.VARIABLE},
				sym.Symbol{` `, 1, 2, 123, sym.WHITESPACE},
				sym.Symbol{`<-`, 2, 4, 123, sym.ASSIGNMENT},
				sym.Symbol{` `, 4, 5, 123, sym.WHITESPACE},
				sym.Symbol{`true`, 5, 9, 123, sym.BOOLEAN},
			},
		},
		symArrayTest{
			Input: `@Println["Whelp"]`,
			Expects: []sym.Symbol{
				sym.Symbol{`@Println`, 0, 8, 0, sym.SOURCERY},
				sym.Symbol{`[`, 8, 9, 0, sym.SQUARE_BRACE_OPEN},
				sym.Symbol{`"Whelp"`, 9, 16, 0, sym.STRING},
				sym.Symbol{`]`, 16, 17, 0, sym.SQUARE_BRACE_CLOSE},
			},
		},
		symArrayTest{
			Input: "\tresult <- spell(a, b) r, err     ",
			Expects: []sym.Symbol{
				sym.Symbol{"\t", 0, 1, 0, sym.WHITESPACE},
				sym.Symbol{`result`, 1, 7, 0, sym.VARIABLE},
				sym.Symbol{` `, 7, 8, 0, sym.WHITESPACE},
				sym.Symbol{`<-`, 8, 10, 0, sym.ASSIGNMENT},
				sym.Symbol{` `, 10, 11, 0, sym.WHITESPACE},
				sym.Symbol{`spell`, 11, 16, 0, sym.KEYWORD_SPELL},
				sym.Symbol{`(`, 16, 17, 0, sym.CURVED_BRACE_OPEN},
				sym.Symbol{`a`, 17, 18, 0, sym.VARIABLE},
				sym.Symbol{`,`, 18, 19, 0, sym.VALUE_SEPARATOR},
				sym.Symbol{` `, 19, 20, 0, sym.WHITESPACE},
				sym.Symbol{`b`, 20, 21, 0, sym.VARIABLE},
				sym.Symbol{`)`, 21, 22, 0, sym.CURVED_BRACE_CLOSE},
				sym.Symbol{` `, 22, 23, 0, sym.WHITESPACE},
				sym.Symbol{`r`, 23, 24, 0, sym.VARIABLE},
				sym.Symbol{`,`, 24, 25, 0, sym.VALUE_SEPARATOR},
				sym.Symbol{` `, 25, 26, 0, sym.WHITESPACE},
				sym.Symbol{`err`, 26, 29, 0, sym.VARIABLE},
				sym.Symbol{`     `, 29, 34, 0, sym.WHITESPACE},
			},
		},
		symArrayTest{
			Input: `keyValue <- "pi": 3.1419`,
			Expects: []sym.Symbol{
				sym.Symbol{`keyValue`, 0, 8, 0, sym.VARIABLE},
				sym.Symbol{` `, 8, 9, 0, sym.WHITESPACE},
				sym.Symbol{`<-`, 9, 11, 0, sym.ASSIGNMENT},
				sym.Symbol{` `, 11, 12, 0, sym.WHITESPACE},
				sym.Symbol{`"pi"`, 12, 16, 0, sym.STRING},
				sym.Symbol{`:`, 16, 17, 0, sym.KEY_VALUE_SEPARATOR},
				sym.Symbol{` `, 17, 18, 0, sym.WHITESPACE},
				sym.Symbol{`3.1419`, 18, 24, 0, sym.NUMBER},
			},
		},
		symArrayTest{
			Input: `alphabet <- ["a", "b", "c"]`,
			Expects: []sym.Symbol{
				sym.Symbol{`alphabet`, 0, 8, 0, sym.VARIABLE},
				sym.Symbol{` `, 8, 9, 0, sym.WHITESPACE},
				sym.Symbol{`<-`, 9, 11, 0, sym.ASSIGNMENT},
				sym.Symbol{` `, 11, 12, 0, sym.WHITESPACE},
				sym.Symbol{`[`, 12, 13, 0, sym.SQUARE_BRACE_OPEN},
				sym.Symbol{`"a"`, 13, 16, 0, sym.STRING},
				sym.Symbol{`,`, 16, 17, 0, sym.VALUE_SEPARATOR},
				sym.Symbol{` `, 17, 18, 0, sym.WHITESPACE},
				sym.Symbol{`"b"`, 18, 21, 0, sym.STRING},
				sym.Symbol{`,`, 21, 22, 0, sym.VALUE_SEPARATOR},
				sym.Symbol{` `, 22, 23, 0, sym.WHITESPACE},
				sym.Symbol{`"c"`, 23, 26, 0, sym.STRING},
				sym.Symbol{`]`, 26, 27, 0, sym.SQUARE_BRACE_CLOSE},
			},
		},
		symArrayTest{
			Input: `loop i <- 0..5`,
			Expects: []sym.Symbol{
				sym.Symbol{`loop`, 0, 4, 0, sym.KEYWORD_LOOP},
				sym.Symbol{` `, 4, 5, 0, sym.WHITESPACE},
				sym.Symbol{`i`, 5, 6, 0, sym.VARIABLE},
				sym.Symbol{` `, 6, 7, 0, sym.WHITESPACE},
				sym.Symbol{`<-`, 7, 9, 0, sym.ASSIGNMENT},
				sym.Symbol{` `, 9, 10, 0, sym.WHITESPACE},
				sym.Symbol{`0`, 10, 11, 0, sym.NUMBER},
				sym.Symbol{`..`, 11, 13, 0, sym.RANGE},
				sym.Symbol{`5`, 13, 14, 0, sym.NUMBER},
			},
		},
		symArrayTest{
			Input: `x<-2 // The value of x is now 2`,
			Expects: []sym.Symbol{
				sym.Symbol{`x`, 0, 1, 0, sym.VARIABLE},
				sym.Symbol{`<-`, 1, 3, 0, sym.ASSIGNMENT},
				sym.Symbol{`2`, 3, 4, 0, sym.NUMBER},
				sym.Symbol{` `, 4, 5, 0, sym.WHITESPACE},
				sym.Symbol{`// The value of x is now 2`, 5, 31, 0, sym.COMMENT},
			},
		},
		symArrayTest{
			Input: `isLandscape<-length<height`,
			Expects: []sym.Symbol{
				sym.Symbol{`isLandscape`, 0, 11, 0, sym.VARIABLE},
				sym.Symbol{`<-`, 11, 13, 0, sym.ASSIGNMENT},
				sym.Symbol{`length`, 13, 19, 0, sym.VARIABLE},
				sym.Symbol{`<`, 19, 20, 0, sym.LESS_THAN},
				sym.Symbol{`height`, 20, 26, 0, sym.VARIABLE},
			},
		},
		symArrayTest{
			Input: `x<-3.14*(1-2+3)`,
			Expects: []sym.Symbol{
				sym.Symbol{`x`, 0, 1, 0, sym.VARIABLE},
				sym.Symbol{`<-`, 1, 3, 0, sym.ASSIGNMENT},
				sym.Symbol{`3.14`, 3, 7, 0, sym.NUMBER},
				sym.Symbol{`*`, 7, 8, 0, sym.MULTIPLY},
				sym.Symbol{`(`, 8, 9, 0, sym.CURVED_BRACE_OPEN},
				sym.Symbol{`1`, 9, 10, 0, sym.NUMBER},
				sym.Symbol{`-`, 10, 11, 0, sym.SUBTRACT},
				sym.Symbol{`2`, 11, 12, 0, sym.NUMBER},
				sym.Symbol{`+`, 12, 13, 0, sym.ADD},
				sym.Symbol{`3`, 13, 14, 0, sym.NUMBER},
				sym.Symbol{`)`, 14, 15, 0, sym.CURVED_BRACE_CLOSE},
			},
		},
		symArrayTest{
			Input: `!x => y <- _`,
			Expects: []sym.Symbol{
				sym.Symbol{`!`, 0, 1, 0, sym.NEGATION},
				sym.Symbol{`x`, 1, 2, 0, sym.VARIABLE},
				sym.Symbol{` `, 2, 3, 0, sym.WHITESPACE},
				sym.Symbol{`=>`, 3, 5, 0, sym.IF_TRUE_THEN},
				sym.Symbol{` `, 5, 6, 0, sym.WHITESPACE},
				sym.Symbol{`y`, 6, 7, 0, sym.VARIABLE},
				sym.Symbol{` `, 7, 8, 0, sym.WHITESPACE},
				sym.Symbol{`<-`, 8, 10, 0, sym.ASSIGNMENT},
				sym.Symbol{` `, 10, 11, 0, sym.WHITESPACE},
				sym.Symbol{`_`, 11, 12, 0, sym.VOID},
			},
		},
	}
}
