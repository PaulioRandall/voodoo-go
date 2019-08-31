package token

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertSliceEqual asserts that the slice of actual tokens matches the slice
// of expected tokens. Note that the only the responses from the token interface
// methods are checked, the underlying type and its fields are not considered.
func AssertSliceEqual(t *testing.T, exp []Token, act []Token) {
	for i, _ := range exp {
		if !assert.True(t, i < len(act), `Token[%d] missing`, i) {
			break
		}

		AssertEqual(t, exp[i], act[i])
	}

	assert.Equal(t, len(exp), len(act), `len(exp) == len(act)`)
}

// AssertEqual asserts that the actual token equals the expected token. Note
// that the only the responses from the token interface methods are checked, the
// underlying type and its fields are not considered.
func AssertEqual(t *testing.T, exp Token, act Token) bool {
	if exp == nil {
		return assert.Nil(t, act, `Token == nil`)
	}

	require.NotNil(t, act, `Token != nil`)

	return logicalConjunction(
		assert.Equal(t, exp.Text(), act.Text(), `Token.Text()`),
		assert.Equal(t, exp.Kind(), act.Kind(), `Token.Kind()`),
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

// Dummy represents a dummy token.
type Dummy struct {
	T string
	K Kind
}

// Text satisfies the Token interface.
func (d Dummy) Text() string {
	return d.T
}

// Kind satisfies the Token interface.
func (d Dummy) Kind() Kind {
	return d.K
}

// String satisfies the Token interface.
func (d Dummy) String() string {
	return `(` + KindName(d.K) + `) ` + d.T
}

// DummyOfKind creates a new dummy token initialised to the specified token
// kind.
func DummyOfKind(k Kind) Token {
	return Dummy{
		K: k,
	}
}

// UniqueDummy creates a new dummy token initialised to the specified token
// kind and unique text.
func UniqueDummy(k Kind) Token {
	return Dummy{
		T: strconv.FormatUint(rand.Uint64(), 10),
		K: k,
	}
}
