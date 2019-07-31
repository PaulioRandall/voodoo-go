package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Arithmetic represents an expression which performs
// one of the four basic arithmetic operations. Addition,
// subtraction, multiplication, or division and returns
// the result.
type Arithmetic struct {
	Operator token.Token    // Arithmetic token defined in the scroll
	Left     ctx.Expression // Left value or numerator
	Right    ctx.Expression // Right value or denominator
}

// Evaluate satisfies the Expression interface.
func (a Arithmetic) Evaluate(c *ctx.Context) (val ctx.Value, err fault.Fault) {
	left, right, err := a.evalChildren(c)
	if err != nil {
		return
	}

	switch a.Operator.Type {
	case token.CALC_ADD:
		val = add(left, right)
	case token.CALC_SUBTRACT:
		val = subtract(left, right)
	case token.CALC_MULTIPLY:
		val = multiply(left, right)
	case token.CALC_DIVIDE:
		val = divide(left, right)
	default:
		err = ctx.EvalFault{
			ExprType: `Arithmetic`,
			Msgs: []string{
				`Unknown operator`,
				`NOTE: Should this be a bug?`,
			},
		}
	}

	return
}

// add adds the left value to the right returning the result.
func add(left, right ctx.Value) ctx.Value {
	return left.(ctx.NumberValue) + right.(ctx.NumberValue)
}

// subtract subtracts the right value from the left returning the
// result.
func subtract(left, right ctx.Value) ctx.Value {
	return left.(ctx.NumberValue) - right.(ctx.NumberValue)
}

// multiply multiplies the left value with the right returning the
// result.
func multiply(left, right ctx.Value) ctx.Value {
	return left.(ctx.NumberValue) * right.(ctx.NumberValue)
}

// divide divides the right value from the right returning the
// result.
func divide(left, right ctx.Value) ctx.Value {
	return left.(ctx.NumberValue) / right.(ctx.NumberValue)
}

// evalChildren evaluates the left and right child expressions.
func (a Arithmetic) evalChildren(c *ctx.Context) (left ctx.Value, right ctx.Value, err fault.Fault) {
	left, err = a.Left.Evaluate(c)
	if err != nil {
		return
	}

	err = checkIsNumber(left, true)
	if err != nil {
		return
	}

	right, err = a.Right.Evaluate(c)
	if err != nil {
		return
	}

	err = checkIsNumber(right, false)
	return
}

// checkIsNumber returns a fault if the input value is not
// a number.
func checkIsNumber(v ctx.Value, isLeft bool) fault.Fault {
	if v.Type() != ctx.NUMBER {

		var m string
		if isLeft {
			m = `The left expression or operator did not equate to a number`
		} else {
			m = `The right expression or operator did not equate to a number`
		}

		return ctx.EvalFault{
			ExprType: `Arithmetic`,
			Msgs: []string{
				m,
			},
		}
	}

	return nil
}
