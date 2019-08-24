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

func TestStatement_Add_6(t *testing.T) {
	in := []*token.Token{
		token.PtrDummyToken(0, 0, 31, `// There's a snake in my boot`, token.TT_COMMENT),
	}
	exp := []*token.Token{
		nil,
	}
	doTestStatement(t, in, exp, false)
}

func TestStatement_Add_7(t *testing.T) {
	in := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		token.PtrDummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 6, `2`, token.TT_NUMBER),
		token.PtrDummyToken(0, 6, 7, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 7, 38, `// 'There's a snake in my boot'`, token.TT_COMMENT),
		token.PtrDummyToken(0, 38, 39, "\n", token.TT_EOS),
	}
	exp := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(0, 5, 6, `2`, token.TT_NUMBER),
		nil,
		nil,
		nil,
	}
	doTestStatement(t, in, exp, true)
}

func TestStatement_Add_8(t *testing.T) {
	in := []*token.Token{
		token.PtrDummyToken(0, 5, 20, `"Howdy partner"`, token.TT_STRING),
	}
	exp := []*token.Token{
		token.PtrDummyToken(0, 5, 20, `Howdy partner`, token.TT_STRING),
	}
	doTestStatement(t, in, exp, false)
}

func TestStatement_Add_9(t *testing.T) {
	// f <- fUnC()
	//   @PrintKanji(語)
	// DONE

	in := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `f`, token.TT_ID),
		token.PtrDummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 9, `fUnC`, token.TT_FUNC),
		token.PtrDummyToken(0, 9, 10, `(`, token.TT_CURVY_OPEN),
		token.PtrDummyToken(0, 10, 11, `)`, token.TT_CURVY_CLOSE),
		token.PtrDummyToken(0, 11, 12, "\n", token.TT_NEWLINE),
	}
	exp := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `f`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(0, 5, 9, `func`, token.TT_FUNC),
		token.PtrDummyToken(0, 9, 10, `(`, token.TT_CURVY_OPEN),
		token.PtrDummyToken(0, 10, 11, `)`, token.TT_CURVY_CLOSE),
		nil,
	}
	doTestStatement(t, in, exp, true)

	in = []*token.Token{
		token.PtrDummyToken(0, 12, 23, `@PrintKanji`, token.TT_SPELL),
		token.PtrDummyToken(0, 23, 24, `(`, token.TT_CURVY_OPEN),
		token.PtrDummyToken(0, 24, 25, `語`, token.TT_ID),
		token.PtrDummyToken(0, 25, 26, `)`, token.TT_CURVY_CLOSE),
		token.PtrDummyToken(0, 26, 27, "\n", token.TT_NEWLINE),
	}
	exp = []*token.Token{
		token.PtrDummyToken(0, 12, 23, `@printkanji`, token.TT_SPELL),
		token.PtrDummyToken(0, 23, 24, `(`, token.TT_CURVY_OPEN),
		token.PtrDummyToken(0, 24, 25, `語`, token.TT_ID),
		token.PtrDummyToken(0, 25, 26, `)`, token.TT_CURVY_CLOSE),
		nil,
	}
	doTestStatement(t, in, exp, true)

	in = []*token.Token{
		token.PtrDummyToken(0, 27, 31, `DONE`, token.TT_DONE),
		token.PtrDummyToken(0, 31, 32, "\n", token.TT_NEWLINE),
	}
	exp = []*token.Token{
		token.PtrDummyToken(0, 27, 31, `done`, token.TT_DONE),
		nil,
	}
	doTestStatement(t, in, exp, true)
}

func TestStatement_Add_10(t *testing.T) {
	// x <- [
	//   1,
	//   2, 3,
	// ]

	in := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		token.PtrDummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 6, `[`, token.TT_SQUARE_OPEN),
		token.PtrDummyToken(0, 6, 7, "\n", token.TT_NEWLINE),
		token.PtrDummyToken(0, 0, 2, `  `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 3, `1`, token.TT_NUMBER),
		token.PtrDummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
		token.PtrDummyToken(0, 4, 5, "\n", token.TT_NEWLINE),
		token.PtrDummyToken(0, 0, 2, `  `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 3, `2`, token.TT_NUMBER),
		token.PtrDummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 6, `3`, token.TT_NUMBER),
		token.PtrDummyToken(0, 6, 7, `,`, token.TT_VALUE_DELIM),
		token.PtrDummyToken(0, 7, 8, "\n", token.TT_NEWLINE),
		token.PtrDummyToken(0, 0, 1, `]`, token.TT_SQUARE_CLOSE),
		token.PtrDummyToken(0, 1, 2, "\n", token.TT_NEWLINE),
	}
	exp := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `x`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(0, 5, 6, `[`, token.TT_SQUARE_OPEN),
		nil,
		nil,
		token.PtrDummyToken(0, 2, 3, `1`, token.TT_NUMBER),
		token.PtrDummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
		nil,
		nil,
		token.PtrDummyToken(0, 2, 3, `2`, token.TT_NUMBER),
		token.PtrDummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
		nil,
		token.PtrDummyToken(0, 5, 6, `3`, token.TT_NUMBER),
		token.PtrDummyToken(0, 6, 7, `,`, token.TT_VALUE_DELIM),
		nil,
		token.PtrDummyToken(0, 0, 1, `]`, token.TT_SQUARE_CLOSE),
		nil,
	}
	doTestStatement(t, in, exp, true)
}
