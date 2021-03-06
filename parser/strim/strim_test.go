package strim

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doStrim(t *testing.T, in, exp []token.Token) {
	act := []token.Token{}
	for _, v := range in {
		if out := Strim(v); out != nil {
			act = append(act, out)
		}
	}
	token.AssertSliceEqual(t, exp, act)
}

func dummy(t string, k token.Kind) token.Token {
	return token.New(t, 0, 0, 0, k)
}

func TestStrim_1(t *testing.T) {
	in := []token.Token{
		// XyZ <- 1
		dummy(`XyZ`, token.TK_ID),
		dummy(` `, token.TK_SPACE),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(` `, token.TK_SPACE),
		dummy(`1`, token.TK_NUMBER),
		dummy("\n", token.TK_NEWLINE),
	}

	exp := []token.Token{
		// xyz<-1
		dummy(`xyz`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`1`, token.TK_NUMBER),
		dummy("\n", token.TK_NEWLINE),
	}

	doStrim(t, in, exp)
}

func TestStrim_2(t *testing.T) {
	in := []token.Token{
		// x <- TrUe
		dummy(`x`, token.TK_ID),
		dummy(` `, token.TK_SPACE),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(` `, token.TK_SPACE),
		dummy(`TrUe`, token.TK_BOOL),
		dummy("\n", token.TK_NEWLINE),
	}

	exp := []token.Token{
		// x<-true
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`true`, token.TK_BOOL),
		dummy("\n", token.TK_NEWLINE),
	}

	doStrim(t, in, exp)
}

func TestStrim_3(t *testing.T) {
	in := []token.Token{
		// x <- `aBc`
		dummy(`x`, token.TK_ID),
		dummy(` `, token.TK_SPACE),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(` `, token.TK_SPACE),
		dummy("`aBc`", token.TK_STRING),
		dummy("\n", token.TK_NEWLINE),
	}

	exp := []token.Token{
		// x<-aBc
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`aBc`, token.TK_STRING),
		dummy("\n", token.TK_NEWLINE),
	}

	doStrim(t, in, exp)
}
