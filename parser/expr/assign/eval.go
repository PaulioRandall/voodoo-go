package assign

import (
	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Eval satisfies the Expr interface.
func (a assign) Eval(c ctx.Context) (value.Value, perror.Perror) {
	v, e := a.src[0].Eval(c)
	if e != nil {
		return nil, e
	}

	switch t := a.dst[0].Text(); {
	case v == nil:
		delete(c.Vars, t)
	case a.dst[0].Kind() != token.TT_VOID:
		c.Vars[t] = v
	}

	return nil, nil
}
