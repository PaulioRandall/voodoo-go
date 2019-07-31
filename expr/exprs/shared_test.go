package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func dummyToken(val string) token.Token {
	return token.Token{
		Val: val,
	}
}

type dummy struct {
	Val ctx.Value
	Err fault.Fault
}

func (d dummy) Evaluate(c *ctx.Context) (v ctx.Value, err fault.Fault) {
	return d.Val, d.Err
}
