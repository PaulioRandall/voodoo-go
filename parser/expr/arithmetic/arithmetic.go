package arithmetic

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// arithmetic represents an infix arithmetic expression.
type arithmetic struct {
	t  token.Token
	nu expr.Expr // numerator
	de expr.Expr // denominator
}

// New returns a new arithmetic expression.
func New(t token.Token, nu, de expr.Expr) expr.Expr {
	return arithmetic{
		t:  t,
		nu: nu,
		de: de,
	}
}

// Token satisfies the Expr interface.
func (a arithmetic) Token() token.Token {
	return a.t
}

// String satisfies the Expr interface.
func (a arithmetic) String() string {
	sb := strings.Builder{}

	sb.WriteString(a.nu.String())
	sb.WriteRune(' ')
	sb.WriteString(a.t.Text())
	sb.WriteRune(' ')
	sb.WriteString(a.de.String())

	return sb.String()
}
