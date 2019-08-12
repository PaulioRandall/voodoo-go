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

func newFault(i int) fault.SyntaxFault {
	return fault.SyntaxFault{
		Index: i,
	}
}

// dummyToken creates a new dummy token.
func dummyToken(line, start, end int, val string, t token.TokenType) token.Token {
	return token.Token{
		//Line: line,
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
			assert.Empty(t, tk)
			assert.NotNil(t, err)
		}

		if tc.Expect != token.EMPTY {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expect, tk)

			next := readRequireNoErr(t, r)
			assert.Equal(t, tc.NextUnreadRune, next)
		}
	}
}
