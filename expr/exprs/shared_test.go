package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func typedToken(v string, t token.TokenType) token.Token {
	return token.Token{
		Val:  v,
		Type: t,
	}
}

func dummyToken(v string) token.Token {
	return token.Token{
		Val: v,
	}
}

type dummy struct {
	Val ctx.Value
	Err fault.Fault
}

func (d dummy) Evaluate(c *ctx.Context) (v ctx.Value, err fault.Fault) {
	return d.Val, d.Err
}

func valDummy(v ctx.Value) ctx.Expression {
	return dummy{
		Val: v,
	}
}

func errDummy(err fault.Fault) ctx.Expression {
	return dummy{
		Err: err,
	}
}
