package runer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertRuner(t *testing.T, exp *Runer, act *Runer) bool {
	return logicalConjunction(
		assert.Equal(t, exp.line, act.line, `Runer.line`),
		assert.Equal(t, exp.col, act.col, `Runer.col`),
		assert.Equal(t, exp.newline, act.newline, `Runer.newline`),
		assert.Equal(t, exp.eof, act.eof, `Runer.eof`),
		assert.Equal(t, exp.buf1, act.buf1, `Runer.buf1`),
		assert.Equal(t, exp.buf1_eof, act.buf1_eof, `Runer.buf1_eof`),
		assert.Equal(t, exp.buf2, act.buf2, `Runer.buf2`),
		assert.Equal(t, exp.buf2_eof, act.buf2_eof, `Runer.buf2_eof`),
	)
}

func logicalConjunction(operands ...bool) bool {
	for _, b := range operands {
		if b == false {
			return false
		}
	}
	return true
}

func readReqNoErr(t *testing.T, r *Runer) (rune, bool) {
	ru, eof, err := r.Read()
	require.Nil(t, err, `Unexpected Runer error`)
	return ru, eof
}

func doTestRead(t *testing.T, expRu rune, expEOF bool, exp *Runer, r *Runer) bool {
	ru, eof := readReqNoErr(t, r)

	return logicalConjunction(
		assert.Equal(t, expRu, ru, `Runer.Read(): Unexpected rune`),
		assert.Equal(t, expEOF, eof, `Runer.Read(): Unexpected EOF`),
		assertRuner(t, exp, r),
	)
}

func TestRuner_Read(t *testing.T) {
	r := NewByStr("ab\ncd")

	exp := &Runer{nil, 0, 0, false, false, 'b', false, '\n', false}
	doTestRead(t, 'a', false, exp, r)

	exp = &Runer{nil, 0, 1, false, false, '\n', false, 'c', false}
	doTestRead(t, 'b', false, exp, r)

	exp = &Runer{nil, 0, 2, true, false, 'c', false, 'd', false}
	doTestRead(t, '\n', false, exp, r)

	exp = &Runer{nil, 1, 0, false, false, 'd', false, 0, true}
	doTestRead(t, 'c', false, exp, r)

	exp = &Runer{nil, 1, 1, false, false, 0, true, 0, true}
	doTestRead(t, 'd', false, exp, r)

	exp = &Runer{nil, 1, 2, false, true, 0, true, 0, true}
	doTestRead(t, 0, true, exp, r)
}

func TestRuner_ReadIf(t *testing.T) {
	r := NewByStr(`123`)

	f := func(ru1 rune, ru2 rune) (bool, error) {
		a := ru1 > '0' && ru2 != 0
		return a, nil
	}

	ru, read, e := r.ReadIf(f)
	require.Nil(t, e)
	assert.True(t, read)
	assert.Equal(t, '1', ru)

	ru, read, e = r.ReadIf(f)
	require.Nil(t, e)
	assert.True(t, read)
	assert.Equal(t, '2', ru)

	ru, read, e = r.ReadIf(f)
	require.Nil(t, e)
	assert.False(t, read)
	assert.Equal(t, rune(0), ru)
}

func TestRuner_ReadWhile(t *testing.T) {
	r := NewByStr(`abc 123`)

	f := func(ru, _ rune) (bool, error) {
		return ru != ' ', nil
	}

	s, e := r.ReadWhile(f)
	require.Nil(t, e)
	assert.Equal(t, `abc`, s)

	eof, e := r.Skip()
	require.Nil(t, e)
	require.False(t, eof)

	s, e = r.ReadWhile(f)
	require.Nil(t, e)
	assert.Equal(t, `123`, s)

	eof, e = r.Skip()
	require.Nil(t, e)
	require.True(t, eof)

	s, e = r.ReadWhile(f)
	require.Nil(t, e)
	assert.Equal(t, ``, s)

	eof, e = r.Skip()
	require.Nil(t, e)
	require.True(t, eof)
}
