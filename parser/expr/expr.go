package expr

// Expr represents an expression.
type Expr interface {

	// Kind returns the type of the expression
	Kind() int
}
