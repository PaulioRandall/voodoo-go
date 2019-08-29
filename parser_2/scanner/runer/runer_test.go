package runer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertRead(t *testing.T, expRu rune, expEOF bool, exp *Runer, r *Runer) bool {
	ru, eof := readReqNoErr(t, r)

	return logicalConjunction(
		assert.Equal(t, expRu, ru, `Runer.Read(): Unexpected rune`),
		assert.Equal(t, expEOF, eof, `Runer.Read(): Unexpected EOF`),

		assert.Equal(t, exp.line, r.line, `Runer.line`),
		assert.Equal(t, exp.col, r.col, `Runer.col`),
		assert.Equal(t, exp.newline, r.newline, `Runer.newline`),
		assert.Equal(t, exp.eof, r.eof, `Runer.eof`),
		assert.Equal(t, exp.buf, r.buf, `Runer.buf`),
		assert.Equal(t, exp.bufEOF, r.bufEOF, `Runer.bufEOF`),
	)
}

func readReqNoErr(t *testing.T, r *Runer) (rune, bool) {
	ru, eof, err := r.Read()
	require.Nil(t, err, `Unexpected Runer error`)
	return ru, eof
}

func logicalConjunction(operands ...bool) bool {
	for _, b := range operands {
		if b == false {
			return false
		}
	}
	return true
}

func TestRuner_Read(t *testing.T) {
	r := NewByStr("ab\ncd")

	exp := &Runer{
		buf: 'b',
	}
	assertRead(t, 'a', false, exp, r)

	exp = &Runer{
		col: 1,
		buf: '\n',
	}
	assertRead(t, 'b', false, exp, r)

	exp = &Runer{
		col:     2,
		newline: true,
		buf:     'c',
	}
	assertRead(t, '\n', false, exp, r)

	exp = &Runer{
		line: 1,
		col:  0,
		buf:  'd',
	}
	assertRead(t, 'c', false, exp, r)

	exp = &Runer{
		line:   1,
		col:    1,
		buf:    0,
		bufEOF: true,
	}
	assertRead(t, 'd', false, exp, r)

	exp = &Runer{
		line:   1,
		col:    2,
		eof:    true,
		buf:    0,
		bufEOF: true,
	}
	assertRead(t, 0, true, exp, r)
}
