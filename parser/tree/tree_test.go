package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Is(t *testing.T) {
	tr := New()
	tr.Kind = KD_ID
	assert.True(t, tr.Is(KD_ID))
	assert.False(t, tr.Is(KD_OPERAND))
}

func TestTree_IsLeft(t *testing.T) {
	tr := New()
	assert.True(t, tr.IsLeft(KD_UNDEFINED))
	assert.False(t, tr.IsLeft(KD_ID))

	tr.Left = New()
	tr.Left.Kind = KD_ID
	assert.False(t, tr.IsLeft(KD_UNDEFINED))
	assert.True(t, tr.IsLeft(KD_ID))
	assert.False(t, tr.IsLeft(KD_OPERAND))
}

func TestTree_IsRight(t *testing.T) {
	tr := New()
	assert.True(t, tr.IsRight(KD_UNDEFINED))
	assert.False(t, tr.IsRight(KD_ID))

	tr.Right = New()
	tr.Right.Kind = KD_ID
	assert.False(t, tr.IsRight(KD_UNDEFINED))
	assert.True(t, tr.IsRight(KD_ID))
	assert.False(t, tr.IsRight(KD_OPERAND))
}
