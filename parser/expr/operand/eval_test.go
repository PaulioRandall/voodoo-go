package operand

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tok(t string, k token.Kind) token.Token {
	return token.New(t, 0, 0, 0, k)
}

func TestOperand_Eval_1(t *testing.T) {
	o := operand{
		t: tok(`1`, token.TT_NUMBER),
	}

	c := ctx.New(nil)
	exp := value.Number(1)

	act, e := o.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}

func TestOperand_Eval_2(t *testing.T) {
	o := operand{
		t: tok(`3.14159`, token.TT_NUMBER),
	}

	c := ctx.New(nil)
	exp := value.Number(3.14159)

	act, e := o.Eval(c)
	require.Nil(t, e)
	assert.Equal(t, exp, act)
}
