package strimmer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// strimTestID is used so target printing for specific test debugging.
var strimTestID int

func doTestStrim(t *testing.T, in, exp []token.Token) {

	act := []token.Token{}
	tt := token.TT_UNDEFINED

	for _, inTk := range in {
		outTk := Strim(inTk, tt)
		if outTk != nil {
			act = append(act, *outTk)
		}
		tt = inTk.Type
	}

	token.AssertTokens(t, exp, act)
}

func TestStrimmer_1(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.DummyToken(0, 5, 6, `1`, token.TT_NUMBER),
	}
	exp := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 5, 6, `1`, token.TT_NUMBER),
	}
	doTestStrim(t, in, exp)
}

func TestStrimmer_2(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 31, `// There's a snake in my boot`, token.TT_COMMENT),
	}
	exp := []token.Token{}
	doTestStrim(t, in, exp)
}

func TestStrimmer_3(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.DummyToken(0, 5, 6, `2`, token.TT_NUMBER),
		token.DummyToken(0, 6, 7, ` `, token.TT_SPACE),
		token.DummyToken(0, 7, 38, `// 'There's a snake in my boot'`, token.TT_COMMENT),
	}
	exp := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 5, 6, `2`, token.TT_NUMBER),
	}
	doTestStrim(t, in, exp)
}

func TestStrimmer_4(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 5, 20, `"Howdy partner"`, token.TT_STRING),
	}
	exp := []token.Token{
		token.DummyToken(0, 5, 20, `Howdy partner`, token.TT_STRING),
	}
	doTestStrim(t, in, exp)
}

func TestStrimmer_5(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 7, `123_456`, token.TT_NUMBER),
	}
	exp := []token.Token{
		token.DummyToken(0, 0, 7, `123456`, token.TT_NUMBER),
	}
	doTestStrim(t, in, exp)
}

func TestStrimmer_6(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 21, `1__2__3__.__4__5__6__`, token.TT_NUMBER),
	}
	exp := []token.Token{
		token.DummyToken(0, 0, 21, `123.456`, token.TT_NUMBER),
	}
	doTestStrim(t, in, exp)
}

func TestStrimmer_7(t *testing.T) {
	in := []token.Token{
		// f <- fUnC()
		//   @PrintKanji(語)
		// DONE
		token.DummyToken(0, 0, 1, `f`, token.TT_ID),
		token.DummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.DummyToken(0, 5, 9, `fUnC`, token.TT_FUNC),
		token.DummyToken(0, 9, 10, `(`, token.TT_CURVY_OPEN),
		token.DummyToken(0, 10, 11, `)`, token.TT_CURVY_CLOSE),
		token.DummyToken(0, 11, 12, "\n", token.TT_NEWLINE),
		token.DummyToken(0, 12, 23, `@PrintKanji`, token.TT_SPELL),
		token.DummyToken(0, 23, 24, `(`, token.TT_CURVY_OPEN),
		token.DummyToken(0, 24, 25, `語`, token.TT_ID),
		token.DummyToken(0, 25, 26, `)`, token.TT_CURVY_CLOSE),
		token.DummyToken(0, 26, 27, "\n", token.TT_NEWLINE),
		token.DummyToken(0, 27, 31, `DONE`, token.TT_DONE),
	}
	exp := []token.Token{
		token.DummyToken(0, 0, 1, `f`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 5, 9, `func`, token.TT_FUNC),
		token.DummyToken(0, 9, 10, `(`, token.TT_CURVY_OPEN),
		token.DummyToken(0, 10, 11, `)`, token.TT_CURVY_CLOSE),
		token.DummyToken(0, 11, 12, "\n", token.TT_EOS),
		token.DummyToken(0, 12, 23, `@printkanji`, token.TT_SPELL),
		token.DummyToken(0, 23, 24, `(`, token.TT_CURVY_OPEN),
		token.DummyToken(0, 24, 25, `語`, token.TT_ID),
		token.DummyToken(0, 25, 26, `)`, token.TT_CURVY_CLOSE),
		token.DummyToken(0, 26, 27, "\n", token.TT_EOS),
		token.DummyToken(0, 27, 31, `done`, token.TT_DONE),
	}
	doTestStrim(t, in, exp)
}

func TestStrimmer_8(t *testing.T) {
	in := []token.Token{
		// x <- [
		//   1,
		//   2, 3,
		// ]
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 1, 2, ` `, token.TT_SPACE),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.DummyToken(0, 5, 6, `[`, token.TT_SQUARE_OPEN),
		token.DummyToken(0, 6, 7, "\n", token.TT_NEWLINE),
		token.DummyToken(0, 0, 2, `  `, token.TT_SPACE),
		token.DummyToken(0, 2, 3, `1`, token.TT_NUMBER),
		token.DummyToken(0, 3, 4, "\n", token.TT_NEWLINE),
		token.DummyToken(0, 0, 2, `  `, token.TT_SPACE),
		token.DummyToken(0, 2, 3, `2`, token.TT_NUMBER),
		token.DummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
		token.DummyToken(0, 4, 5, ` `, token.TT_SPACE),
		token.DummyToken(0, 5, 6, `3`, token.TT_NUMBER),
		token.DummyToken(0, 6, 7, "\n", token.TT_NEWLINE),
		token.DummyToken(0, 0, 1, `]`, token.TT_SQUARE_CLOSE),
		token.DummyToken(0, 1, 2, "\n", token.TT_NEWLINE),
	}
	exp := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 5, 6, `[`, token.TT_SQUARE_OPEN),
		token.DummyToken(0, 2, 3, `1`, token.TT_NUMBER),
		token.DummyToken(0, 3, 4, "\n", token.TT_EOS),
		token.DummyToken(0, 2, 3, `2`, token.TT_NUMBER),
		token.DummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
		token.DummyToken(0, 5, 6, `3`, token.TT_NUMBER),
		token.DummyToken(0, 6, 7, "\n", token.TT_EOS),
		token.DummyToken(0, 0, 1, `]`, token.TT_SQUARE_CLOSE),
		token.DummyToken(0, 1, 2, "\n", token.TT_EOS),
	}
	doTestStrim(t, in, exp)
}
