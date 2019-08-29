package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertScanTokEqual asserts that the the actual Token is a scanTok and that
// it equals the expected scanTok.
func AssertScanTokEqual(t *testing.T, exp *scanTok, act token.Token) bool {
	if exp == nil {
		return assert.Nil(t, act, `Token == nil`)
	}

	require.NotNil(t, act, `Token != nil`)
	actTk, ok := act.(scanTok)
	require.True(t, ok, `Token.(scanTok)`)

	return logicalConjunction(
		assert.Equal(t, exp.text, actTk.text, `scanTok.text`),
		assert.Equal(t, exp.line, actTk.line, `scanTok.line`),
		assert.Equal(t, exp.start, actTk.start, `scanTok.start`),
		assert.Equal(t, exp.end, actTk.end, `scanTok.end`),
		assert.Equal(t, exp.kind, actTk.kind, `scanTok.kind`),
	)
}

func logicalConjunction(operands ...bool) bool {
	for _, b := range operands {
		if b == false {
			return false
		}
	}
	return true
}
