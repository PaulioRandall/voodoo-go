package tree

// TreeKind represents the type of an expression.
type TreeKind int

const (
	KD_UNDEFINED TreeKind = iota
	KD_ASSIGNMENT
	KD_ID
)

// TreeKindName returns the name of the input kind.
func TreeKindName(t TreeKind) string {
	switch t {
	case KD_ASSIGNMENT:
		return "ASSIGNMENT"
	case KD_ID:
		return "IDENTIFIER"
	}

	return "UNDEFINED"
}
