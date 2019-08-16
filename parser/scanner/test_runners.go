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

func dummyRuner(s string) *Runer {
	sr := strings.NewReader(s)
	br := bufio.NewReader(sr)
	return NewRuner(br)
}

func readRequireNoErr(t *testing.T, r *Runer) rune {
	ru, err := r.ReadRune()
	require.Nil(t, err)
	return ru
}

// scanFuncTest represents a test case for any of the delegate scanning
// functions.
type scanFuncTest struct {
	TestLine       int
	Input          string
	Expect         token.Token
	NextUnreadRune rune
	Error          fault.Fault
}

func newFault(i int) fault.Fault {
	return token.SyntaxFault{
		Col: i,
	}
}

// dummyToken creates a new dummy token.
func dummyToken(line, start, end int, val string, t token.TokenType) token.Token {
	return token.Token{
		Line:  line,
		Start: start,
		End:   end,
		Val:   val,
		Type:  t,
	}
}

// runScanTest runs the input test cases on the input function.
func runScanTest(
	t *testing.T,
	fileName string,
	f func(*Runer) (token.Token, fault.Fault),
	tests []scanFuncTest) {

	for _, tc := range tests {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		r := dummyRuner(tc.Input)
		tk, err := f(r)

		if tc.Error != nil {
			assert.NotNil(t, err)
		}

		assertToken(t, tc.Expect, tk)

		if tk.Type != token.TT_ERROR_UPSTREAM {
			next := readRequireNoErr(t, r)
			assert.Equal(t, tc.NextUnreadRune, next)
		}
	}
}

// assertToken asserts that the actual token equals the expected token except
// for the error messages.
func assertToken(t *testing.T, exp token.Token, act token.Token) {
	assert.Equal(t, exp.Val, act.Val)
	assert.Equal(t, exp.Line, act.Line)
	assert.Equal(t, exp.Start, act.Start)
	assert.Equal(t, exp.End, act.End)
	assert.Equal(t, exp.Type, act.Type)
}
