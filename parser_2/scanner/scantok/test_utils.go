package scantok

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/token"
	"github.com/PaulioRandall/voodoo-go/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertEqual asserts that the the actual Token is a scanTok and that
// it equals the expected Token which must also be a scanTok.
func AssertEqual(t *testing.T, exp token.Token, act token.Token) bool {
	if exp == nil {
		return assert.Nil(t, act, `Token == nil`)
	}

	require.NotNil(t, act, `Token != nil`)

	actTk, ok := act.(scanTok)
	require.True(t, ok, `Expected-Token.(scanTok)`)
	expTk, ok := act.(scanTok)
	require.True(t, ok, `Actual-Token.(scanTok)`)

	return utils.LogicalConjunction(
		assert.Equal(t, expTk.text, actTk.text, `scanTok.text`),
		assert.Equal(t, expTk.line, actTk.line, `scanTok.line`),
		assert.Equal(t, expTk.start, actTk.start, `scanTok.start`),
		assert.Equal(t, expTk.end, actTk.end, `scanTok.end`),
		assert.Equal(t, expTk.kind, actTk.kind, `scanTok.kind`),
	)
}
