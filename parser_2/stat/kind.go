package stat

// Kind represents the type of a statement.
type Kind int

const (
	SK_UNDEFINED Kind = iota
	SK_ASSIGN         // x: 1
)

// KindName returns the name of the statement kind.
func KindName(t Kind) string {
	switch t {
	case SK_ASSIGN:
		return `ASSIGN`
	default:
		return `UNDEFINED`
	}
}
