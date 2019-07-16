package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneItr_Length(t *testing.T) {
	s := `Cry Out For A Hero`
	itr := NewRuneItr(s)
	assert.Equal(t, len(s), itr.Length())
}

func TestRuneItr_HasRelRune(t *testing.T) {
	s := `Cry Out For A Hero`
	itr := NewRuneItr(s)
	itr.index = 2

	assert.True(t, itr.HasRelRune(0))
	assert.True(t, itr.HasRelRune(1))
	assert.True(t, itr.HasRelRune(-1))

	assert.False(t, itr.HasRelRune(-3))
	assert.False(t, itr.HasRelRune(20))
}

func TestRuneItr_RelRune(t *testing.T) {
	s := `Cry Out For A Hero`
	itr := NewRuneItr(s)
	itr.index = 2

	assert.Equal(t, 'y', itr.RelRune(0))
	assert.Equal(t, ' ', itr.RelRune(1))
	assert.Equal(t, 'r', itr.RelRune(-1))

	assert.Equal(t, int32(-1), itr.RelRune(-3))
	assert.Equal(t, int32(-1), itr.RelRune(20))
}

func TestRuneItr_NextRune(t *testing.T) {
	s := `abc`
	itr := NewRuneItr(s)

	assert.Equal(t, 'a', itr.NextRune())
	assert.Equal(t, 'b', itr.NextRune())
	assert.Equal(t, 'c', itr.NextRune())
	assert.Equal(t, int32(-1), itr.NextRune())
	assert.Equal(t, int32(-1), itr.NextRune())
}
