package new_scanner

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
)

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

/*
// runScanTest runs the input test cases on the input
// function.
func runScanTest(
	t *testing.T,
	fileName string,
	f func([]rune, int) (*token.Token, []rune),
	tests []scanFuncTest) {

	for _, tc := range tests {
		require.NotNil(t, tc.Expect)
		require.Nil(t, tc.Error)

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		tk, out := f(tc.Input, tc.col)

		require.NotNil(t, tk, "Did not expect token to be nil")
		assert.Equal(t, tc.Output, out, "Expected a different array of leftover runes")
		assert.Equal(t, tc.Expect, *tk, "Expected a different token")
	}
}
*/

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

// runFailableScanTest runs the input test cases on the input function.
func runFailableScanTest(
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
