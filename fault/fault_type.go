package fault

// FaultType represents the type of Fault
type FaultType int

const (
	Undefined FaultType = iota
	DevBug
	Parenthesis
	Number
	Function
	String
	Symbol
)

// FaultName returns the name of the supplied fault or
// `undefined` if the type is not known.
func FaultName(t FaultType) string {
	switch t {
	case DevBug:
		return `bug`
	case Parenthesis:
		return `parenthesis`
	case Number:
		return `number`
	case Function:
		return `function`
	case String:
		return `string`
	case Symbol:
		return `symbol`
	}

	return `undefined`
}
