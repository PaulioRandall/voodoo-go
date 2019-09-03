package operation

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// operation represents an infix expression.
type operation struct {
	t  token.Token
	nu expr.Expr // numerator
	de expr.Expr // denominator
}

// New returns a new assignment expression.
func New(t token.Token, nu, de expr.Expr) expr.Expr {
	return operation{
		t:  t,
		nu: nu,
		de: de,
	}
}

// Token satisfies the Expr interface.
func (o operation) Token() token.Token {
	return o.t
}

// Exe satisfies the Expr interface.
func (o operation) Exe(ctx.Context) (value.Value, perror.Perror) {
	// TODO
	panic(`TODO: operation.Exe(ctx.Context)`)
	return nil, nil
}

// String satisfies the Expr interface.
func (o operation) String() string {
	sb := strings.Builder{}

	sb.WriteString(o.nu.String())
	sb.WriteRune(' ')
	sb.WriteString(o.t.Text())
	sb.WriteRune(' ')
	sb.WriteString(o.de.String())

	return sb.String()
}
