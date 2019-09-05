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

func TestAssign_Eval_1(t *testing.T) {
	a := assign{
		t: tok(`<-`, token.TT_ASSIGN),
		dst: []token.Token{
			tok(`x`, token.TT_ID),
		},
		src: []expr.Expr{
			dummy(`1`, token.TT_NUMBER, value.Number(1)),
		},
	}

	c := ctx.New(nil)

	expCtx := ctx.New(nil)
	expCtx.Vars[`x`] = value.Number(1)

	doTestAssign_Eval(t, a, c, expCtx)
}
