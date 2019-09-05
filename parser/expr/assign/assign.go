package assign

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// assign represents an assignment expression.
type assign struct {
	t   token.Token
	src []expr.Expr
	dst []token.Token
}

// New returns a new assignment expression.
func New(t token.Token, src []expr.Expr, dst []token.Token) expr.Expr {
	return assign{
		t:   t,
		src: src,
		dst: dst,
	}
}

// Token satisfies the Expr interface.
func (a assign) Token() token.Token {
	return a.t
}

// String satisfies the Expr interface.
func (a assign) String() string {
	sb := strings.Builder{}

	a.writeTokens(&sb)
	sb.WriteString(a.t.Text())
	a.writeExprs(&sb)

	return sb.String()
}

// writeTokens writes the text of each supplied token into the string builder
// as a comma separated list.
func (a assign) writeTokens(sb *strings.Builder) {
	for i, tk := range a.dst {
		if i != 0 {
			sb.WriteString(`, `)
		}

		sb.WriteString(tk.Text())
	}
}

// writeExprs writes expression placeholder instead of printing the expression
// as it could be quite lengthy.
func (a assign) writeExprs(sb *strings.Builder) {
	for i, ex := range a.src {
		if i != 0 {
			sb.WriteString(`, `)
		}

		sb.WriteString(ex.String())
	}
}
