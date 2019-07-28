package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanInt_1(t *testing.T) {
	in := []rune(`_1_2_3_`)
	act, out := scanInt(in)
	assert.Equal(t, `_1_2_3_`, act)
	assert.Nil(t, out)
}

func TestScanInt_2(t *testing.T) {
	in := []rune(`123.456`)
	act, out := scanInt(in)
	assert.Equal(t, `123`, act)
	assert.Equal(t, []rune(`.456`), out)
}

func TestScanWordStr_1(t *testing.T) {
	in := []rune(`Happi_123_ness`)
	act, out := scanWordStr(in)
	assert.Equal(t, `Happi_123_ness`, act)
	assert.Nil(t, out)
}

func TestScanWordStr_2(t *testing.T) {
	in := []rune(`Happi ness`)
	act, out := scanWordStr(in)
	assert.Equal(t, `Happi`, act)
	assert.Equal(t, []rune(` ness`), out)
}
