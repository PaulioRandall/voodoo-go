package preparser

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func doTestAdd(t *testing.T, stat *Statement, in, exp *token.Token, expComplete bool) {
	i := len(stat.Tokens)
	complete := Add(stat, in)

	v := strconv.QuoteToGraphic(in.Val)
	assert.Equal(t, expComplete, complete, "%s cause statement completion?", v)

	if exp == nil {
		assert.Equal(t, i, len(stat.Tokens), "Unexpected removal of token[%d]", i)
	} else {
		require.Equal(t, i+1, len(stat.Tokens), "Unexpected additional token[%d]", i)
		token.AssertToken(t, exp, &stat.Tokens[i])
	}
}

func doTestStatement(t *testing.T, in, exp []*token.Token, expComplete bool) {
	require.Equal(t, len(in), len(exp), `Test requirment: len(in) == len(exp)`)
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

	in := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `f`, token.TT_ID),
		token.PtrDummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 9, `fUnC`, token.TT_FUNC),
		token.PtrDummyToken(0, 9, 10, `(`, token.TT_CURVED_OPEN),
		token.PtrDummyToken(0, 10, 11, `)`, token.TT_CURVED_CLOSE),
		token.PtrDummyToken(0, 11, 12, "\n", token.TT_NEWLINE),
	}
	exp := []*token.Token{
		token.PtrDummyToken(0, 0, 1, `f`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(0, 5, 9, `func`, token.TT_FUNC),
		token.PtrDummyToken(0, 9, 10, `(`, token.TT_CURVED_OPEN),
		token.PtrDummyToken(0, 10, 11, `)`, token.TT_CURVED_CLOSE),
		nil,
	}
	doTestStatement(t, in, exp, true)

	in = []*token.Token{
		token.PtrDummyToken(0, 12, 23, `@PrintKanji`, token.TT_SPELL),
		token.PtrDummyToken(0, 23, 24, `(`, token.TT_CURVED_OPEN),
		token.PtrDummyToken(0, 24, 25, `語`, token.TT_ID),
		token.PtrDummyToken(0, 25, 26, `)`, token.TT_CURVED_CLOSE),
		token.PtrDummyToken(0, 26, 27, "\n", token.TT_NEWLINE),
	}
	exp = []*token.Token{
		token.PtrDummyToken(0, 12, 23, `@printkanji`, token.TT_SPELL),
		token.PtrDummyToken(0, 23, 24, `(`, token.TT_CURVED_OPEN),
		token.PtrDummyToken(0, 24, 25, `語`, token.TT_ID),
		token.PtrDummyToken(0, 25, 26, `)`, token.TT_CURVED_CLOSE),
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

func TestStatement_Add_11(t *testing.T) {
	// f <- func(a, b) r, err {
	// 	 b == 0 => err := "Can't divide by zero"
	//   b != 0 => r := a / b
	// }
	in := []*token.Token{
		// f <- func(a, b) r, err {
		token.PtrDummyToken(0, 0, 1, `f`, token.TT_ID),
		token.PtrDummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.PtrDummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 5, 9, `func`, token.TT_FUNC),
		token.PtrDummyToken(0, 9, 10, `(`, token.TT_CURVED_OPEN),
		token.PtrDummyToken(0, 10, 11, `a`, token.TT_ID),
		token.PtrDummyToken(0, 11, 12, `,`, token.TT_VALUE_DELIM),
		token.PtrDummyToken(0, 12, 13, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 13, 14, `b`, token.TT_ID),
		token.PtrDummyToken(0, 14, 15, `)`, token.TT_CURVED_CLOSE),
		token.PtrDummyToken(0, 15, 16, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 16, 17, `r`, token.TT_ID),
		token.PtrDummyToken(0, 17, 18, `,`, token.TT_VALUE_DELIM),
		token.PtrDummyToken(0, 18, 19, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 19, 22, `err`, token.TT_ID),
		token.PtrDummyToken(0, 22, 23, ` `, token.TT_SPACE),
		token.PtrDummyToken(0, 23, 24, `{`, token.TT_CURLY_OPEN),
		token.PtrDummyToken(0, 24, 25, "\n", token.TT_NEWLINE),
	}
	exp := []*token.Token{
		// f <- func(a, b) r, err {
		token.PtrDummyToken(0, 0, 1, `f`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(0, 5, 9, `func`, token.TT_FUNC),
		token.PtrDummyToken(0, 9, 10, `(`, token.TT_CURVED_OPEN),
		token.PtrDummyToken(0, 10, 11, `a`, token.TT_ID),
		token.PtrDummyToken(0, 11, 12, `,`, token.TT_VALUE_DELIM),
		nil,
		token.PtrDummyToken(0, 13, 14, `b`, token.TT_ID),
		token.PtrDummyToken(0, 14, 15, `)`, token.TT_CURVED_CLOSE),
		nil,
		token.PtrDummyToken(0, 16, 17, `r`, token.TT_ID),
		token.PtrDummyToken(0, 17, 18, `,`, token.TT_VALUE_DELIM),
		nil,
		token.PtrDummyToken(0, 19, 22, `err`, token.TT_ID),
		nil,
		token.PtrDummyToken(0, 23, 24, `{`, token.TT_CURLY_OPEN),
		nil,
	}
	doTestStatement(t, in, exp, true)

	in = []*token.Token{
		// 	 b == 0 => err := "Can't divide by zero"
		token.PtrDummyToken(1, 0, 1, "\t", token.TT_SPACE),
		token.PtrDummyToken(1, 1, 2, `b`, token.TT_ID),
		token.PtrDummyToken(1, 2, 3, ` `, token.TT_SPACE),
		token.PtrDummyToken(1, 3, 5, `==`, token.TT_CMP_EQ),
		token.PtrDummyToken(1, 5, 6, ` `, token.TT_SPACE),
		token.PtrDummyToken(1, 6, 7, `0`, token.TT_NUMBER),
		token.PtrDummyToken(1, 7, 8, ` `, token.TT_SPACE),
		token.PtrDummyToken(1, 8, 10, `=>`, token.TT_IF_THEN),
		token.PtrDummyToken(1, 10, 11, ` `, token.TT_SPACE),
		token.PtrDummyToken(1, 11, 14, `err`, token.TT_ID),
		token.PtrDummyToken(1, 14, 15, ` `, token.TT_SPACE),
		token.PtrDummyToken(1, 15, 17, `:=`, token.TT_ASSIGN),
		token.PtrDummyToken(1, 17, 18, ` `, token.TT_SPACE),
		token.PtrDummyToken(1, 18, 40, `"Can't divide by zero"`, token.TT_STRING),
		token.PtrDummyToken(1, 40, 41, "\n", token.TT_NEWLINE),
	}
	exp = []*token.Token{
		// 	 b == 0 => err := "Can't divide by zero"
		nil,
		token.PtrDummyToken(1, 1, 2, `b`, token.TT_ID),
		nil,
		token.PtrDummyToken(1, 3, 5, `==`, token.TT_CMP_EQ),
		nil,
		token.PtrDummyToken(1, 6, 7, `0`, token.TT_NUMBER),
		nil,
		token.PtrDummyToken(1, 8, 10, `=>`, token.TT_IF_THEN),
		nil,
		token.PtrDummyToken(1, 11, 14, `err`, token.TT_ID),
		nil,
		token.PtrDummyToken(1, 15, 17, `:=`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(1, 18, 40, `Can't divide by zero`, token.TT_STRING),
		nil,
	}
	doTestStatement(t, in, exp, true)

	in = []*token.Token{
		//   b != 0 => r := a / b
		token.PtrDummyToken(2, 0, 1, "\t", token.TT_SPACE),
		token.PtrDummyToken(2, 1, 2, `b`, token.TT_ID),
		token.PtrDummyToken(2, 2, 3, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 3, 5, `!=`, token.TT_CMP_NOT_EQ),
		token.PtrDummyToken(2, 5, 6, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 6, 7, `0`, token.TT_NUMBER),
		token.PtrDummyToken(2, 7, 8, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 8, 10, `=>`, token.TT_IF_THEN),
		token.PtrDummyToken(2, 10, 11, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 11, 12, `r`, token.TT_ID),
		token.PtrDummyToken(2, 12, 13, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 13, 15, `:=`, token.TT_ASSIGN),
		token.PtrDummyToken(2, 15, 16, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 16, 17, "a", token.TT_ID),
		token.PtrDummyToken(2, 17, 18, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 18, 19, "/", token.TT_DIVIDE),
		token.PtrDummyToken(2, 19, 20, ` `, token.TT_SPACE),
		token.PtrDummyToken(2, 20, 21, "b", token.TT_ID),
		token.PtrDummyToken(2, 21, 22, "\n", token.TT_NEWLINE),
	}
	exp = []*token.Token{
		//   b != 0 => r := a / b
		nil,
		token.PtrDummyToken(2, 1, 2, `b`, token.TT_ID),
		nil,
		token.PtrDummyToken(2, 3, 5, `!=`, token.TT_CMP_NOT_EQ),
		nil,
		token.PtrDummyToken(2, 6, 7, `0`, token.TT_NUMBER),
		nil,
		token.PtrDummyToken(2, 8, 10, `=>`, token.TT_IF_THEN),
		nil,
		token.PtrDummyToken(2, 11, 12, `r`, token.TT_ID),
		nil,
		token.PtrDummyToken(2, 13, 15, `:=`, token.TT_ASSIGN),
		nil,
		token.PtrDummyToken(2, 16, 17, "a", token.TT_ID),
		nil,
		token.PtrDummyToken(2, 18, 19, "/", token.TT_DIVIDE),
		nil,
		token.PtrDummyToken(2, 20, 21, "b", token.TT_ID),
		nil,
	}
	doTestStatement(t, in, exp, true)

	in = []*token.Token{
		// }
		token.PtrDummyToken(3, 0, 1, "}", token.TT_CURLY_CLOSE),
		token.PtrDummyToken(3, 1, 2, "\n", token.TT_NEWLINE),
	}
	exp = []*token.Token{
		// }
		token.PtrDummyToken(3, 0, 1, "}", token.TT_CURLY_CLOSE),
		nil,
	}
	doTestStatement(t, in, exp, true)
}
