package tree

// TreeKind represents the type of an expression.
type TreeKind int

const (
	UNDEFINED TreeKind = iota
	ASSIGNMENT
	LITERAL
	IDENTIFIER
)

// TreeKindName returns the name of the input kind.
func TreeKindName(t TreeKind) string {
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
