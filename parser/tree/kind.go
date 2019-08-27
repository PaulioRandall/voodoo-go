package tree

// Kind represents the type of an expression.
type Kind int

const (
	KD_UNDEFINED Kind = iota
	KD_ASSIGN
	KD_ID
	KD_OPERAND
	KD_OPERATION
	KD_UNION // Joins the left and the right
)

// KindName returns the name of the input kind.
func KindName(t Kind) string {
	switch t {
	case KD_ASSIGN:
		return "ASSIGNMENT"
	case KD_ID:
		return "IDENTIFIER"
	case KD_OPERAND:
		return "OPERAND"
	case KD_OPERATION:
		return "OPERATION"
	}

	return "UNDEFINED"
}
