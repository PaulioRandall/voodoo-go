package scanner

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// scanFuncTest represents a test case for any of the
// delegate scanning functions.
type scanFuncTest struct {
	TestLine  int
	Input     string
	Expect    token.Token
	ExpectErr fault.Fault
}

// runScanTest runs the input test cases on the input
// function.
func runScanTest(
	t *testing.T,
	fileName string,
	f func(*runer.RuneItr) *token.Token,
	tests []scanFuncTest) {

	for _, tc := range tests {
		require.NotNil(t, tc.Expect)
		require.Nil(t, tc.ExpectErr)

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		itr := runer.NewRuneItr(tc.Input)
		act := f(itr)

		require.NotNil(t, act)
		assert.Equal(t, tc.Expect, *act)
	}
}

// runFailableScanTest runs the input test cases on the
// input function.
func runFailableScanTest(
	t *testing.T,
	fileName string,
	f func(*runer.RuneItr) (*token.Token, fault.Fault),
	tests []scanFuncTest) {

	for _, tc := range tests {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> " + fileName + " : " + testLine)

		itr := runer.NewRuneItr(tc.Input)
		act, err := f(itr)

		if tc.ExpectErr != nil {
			assert.Nil(t, act)
			require.NotNil(t, err)
			assert.NotEmpty(t, err.Error())
			fault.Assert(t, tc.ExpectErr, err)

		} else {
			assert.Nil(t, err)
			require.NotNil(t, act)
			assert.Equal(t, tc.Expect, *act)
		}
	}
}
