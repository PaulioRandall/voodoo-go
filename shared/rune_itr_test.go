package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneItr_Length(t *testing.T) {
	s := `Cry Out For A Hero`
	itr := NewRuneItr(s)

	exp := len(s)
	act := itr.Length()

	assert.Equal(t, exp, act)
}
