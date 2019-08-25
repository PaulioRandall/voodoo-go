package expr

// ExprKind represents the type of an expression.
type ExprKind int

const (
	UNDEFINED ExprKind = iota
	ASSIGNMENT
	LITERAL
	IDENTIFIER
)

// ExprKindName returns the name of the input kind.
func ExprKindName(t ExprKind) string {
	switch t {
	case ASSIGNMENT:
		return "ASSIGNMENT"
	case LITERAL:
		return "LITERAL"
	case IDENTIFIER:
		return "IDENTIFIER"
	}

	return "UNDEFINED"
}
