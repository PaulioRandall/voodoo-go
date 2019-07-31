package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Assignment represents an assignment expression.
type Assignment struct {
	Operator   token.Token    // Scroll token representing the operator used
	Identifier token.Token    // Identifier that will reference the result
	Expression ctx.Expression // Expression to evaluate that produces the result
}

// Evaluate satisfies the Expression interface.
func (a Assignment) Evaluate(c *ctx.Context) (v ctx.Value, err fault.Fault) {
	val, err := a.Expression.Evaluate(c)

	if err != nil {
		return
	}

	if val == nil {
		err = ctx.EvalFault{
			ExprType: `Assignment`,
			Msgs: []string{
				`The expression to the right did not evaluate to a value`,
				`Can't assign nothing to an identifier`,
				`NOTE: Should this be a bug?`,
				`NOTE: Variable deletion as a feature?`,
			},
		}
		return
	}

	id := a.Identifier.Val
	c.Assign(id, val)
	return
}
