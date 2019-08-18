package scanner

import (
	"bufio"
	"strconv"
	"strings"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// scanTest represents a test case for the Scan() function.
type scanTest struct {
	TestLine int
	Input    string
	Expect   []token.Token
	Error    fault.Fault
}

// collateLine collates a single line from a channel of tokens.
func collateLine(in chan token.Token, out chan []token.Token) {
	defer close(out)
	tks := []token.Token{}
	for tk := range in {
		tks = append(tks, tk)
	}
	out <- tks
}

// printTestTitle prints the title and line number of the test.
func printTestTitle(t *testing.T, lineNum int) {
	testLine := strconv.Itoa(lineNum)
	t.Log("-> scanner_test.go : " + testLine)
}

// newRuner returns a new Runer instance.
func newRuner(s string) *Runer {
	sr := strings.NewReader(s)
	br := bufio.NewReader(sr)
	return NewRuner(br)
}

// makeChans makes some channels to use while testing.
func makeChans() (chan token.Token, chan []token.Token) {
	inChan := make(chan token.Token)
	outChan := make(chan []token.Token)
	return inChan, outChan
}

// TestScan runs the test cases for the Scan() function.
func TestScan(t *testing.T) {
	for _, tc := range scanTests() {
		printTestTitle(t, tc.TestLine)

		inChan, outChan := makeChans()
		go collateLine(inChan, outChan)

		r := newRuner(tc.Input)
		_ = Scan(r, false, inChan)
		act := <-outChan

		for i, expTk := range tc.Expect {
			require.True(t, i < len(act))
			actTk := act[i]
			assertToken(t, expTk, actTk)
		}

		assert.Equal(t, len(tc.Expect), len(act))
	}
}

// scanTests creates a list of Scan() function tests.
func scanTests() []scanTest {
	return []scanTest{
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x # 1`,
			Expect: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				errDummyToken(0, 2, 3),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `123.456.789`,
			Expect: []token.Token{
				dummyToken(0, 0, 7, `123.456`, token.TT_NUMBER),
				errDummyToken(0, 7, 8),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- 1`,
			Expect: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 6, `1`, token.TT_NUMBER),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `y <- -1.1`,
			Expect: []token.Token{
				dummyToken(0, 0, 1, `y`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 6, `-`, token.TT_SUBTRACT),
				dummyToken(0, 6, 9, `1.1`, token.TT_NUMBER),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x <- true`,
			Expect: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 9, `true`, token.TT_TRUE),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `@Println["Whelp"]`,
			Expect: []token.Token{
				dummyToken(0, 0, 8, `@Println`, token.TT_SPELL),
				dummyToken(0, 8, 9, `[`, token.TT_SQUARE_OPEN),
				dummyToken(0, 9, 16, `"Whelp"`, token.TT_STRING),
				dummyToken(0, 16, 17, `]`, token.TT_SQUARE_CLOSE),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    "\tresult <- func(a, b) r, err     ",
			Expect: []token.Token{
				dummyToken(0, 0, 1, "\t", token.TT_SPACE),
				dummyToken(0, 1, 7, `result`, token.TT_ID),
				dummyToken(0, 7, 8, ` `, token.TT_SPACE),
				dummyToken(0, 8, 10, `<-`, token.TT_ASSIGN),
				dummyToken(0, 10, 11, ` `, token.TT_SPACE),
				dummyToken(0, 11, 15, `func`, token.TT_FUNC),
				dummyToken(0, 15, 16, `(`, token.TT_CURVY_OPEN),
				dummyToken(0, 16, 17, `a`, token.TT_ID),
				dummyToken(0, 17, 18, `,`, token.TT_VALUE_DELIM),
				dummyToken(0, 18, 19, ` `, token.TT_SPACE),
				dummyToken(0, 19, 20, `b`, token.TT_ID),
				dummyToken(0, 20, 21, `)`, token.TT_CURVY_CLOSE),
				dummyToken(0, 21, 22, ` `, token.TT_SPACE),
				dummyToken(0, 22, 23, `r`, token.TT_ID),
				dummyToken(0, 23, 24, `,`, token.TT_VALUE_DELIM),
				dummyToken(0, 24, 25, ` `, token.TT_SPACE),
				dummyToken(0, 25, 28, `err`, token.TT_ID),
				dummyToken(0, 28, 33, `     `, token.TT_SPACE),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `alphabet <- ["a", "b", "c"]`,
			Expect: []token.Token{
				dummyToken(0, 0, 8, `alphabet`, token.TT_ID),
				dummyToken(0, 8, 9, ` `, token.TT_SPACE),
				dummyToken(0, 9, 11, `<-`, token.TT_ASSIGN),
				dummyToken(0, 11, 12, ` `, token.TT_SPACE),
				dummyToken(0, 12, 13, `[`, token.TT_SQUARE_OPEN),
				dummyToken(0, 13, 16, `"a"`, token.TT_STRING),
				dummyToken(0, 16, 17, `,`, token.TT_VALUE_DELIM),
				dummyToken(0, 17, 18, ` `, token.TT_SPACE),
				dummyToken(0, 18, 21, `"b"`, token.TT_STRING),
				dummyToken(0, 21, 22, `,`, token.TT_VALUE_DELIM),
				dummyToken(0, 22, 23, ` `, token.TT_SPACE),
				dummyToken(0, 23, 26, `"c"`, token.TT_STRING),
				dummyToken(0, 26, 27, `]`, token.TT_SQUARE_CLOSE),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-2 // The value of x is now 2`,
			Expect: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 3, `<-`, token.TT_ASSIGN),
				dummyToken(0, 3, 4, `2`, token.TT_NUMBER),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 31, `// The value of x is now 2`, token.TT_COMMENT),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `isLandscape<-length<height`,
			Expect: []token.Token{
				dummyToken(0, 0, 11, `isLandscape`, token.TT_ID),
				dummyToken(0, 11, 13, `<-`, token.TT_ASSIGN),
				dummyToken(0, 13, 19, `length`, token.TT_ID),
				dummyToken(0, 19, 20, `<`, token.TT_CMP_LT),
				dummyToken(0, 20, 26, `height`, token.TT_ID),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `x<-3.14*(1-2+3)`,
			Expect: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 3, `<-`, token.TT_ASSIGN),
				dummyToken(0, 3, 7, `3.14`, token.TT_NUMBER),
				dummyToken(0, 7, 8, `*`, token.TT_MULTIPLY),
				dummyToken(0, 8, 9, `(`, token.TT_CURVY_OPEN),
				dummyToken(0, 9, 10, `1`, token.TT_NUMBER),
				dummyToken(0, 10, 11, `-`, token.TT_SUBTRACT),
				dummyToken(0, 11, 12, `2`, token.TT_NUMBER),
				dummyToken(0, 12, 13, `+`, token.TT_ADD),
				dummyToken(0, 13, 14, `3`, token.TT_NUMBER),
				dummyToken(0, 14, 15, `)`, token.TT_CURVY_CLOSE),
			},
		},
		scanTest{
			TestLine: fault.CurrLine(),
			Input:    `!x => y <- _`,
			Expect: []token.Token{
				dummyToken(0, 0, 1, `!`, token.TT_NOT),
				dummyToken(0, 1, 2, `x`, token.TT_ID),
				dummyToken(0, 2, 3, ` `, token.TT_SPACE),
				dummyToken(0, 3, 5, `=>`, token.TT_MATCH),
				dummyToken(0, 5, 6, ` `, token.TT_SPACE),
				dummyToken(0, 6, 7, `y`, token.TT_ID),
				dummyToken(0, 7, 8, ` `, token.TT_SPACE),
				dummyToken(0, 8, 10, `<-`, token.TT_ASSIGN),
				dummyToken(0, 10, 11, ` `, token.TT_SPACE),
				dummyToken(0, 11, 12, `_`, token.TT_VOID),
			},
		},
	}
}
