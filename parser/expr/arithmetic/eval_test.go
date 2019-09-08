package arithmetic

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

func doTestArithmetic_Eval(t *testing.T, a arithmetic, c ctx.Context, exp value.Value) {
	act, e := a.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}

func doErrTestArithmetic_Eval(t *testing.T, a arithmetic, c ctx.Context) {
	_, e := a.Eval(c)
	require.NotNil(t, e)
}

func TestArithmetic_Eval_1(t *testing.T) {
	a := arithmetic{
		t:  tok(`+`, token.TK_ADD),
		nu: dummy(`1`, token.TK_NUMBER, value.Number(1)),
		de: dummy(`2`, token.TK_NUMBER, value.Number(2)),
	}

	exp := value.Number(3)

	c := ctx.New(nil)
	doTestArithmetic_Eval(t, a, c, exp)
}

func TestArithmetic_Eval_2(t *testing.T) {
	a := arithmetic{
		t:  tok(`-`, token.TK_SUBTRACT),
		nu: dummy(`8`, token.TK_NUMBER, value.Number(8)),
		de: dummy(`3`, token.TK_NUMBER, value.Number(3)),
	}

	exp := value.Number(5)

	c := ctx.New(nil)
	doTestArithmetic_Eval(t, a, c, exp)
}

func TestArithmetic_Eval_3(t *testing.T) {
	a := arithmetic{
		t:  tok(`*`, token.TK_MULTIPLY),
		nu: dummy(`6`, token.TK_NUMBER, value.Number(6)),
		de: dummy(`7`, token.TK_NUMBER, value.Number(7)),
	}

	exp := value.Number(42)

	c := ctx.New(nil)
	doTestArithmetic_Eval(t, a, c, exp)
}

func TestArithmetic_Eval_4(t *testing.T) {
	a := arithmetic{
		t:  tok(`/`, token.TK_DIVIDE),
		nu: dummy(`64`, token.TK_NUMBER, value.Number(64)),
		de: dummy(`4`, token.TK_NUMBER, value.Number(4)),
	}

	exp := value.Number(16)

	c := ctx.New(nil)
	doTestArithmetic_Eval(t, a, c, exp)
}

func TestArithmetic_Eval_5(t *testing.T) {
	a := arithmetic{
		t:  tok(`/`, token.TK_DIVIDE),
		nu: dummy(`5`, token.TK_NUMBER, value.Number(5)),
		de: dummy(`0`, token.TK_NUMBER, value.Number(0)),
	}

	c := ctx.New(nil)
	doErrTestArithmetic_Eval(t, a, c)
}

func TestArithmetic_Eval_6(t *testing.T) {
	a := arithmetic{
		t:  tok(`%`, token.TK_MODULO),
		nu: dummy(`12`, token.TK_NUMBER, value.Number(12)),
		de: dummy(`5`, token.TK_NUMBER, value.Number(5)),
	}

	exp := value.Number(2)

	c := ctx.New(nil)
	doTestArithmetic_Eval(t, a, c, exp)
}
