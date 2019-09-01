package farm

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func doTestFarm_Feed(
	t *testing.T,
	in token.Token,
	added, err bool,
	f *Farm,
	multiline, mature bool) {

	size := len(f.tokens)

	m, e := f.Feed(in)
	if err {
		require.NotNil(t, e, `Farm.Feed(): Was error expected?`)
		return
	}

	require.Nil(t, e, `Farm.Feed(): Was error expected?`)
	assert.Equal(t, mature, m, `Farm.Feed(): mature?`)

	if added {
		size++
	}

	assert.Equal(t, multiline, f.multiline, `Farm.multiline`)
	assert.Equal(t, mature, f.mature, `Farm.mature`)
	assert.Equal(t, size, len(f.tokens),
		`Farm.tokens: len(exp) == len(act)`)
}

func TestFarm_Feed_1(t *testing.T) {
	f := New()

	tk := scantok.New(`x`, 0, 0, 1, token.TT_ID)
	doTestFarm_Feed(t, tk, true, false, f, false, false)

	tk = scantok.New(`:`, 0, 1, 2, token.TT_ASSIGN)
	doTestFarm_Feed(t, tk, true, false, f, false, false)

	tk = scantok.New(` `, 0, 2, 3, token.TT_SPACE)
	doTestFarm_Feed(t, tk, false, false, f, false, false)

	tk = scantok.New(`1`, 0, 3, 4, token.TT_NUMBER)
	doTestFarm_Feed(t, tk, true, false, f, false, false)

	tk = scantok.New("\n", 0, 4, 5, token.TT_NEWLINE)
	doTestFarm_Feed(t, tk, false, false, f, false, true)

	tk = scantok.New("y", 0, 0, 1, token.TT_ID)
	assert.Panics(t, func() {
		doTestFarm_Feed(t, tk, false, false, f, false, false)
	})
}
