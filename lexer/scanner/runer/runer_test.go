package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanWordStr_1(t *testing.T) {
	in := []rune(`Happi_123_ness`)
	word, out := scanWordStr(in)
	assert.Equal(t, `Happi_123_ness`, word)
	assert.Nil(t, out)
}

func TestScanWordStr_2(t *testing.T) {
	in := []rune(`Happi ness`)
	word, out := scanWordStr(in)
	assert.Equal(t, `Happi`, word)
	assert.Equal(t, []rune(` ness`), out)
}
