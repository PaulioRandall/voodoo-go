package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneItr_Length(t *testing.T) {
	// Japanese rune has 3 bytes
	s := `Cry Out For A Hero 語`
	itr := NewRuneItr(s)
	exp := len([]rune(s))
	assert.Equal(t, exp, itr.Length())
}

func TestRuneItr_Index(t *testing.T) {
	s := `Cry Out For A Hero 語`
	itr := NewRuneItr(s)
	itr.index = 12
	assert.Equal(t, 12, itr.Index())
}

func TestRuneItr_HasRelRune(t *testing.T) {
	s := `Cry Out For A Hero 語`
	itr := NewRuneItr(s)
	itr.index = 2

	assert.True(t, itr.HasRelRune(0))
	assert.True(t, itr.HasRelRune(1))
	assert.True(t, itr.HasRelRune(-1))

	assert.False(t, itr.HasRelRune(-3))
	assert.False(t, itr.HasRelRune(20))
}

func TestRuneItr_PeekRelRune(t *testing.T) {
	s := `Cry Out For A Hero 語`
	itr := NewRuneItr(s)
	itr.index = 2

	assert.Equal(t, 'y', itr.PeekRelRune(0))
	assert.Equal(t, ' ', itr.PeekRelRune(1))
	assert.Equal(t, 'r', itr.PeekRelRune(-1))

	assert.Equal(t, int32(-1), itr.PeekRelRune(-3))
	assert.Equal(t, int32(-1), itr.PeekRelRune(20))
}

func TestRuneItr_NextRune(t *testing.T) {
	s := `ab語`
	itr := NewRuneItr(s)

	assert.Equal(t, 'a', itr.NextRune())
	assert.Equal(t, 'b', itr.NextRune())
	assert.Equal(t, '語', itr.NextRune())
	assert.Equal(t, int32(-1), itr.NextRune())
	assert.Equal(t, int32(-1), itr.NextRune())
}

func TestRuneItr_PeekRune(t *testing.T) {
	s := `ab語`
	itr := NewRuneItr(s)

	assert.Equal(t, 'a', itr.PeekRune())
	itr.index = 5
	assert.Equal(t, int32(-1), itr.PeekRune())
}

func TestRuneItr_HasNext(t *testing.T) {
	s := `ab語`
	itr := NewRuneItr(s)

	assert.True(t, itr.HasNext())
	itr.index += 1
	assert.True(t, itr.HasNext())
	itr.index += 1
	assert.True(t, itr.HasNext())
	itr.index += 1
	assert.False(t, itr.HasNext())
	itr.index += 1
	assert.False(t, itr.HasNext())
}

func TestRuneItr_IsNext(t *testing.T) {
	s := `ab語`
	itr := NewRuneItr(s)

	assert.True(t, itr.IsNext('a'))
	assert.False(t, itr.IsNext('b'))
	itr.index += 1
	assert.True(t, itr.IsNext('b'))
	assert.False(t, itr.IsNext('語'))
	itr.index += 1
	assert.True(t, itr.IsNext('語'))
	itr.index += 1
	assert.False(t, itr.IsNext('語'))
}

func TestRuneItr_IsNextIn(t *testing.T) {
	s := `ab語`
	itr := NewRuneItr(s)

	assert.True(t, itr.IsNextIn("ab語"))
	assert.False(t, itr.IsNextIn("xyz"))
	itr.index += 1
	assert.True(t, itr.IsNextIn("ab語"))
	assert.False(t, itr.IsNextIn("xyz"))
	itr.index += 1
	itr.index += 1
	assert.False(t, itr.IsNextIn("ab語"))
}
