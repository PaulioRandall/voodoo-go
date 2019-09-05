package expr

import (
	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

type Dummy struct {
	T token.Token
	F func(ctx.Context) (value.Value, perror.Perror)
}

func (d Dummy) Token() token.Token {
	return d.T
}

func (d Dummy) Eval(c ctx.Context) (value.Value, perror.Perror) {
	return d.F(c)
}

func (d Dummy) String() string {
	return ``
}
