package assign

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tok(t string, k token.Kind) token.Token {
	return token.New(t, 0, 0, 0, k)
}

func dummy(t string, k token.Kind, v value.Value) expr.Expr {
	return expr.Dummy{
		T: tok(t, k),
		F: func(c ctx.Context) (value.Value, perror.Perror) {
			return v, nil
		},
	}
}

func doTestAssign_Eval(t *testing.T, a assign, c, expCtx ctx.Context) {
	act, e := a.Eval(c)
	require.Nil(t, e)
	assert.Nil(t, act)
	assert.Equal(t, expCtx, c)
}

func doErrTestAssign_Eval(t *testing.T, a assign, c ctx.Context) {
	_, e := a.Eval(c)
	require.NotNil(t, e)
}

func TestAssign_Eval_1(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TK_ASSIGN),
		dst: []token.Token{
			tok(`x`, token.TK_ID),
		},
		src: []expr.Expr{
			dummy(`1`, token.TK_NUMBER, value.Number(1)),
		},
	}

	c := ctx.New(nil)

	exp := ctx.New(nil)
	exp.Vars[`x`] = value.Number(1)

	doTestAssign_Eval(t, a, c, exp)
}

func TestAssign_Eval_2(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TK_ASSIGN),
		dst: []token.Token{
			tok(`x`, token.TK_ID),
		},
		src: []expr.Expr{
			dummy(`_`, token.TK_VOID, nil),
		},
	}

	c := ctx.New(nil)
	exp := ctx.New(nil)

	doTestAssign_Eval(t, a, c, exp)
}

func TestAssign_Eval_3(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TK_ASSIGN),
		dst: []token.Token{
			tok(`_`, token.TK_VOID),
		},
		src: []expr.Expr{
			dummy(`3`, token.TK_NUMBER, value.Number(3)),
		},
	}

	c := ctx.New(nil)
	exp := ctx.New(nil)

	doTestAssign_Eval(t, a, c, exp)
}

func TestAssign_Eval_4(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TK_ASSIGN),
		dst: []token.Token{
			tok(`x`, token.TK_ID),
			tok(`y`, token.TK_ID),
			tok(`z`, token.TK_ID),
		},
		src: []expr.Expr{
			dummy(`4`, token.TK_NUMBER, value.Number(4)),
			dummy(`Dragonfly`, token.TK_STRING, value.String(`Dragonfly`)),
			dummy(`_`, token.TK_VOID, nil),
		},
	}

	c := ctx.New(nil)
	exp := ctx.New(nil)
	exp.Vars[`x`] = value.Number(4)
	exp.Vars[`y`] = value.String(`Dragonfly`)

	doTestAssign_Eval(t, a, c, exp)
}

func TestAssign_Eval_5(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TK_ASSIGN),
		dst: []token.Token{
			tok(`x`, token.TK_ID),
			tok(`y`, token.TK_ID),
		},
		src: []expr.Expr{
			dummy(`4`, token.TK_NUMBER, value.Number(4)),
			dummy(`Dragonfly`, token.TK_STRING, value.String(`Dragonfly`)),
			dummy(`_`, token.TK_VOID, nil),
		},
	}

	c := ctx.New(nil)

	doErrTestAssign_Eval(t, a, c)
}

func TestAssign_Eval_6(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TK_ASSIGN),
		dst: []token.Token{
			tok(`x`, token.TK_ID),
			tok(`y`, token.TK_ID),
			tok(`z`, token.TK_ID),
		},
		src: []expr.Expr{
			dummy(`4`, token.TK_NUMBER, value.Number(4)),
			dummy(`Dragonfly`, token.TK_STRING, value.String(`Dragonfly`)),
		},
	}

	c := ctx.New(nil)

	doErrTestAssign_Eval(t, a, c)
}

func TestAssign_Eval_7(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TK_ASSIGN),
		dst: []token.Token{
			tok(`x`, token.TK_ID),
		},
		src: []expr.Expr{
			dummy(`Dragonfly`, token.TK_STRING, value.String(`Dragonfly`)),
		},
	}

	c := ctx.New(nil)
	c.Vars[`x`] = value.Number(4)

	doErrTestAssign_Eval(t, a, c)
}
