package value

// Kind represents the type of the value.
type Kind int

const (
	VK_UNDEFINED Kind = iota
	VK_BOOL
	VK_NUMBER
	VK_STRING
	VK_TUPLE
)

// Name returns the name of the kind.
func (k Kind) Name() string {
	switch k {
	case VK_BOOL:
		return `BOOL`
	case VK_NUMBER:
		return `NUMBER`
	case VK_STRING:
		return `STRING`
	case VK_TUPLE:
		return `TUPLE`
	default:
		return `UNDEFINED`
	}
}
