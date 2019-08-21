package strimmer_new

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

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

/*
func strimTests() []strimTest {
	return []strimTest{
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 6, `1`, token.TT_NUMBER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 5, 6, `1`, token.TT_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 31, `// 'There's a snake in my boot'`, token.TT_COMMENT),
			},
			ExpectToks: []token.Token{},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 6, `2`, token.TT_NUMBER),
				dummyToken(0, 6, 7, ` `, token.TT_SPACE),
				dummyToken(0, 7, 38, `// 'There's a snake in my boot'`, token.TT_COMMENT),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 5, 6, `2`, token.TT_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 5, 20, `"Howdy partner"`, token.TT_STRING),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 5, 20, `Howdy partner`, token.TT_STRING),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 7, `123_456`, token.TT_NUMBER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 7, `123456`, token.TT_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 21, `1__2__3__.__4__5__6__`, token.TT_NUMBER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 21, `123.456`, token.TT_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 6, `func`, token.TT_FUNC),
				dummyToken(0, 6, 7, ` `, token.TT_SPACE),
				dummyToken(0, 7, 10, `END`, token.TT_DONE),
				dummyToken(0, 10, 11, ` `, token.TT_SPACE),
				dummyToken(0, 11, 19, `@PrInTlN`, token.TT_SPELL),
				dummyToken(0, 19, 20, `語`, token.TT_ID),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 6, `func`, token.TT_FUNC),
				dummyToken(0, 7, 10, `end`, token.TT_DONE),
				dummyToken(0, 11, 19, `@println`, token.TT_SPELL),
				dummyToken(0, 19, 20, `語`, token.TT_ID),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				// x <- [
				//   1,
				//   2, 3,
				// ]
				//
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 6, `[`, token.TT_SQUARE_OPEN),
				dummyToken(0, 6, 7, "\n", token.TT_NEWLINE),
				dummyToken(0, 0, 2, `  `, token.TT_SPACE),
				dummyToken(0, 2, 3, `1`, token.TT_NUMBER),
				dummyToken(0, 3, 4, "\n", token.TT_NEWLINE),
				dummyToken(0, 0, 2, `  `, token.TT_SPACE),
				dummyToken(0, 2, 3, `2`, token.TT_NUMBER),
				dummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
				dummyToken(0, 4, 5, ` `, token.TT_SPACE),
				dummyToken(0, 5, 6, `3`, token.TT_NUMBER),
				dummyToken(0, 6, 7, "\n", token.TT_NEWLINE),
				dummyToken(0, 0, 1, `]`, token.TT_SQUARE_CLOSE),
				dummyToken(0, 1, 2, "\n", token.TT_NEWLINE),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
				dummyToken(0, 5, 6, `[`, token.TT_SQUARE_OPEN),
				dummyToken(0, 2, 3, `1`, token.TT_NUMBER),
				dummyToken(0, 3, 4, "\n", token.TT_EOS),
				dummyToken(0, 2, 3, `2`, token.TT_NUMBER),
				dummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM),
				dummyToken(0, 5, 6, `3`, token.TT_NUMBER),
				dummyToken(0, 6, 7, "\n", token.TT_EOS),
				dummyToken(0, 0, 1, `]`, token.TT_SQUARE_CLOSE),
				dummyToken(0, 1, 2, "\n", token.TT_EOS),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 1, 2, ` `, token.TT_SPACE),
				dummyToken(0, 2, 3, `=`, token.TT_ERROR_UPSTREAM),
				dummyToken(0, 3, 4, ` `, token.TT_SPACE),
				dummyToken(0, 4, 5, `2`, token.TT_NUMBER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 1, `x`, token.TT_ID),
				dummyToken(0, 2, 3, `=`, token.TT_ERROR_UPSTREAM),
			},
		},
	}
}
*/
