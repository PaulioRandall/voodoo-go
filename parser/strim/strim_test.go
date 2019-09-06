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
		dummy(`XyZ`, token.TT_ID),
		dummy(` `, token.TT_SPACE),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(` `, token.TT_SPACE),
		dummy(`1`, token.TT_NUMBER),
		dummy("\n", token.TT_NEWLINE),
	}

	exp := []token.Token{
		// xyz<-1
		dummy(`xyz`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`1`, token.TT_NUMBER),
		dummy("\n", token.TT_NEWLINE),
	}

	doStrim(t, in, exp)
}

func TestStrim_2(t *testing.T) {
	in := []token.Token{
		// x <- TrUe
		dummy(`x`, token.TT_ID),
		dummy(` `, token.TT_SPACE),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(` `, token.TT_SPACE),
		dummy(`TrUe`, token.TT_BOOL),
		dummy("\n", token.TT_NEWLINE),
	}

	exp := []token.Token{
		// x<-true
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`true`, token.TT_BOOL),
		dummy("\n", token.TT_NEWLINE),
	}

	doStrim(t, in, exp)
}

func TestStrim_3(t *testing.T) {
	in := []token.Token{
		// x <- `aBc`
		dummy(`x`, token.TT_ID),
		dummy(` `, token.TT_SPACE),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(` `, token.TT_SPACE),
		dummy("`aBc`", token.TT_STRING),
		dummy("\n", token.TT_NEWLINE),
	}

	exp := []token.Token{
		// x<-aBc
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`aBc`, token.TT_STRING),
		dummy("\n", token.TT_NEWLINE),
	}

	doStrim(t, in, exp)
}
