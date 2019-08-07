package new_scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanWordStr(t *testing.T) {
	r := dummyRuner(`abc`)
	out, err := scanWordStr(r)
	require.Nil(t, err)
	assert.Equal(t, `abc`, out)
}
