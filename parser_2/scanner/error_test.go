package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertScanError asserts the ScanErrors are equal except for the error
// messages. The error messages are checked to ensure they are not empty.
func AssertScanError(t *testing.T, exp ScanError, act ScanError) bool {
	if exp == nil {
		return true
	}

	require.NotEmpty(t, act.Errors(), `ScanError.Errors()`)

	return logicalConjunction(
		assert.Equal(t, exp.Line(), act.Line(), `ScanError.Line()`),
		assert.Equal(t, exp.Index(), act.Index(), `ScanError.Index()`),
		assertScanErrors(t, act),
	)
}

func assertScanErrors(t *testing.T, act ScanError) bool {
	ok := true
	for i, e := range act.Errors() {
		ok = ok && assert.NotEmpty(t, e, `ScanError.Errors()[%d]`, i)
	}
	return ok
}
