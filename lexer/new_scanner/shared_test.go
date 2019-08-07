package new_scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanWordStr_1(t *testing.T) {
	r := dummyRuner(`ab_3`)
	out, size, err := scanWordStr(r)
	require.Nil(t, err)
	assert.Equal(t, `ab_3`, out)
	assert.Equal(t, 4, size)
	assert.Equal(t, EOF, readRequireNoErr(t, r))
}

func TestScanWordStr_2(t *testing.T) {
	r := dummyRuner(`ab cd`)
	out, size, err := scanWordStr(r)
	require.Nil(t, err)
	assert.Equal(t, `ab`, out)
	assert.Equal(t, 2, size)
	assert.Equal(t, ' ', readRequireNoErr(t, r))
}
