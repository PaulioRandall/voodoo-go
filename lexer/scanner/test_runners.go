package scanner

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/new_fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// scanFuncTest represents a test case for any of the
// delegate scanning functions.
type scanFuncTest struct {
	TestLine     int
	Input        []rune
	Output       []rune
	Expect       token.Token
	ExpectNewErr new_fault.Fault
}

func newFault(i int) new_fault.SyntaxFault {
	return new_fault.SyntaxFault{
		Index: i,
	}
}

// runScanTest runs the input test cases on the input
// function.
func runScanTest(
	t *testing.T,
	fileName string,
	f func([]rune) (*token.Token, []rune),
	tests []scanFuncTest) {

	for _, tc := range tests {
		require.NotNil(t, tc.Expect)
		require.Nil(t, tc.ExpectNewErr)

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		tk, out := f(tc.Input)

		require.NotNil(t, tk, "Did not expect token to be nil")
		assert.Equal(t, tc.Output, out, "Expected a different array of leftover runes")
		assert.Equal(t, tc.Expect, *tk, "Expected a different token")
	}
}

// runFailableScanTest runs the input test cases on the
// input function.
func new_runFailableScanTest(
	t *testing.T,
	fileName string,
	f func([]rune) (*token.Token, []rune, new_fault.Fault),
	tests []scanFuncTest) {

	for _, tc := range tests {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		tk, out, err := f(tc.Input)

		if tc.ExpectNewErr != nil {
			assert.Nil(t, tk, "Expected token to be nil")
			require.NotNil(t, err, "Did NOT expect error to be nil")
		} else {
			assert.Nil(t, err, "Expected error to be nil")
			require.NotNil(t, tk, "Did NOT expect token to be nil")
			assert.Equal(t, tc.Output, out, "Expected a different array of leftover runes")
			assert.Equal(t, tc.Expect, *tk, "Expected a different token")
		}
	}
}
