package new_scanner

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummyRuner(s string) *Runer {
	sr := strings.NewReader(s)
	br := bufio.NewReader(sr)
	return NewRuner(br)
}

func readRequireNoErr(t *testing.T, r *Runer) rune {
	ru, err := r.ReadRune()
	require.Nil(t, err)
	return ru
}

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

func TestRuner_PeekRunes(t *testing.T) {
	r := dummyRuner(`abc`)

	ru1, ru2, err := r.PeekRunes()
	require.Nil(t, err)
	assert.Equal(t, 'a', ru1)
	assert.Equal(t, 'b', ru2)

	readRequireNoErr(t, r)
	readRequireNoErr(t, r)

	ru3, ru4, err := r.PeekRunes()
	require.Nil(t, err)
	assert.Equal(t, 'c', ru3)
	assert.Equal(t, EOF, ru4)
}

func TestRuner_Line(t *testing.T) {
	r := dummyRuner("a\nb\nc\nd")

	assert.Equal(t, 0, r.Line())

	_ = readRequireNoErr(t, r)
	_ = readRequireNoErr(t, r)

	assert.Equal(t, 1, r.Line())

	_ = readRequireNoErr(t, r)
	_ = readRequireNoErr(t, r)

	assert.Equal(t, 2, r.Line())

	_ = readRequireNoErr(t, r)
	_ = readRequireNoErr(t, r)

	assert.Equal(t, 3, r.Line())
}
