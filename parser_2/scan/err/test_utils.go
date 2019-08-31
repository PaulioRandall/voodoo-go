package err

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertEqual asserts the ScanErrors are equal except for the error messages.
// The error messages are checked to ensure they are not empty.
func AssertEqual(t *testing.T, exp ScanError, act ScanError) bool {
	if exp == nil {
		return assert.Nil(t, act, `ScanError == nil`)
	} else {
		require.NotNil(t, act, `ScanError != nil`)
	}

	require.NotEmpty(t, act.Errors(), `ScanError.Errors()`)

	return utils.LogicalConjunction(
		assert.Equal(t, exp.Line(), act.Line(), `ScanError.Line()`),
		assert.Equal(t, exp.Index(), act.Index(), `ScanError.Index()`),
		assertErrors(t, act),
	)
}

func assertErrors(t *testing.T, act ScanError) bool {
	ok := true
	for i, e := range act.Errors() {
		ok = ok && assert.NotEmpty(t, e, `ScanError.Errors()[%d]`, i)
	}
	return ok
}
