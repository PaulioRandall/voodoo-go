package tree

// Kind represents the type of an expression.
type Kind int

const (
	KD_DONT_CARE Kind = iota - 1
	KD_UNDEFINED
	KD_ASSIGN
	KD_ID
	KD_OPERAND
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
	}

	return "UNDEFINED"
}
