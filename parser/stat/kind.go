package stat

// Kind represents the type of a statement.
type Kind int

const (
	SK_UNDEFINED  Kind = iota
	SK_EXPRESSION      // 1, 1 + 2, @Print(`abc`)
)

// KindName returns the name of the statement kind.
func KindName(t Kind) string {
	switch t {
	case SK_EXPRESSION:
		return `EXPRESSION`
	default:
		return `UNDEFINED`
	}
}
