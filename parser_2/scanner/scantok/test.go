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
func AssertEqual(t *testing.T, exp, act token.Token) bool {
	if exp == nil {
		return assert.Nil(t, act, `Token == nil`)
	}

	require.NotNil(t, act, `Token != nil`)

	actTk, ok := act.(scanTok)
	require.True(t, ok, `Expected-Token.(scanTok)`)
	expTk, ok := exp.(scanTok)
	require.True(t, ok, `Actual-Token.(scanTok)`)

	return utils.LogicalConjunction(
		assert.Equal(t, expTk.text, actTk.text, `scanTok.text`),
		assert.Equal(t, expTk.line, actTk.line, `scanTok.line`),
		assert.Equal(t, expTk.start, actTk.start, `scanTok.start`),
		assert.Equal(t, expTk.end, actTk.end, `scanTok.end`),
		assert.Equal(t,
			token.KindName(expTk.kind),
			token.KindName(actTk.kind),
			`scanTok.kind`,
		),
	)
}

// AssertSliceEqual asserts that the the actual Token slice contains only
// scanTok instances and that each equals the corresponding one from the
// expected Token slice.
func AssertSliceEqual(t *testing.T, exp, act []token.Token) bool {
	require.NotNil(t, exp, `exp != nil`)
	require.NotNil(t, act, `act != nil`)

	ok := true
	for i, _ := range exp {
		require.True(t, i < len(act), `exp[i]: i < len(act)`)
		ok = ok && AssertEqual(t, exp[i], act[i])
	}

	return ok && assert.Equal(t, len(exp), len(act), `len(exp) == len(act)`)
}
