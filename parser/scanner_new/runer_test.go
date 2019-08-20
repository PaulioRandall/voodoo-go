package scanner_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRuner_ReadRune(t *testing.T) {
	r := dummyRuner(`abc`)

	ru1 := readRequireNoErr(t, r)
	assert.Equal(t, 'a', ru1)

	ru2 := readRequireNoErr(t, r)
	assert.Equal(t, 'b', ru2)

	ru3 := readRequireNoErr(t, r)
	assert.Equal(t, 'c', ru3)

	ru4 := readRequireNoErr(t, r)
	assert.Equal(t, EOF, ru4)
}

func TestRuner_LookAhead(t *testing.T) {
	r := dummyRuner(`abc`)

	ru1, ru2, err := r.LookAhead()
	require.Nil(t, err)
	assert.Equal(t, 'a', ru1)
	assert.Equal(t, 'b', ru2)

	readRequireNoErr(t, r)

	ru3, ru4, err := r.LookAhead()
	require.Nil(t, err)
	assert.Equal(t, 'b', ru3)
	assert.Equal(t, 'c', ru4)

	readRequireNoErr(t, r)

	ru5, ru6, err := r.LookAhead()
	require.Nil(t, err)
	assert.Equal(t, 'c', ru5)
	assert.Equal(t, EOF, ru6)
}
