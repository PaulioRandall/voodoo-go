package preparser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func doTestAdd(t *testing.T, stat *Statement, in, exp *token.Token, expComplete bool) {
	i := len(stat.Tokens)
	complete := stat.Add(in)

	assert.Equal(t, expComplete, complete)

	if exp == nil {
		assert.Equal(t, i, len(stat.Tokens))
	} else {
		require.Equal(t, i+1, len(stat.Tokens))
		token.AssertToken(t, exp, &stat.Tokens[i])
	}
}

func doTestStatement(t *testing.T, in, exp []*token.Token, expComplete bool) {
	require.Equal(t, len(in), len(exp))
	last := len(in) - 1
	stat := New()

	for i := 0; i <= last; i++ {
		expFinished := expComplete && last == i
		doTestAdd(t, stat, in[i], exp[i], expFinished)
	}
}

func TestStatement_Add_1(t *testing.T) {
	stat := New()
	in := token.DummyToken(0, 0, 1, `x`, token.TT_ID)
	exp := token.DummyToken(0, 0, 1, `x`, token.TT_ID)
	doTestAdd(t, stat, &in, &exp, false)
}

func TestStatement_Add_2(t *testing.T) {
	stat := New()

	in := token.DummyToken(0, 0, 1, `x`, token.TT_ID)
	exp := token.DummyToken(0, 0, 1, `x`, token.TT_ID)
	doTestAdd(t, stat, &in, &exp, false)

	in = token.DummyToken(0, 1, 2, "\n", token.TT_EOS)
	doTestAdd(t, stat, &in, nil, true)
}

func TestStatement_Add_3(t *testing.T) {
	stat := New()

	in := token.DummyToken(0, 0, 1, `x`, token.TT_ID)
	exp := token.DummyToken(0, 0, 1, `x`, token.TT_ID)
	doTestAdd(t, stat, &in, &exp, false)

	in = token.DummyToken(0, 1, 2, "\n", token.TT_EOS)
	doTestAdd(t, stat, &in, nil, true)

	assert.Panics(t, func() {
		in = token.DummyToken(0, 2, 3, `y`, token.TT_ID)
		doTestAdd(t, stat, &in, nil, false)
	})
}

func TestStatement_Add_4(t *testing.T) {
	in := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		token.PtrDummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 6, `1`, token.TT_NUMBER),
	}
	exp := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(0, 5, 6, `1`, token.TT_NUMBER),
	}
	doTestStatement(t, in, exp, false)
}

func TestStatement_Add_5(t *testing.T) {
	in := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		token.PtrDummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 6, `1`, token.TT_NUMBER),
		token.PtrDummyToken(0, 6, 7, "\n", token.TT_EOS),
	}
	exp := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(0, 5, 6, `1`, token.TT_NUMBER),
		nil,
	}
	doTestStatement(t, in, exp, true)
}
