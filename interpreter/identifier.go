
package interpreter

// Operator represents a arithmetic or boolean operator.
type Operator string

const (
	BoolEqual Operator = `==`
	BoolNotEqual Operator = `!=`
	BoolLessThan Operator = `<`
	BoolGreaterThan Operator = `>`
	BoolLessThanOrEqual Operator = `<=`
	BoolGreaterThanOrEqual Operator = `>=`
	BoolAnd Operator = `&&`
	BoolOr Operator = `||`
	
	MathAdd Operator = `+`
	MathSub Operator = `-`
	MathMul Operator = `*`
	MathDiv Operator = `/`
)
