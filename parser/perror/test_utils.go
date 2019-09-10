package perror

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertEqual asserts the Perrors are equal except for the error messages.
// The error messages are simple checked to ensure they are not empty.
func AssertEqual(t *testing.T, exp Perror, act Perror) bool {
	if exp == nil {
		return assert.Nil(t, act, `Perror == nil`)
	} else {
		require.NotNil(t, act, `Perror != nil`)
	}

	require.NotEmpty(t, act.Errors(), `Perror.Errors()`)

	return utils.LogicalConjunction(
		assert.Equal(t, exp.Line(), act.Line(), `Perror.Line()`),
		assert.Equal(t, exp.Cols(), act.Cols(), `Perror.Cols()`),
		assertErrors(t, act),
	)
}

func assertErrors(t *testing.T, act Perror) bool {
	ok := true
	for i, e := range act.Errors() {
		ok = ok && assert.NotEmpty(t, e, `Perror.Errors()[%d]`, i)
	}
	return ok
}
