package token

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertEqual asserts that the the actual Token is a scanTok and that
// it equals the expected Token which must also be a scanTok.
func AssertEqual(t *testing.T, exp, act Token) bool {
	if exp == nil {
		return assert.Nil(t, act, `act == nil`)
	}

	require.NotNil(t, act, `act != nil`)

	return utils.LogicalConjunction(
		assert.Equal(t, exp.Text(), act.Text(), `token.Text()`),
		assert.Equal(t, exp.Line(), act.Line(), `token.Line()`),
		assert.Equal(t, exp.Start(), act.Start(), `token.Start()`),
		assert.Equal(t, exp.End(), act.End(), `token.End()`),
		assert.Equal(t,
			KindName(exp.Kind()),
			KindName(act.Kind()),
			`token.Kind()`,
		),
	)
}

// AssertSliceEqual asserts that the the actual Token slice contains only
// scanTok instances and that each equals the corresponding one from the
// expected Token slice.
func AssertSliceEqual(t *testing.T, exp, act []Token) bool {
	if exp == nil {
		return assert.Nil(t, act, `act == nil`)
	}

	require.NotNil(t, act, `act != nil`)

	ok := true
	for i, _ := range exp {
		require.True(t, i < len(act), `exp[i]: i < len(act)`)
		ok = ok && AssertEqual(t, exp[i], act[i])
	}

	return ok && assert.Equal(t, len(exp), len(act), `len(exp) == len(act)`)
}

// UniqueDummy creates a new dummy token initialised to the specified token
// kind and unique text.
func UniqueDummy(k Kind) Token {
	return token{
		text: strconv.FormatUint(rand.Uint64(), 10),
		kind: k,
	}
}
