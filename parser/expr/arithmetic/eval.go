package arithmetic

import (
	"math"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// operation represents an arithmetic operation on two values.
type operation func(nu, de float64, t token.Token) (value.Value, perror.Perror)

// Eval satisfies the Expr interface.
func (a arithmetic) Eval(c ctx.Context) (value.Value, perror.Perror) {
	switch a.t.Kind() {
	case token.TK_ADD:
		return a.doOperation(c, add)
	case token.TK_SUBTRACT:
		return a.doOperation(c, subtract)
	case token.TK_MULTIPLY:
		return a.doOperation(c, multiply)
	case token.TK_DIVIDE:
		return a.doOperation(c, divide)
	case token.TK_MODULO:
		return a.doOperation(c, modulo)
	default:
		return nil, a.invalidKind()
	}
}

// add adds together two numbers.
func add(nu, de float64, t token.Token) (value.Value, perror.Perror) {
	return value.Number(nu + de), nil
}

// subtract subtracts one number from another.
func subtract(nu, de float64, t token.Token) (value.Value, perror.Perror) {
	return value.Number(nu - de), nil
}

// multiply two numbers together.
func multiply(nu, de float64, t token.Token) (value.Value, perror.Perror) {
	return value.Number(nu * de), nil
}

// divide divides one number by another.
func divide(nu, de float64, t token.Token) (value.Value, perror.Perror) {
	if de == 0 {
		return nil, divideByZero(t)
	}
	return value.Number(nu / de), nil
}

// modulo finds the remainder of dividing one number by another.
func modulo(nu, de float64, t token.Token) (value.Value, perror.Perror) {
	return value.Number(math.Mod(nu, de)), nil
}

// doOperation evaluates the arithmetic operation.
func (a arithmetic) doOperation(c ctx.Context, f operation) (value.Value, perror.Perror) {
	nu, e := evalToFloat64(c, a.nu)
	if e != nil {
		return nil, e
	}

	de, e := evalToFloat64(c, a.de)
	if e != nil {
		return nil, e
	}

	return f(nu, de, a.de.Token())
}

// evalToFloat64 evaluates the operand to a float.
func evalToFloat64(c ctx.Context, ex expr.Expr) (float64, perror.Perror) {
	v, e := ex.Eval(c)
	if e != nil {
		return 0, e
	}

	n, ok := v.Num()
	if !ok {
		return 0, notNumber(ex.Token())
	}

	return n, nil
}

// notInteger returns a new Perror for a when an attempt is being made to
// find the remander of a numerator which is not an integer.
func notInteger(t token.Token) perror.Perror {
	return perror.New(
		t.Line(),
		t.Start(),
		[]string{
			"The numerator is not an integer",
			"Can't find the remainder of a non integer",
		},
	)
}

// divideByZero returns a new Perror for a when an attempt is being made to
// divide by zero.
func divideByZero(t token.Token) perror.Perror {
	return perror.New(
		t.Line(),
		t.Start(),
		[]string{
			"The denominator is zero",
			"Can't divide by zero",
		},
	)
}

// notNumber returns a new Perror for a when a value is not a number when a
// number is required.
func notNumber(t token.Token) perror.Perror {
	return perror.New(
		t.Line(),
		t.Start(),
		[]string{
			"Expected value to be a number",
		},
	)
}

// invalidKind returns a new Perror for a when the token kind can not be
// handled by the evaluator.
func (a arithmetic) invalidKind() perror.Perror {
	return a.newPerror([]string{
		"Can't handle aritmetic operators of kind '" + token.KindName(a.t.Kind()) + "'",
	})
}

// newPerror creates a new Perror.
func (a arithmetic) newPerror(m []string) perror.Perror {
	return perror.New(
		a.t.Line(),
		a.t.Start(),
		m,
	)
}
