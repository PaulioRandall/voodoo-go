package ctx

// ValueType represents a type of value
type ValueType int

const (
	UNDEFINED ValueType = iota
	BOOL
	NUMBER
	STRING
	LIST
)

// NameOfValueType returns the name of a value type.
func NameOfValueType(t ValueType) string {
	switch t {
	case BOOL:
		return `bool`
	case NUMBER:
		return `number`
	case STRING:
		return `string`
	case LIST:
		return `list`
	}

	return `undefined`
}

// Value represents a value within a scroll. Not all values
// are explicitly declared within the code.
type Value interface {

	// Type returns the type of the value.
	Type() ValueType
}

// BoolValue represents a boolean.
type BoolValue bool

// NumberValue represents a number.
type NumberValue float64

// StringValue represents a string.
type StringValue string

// ListValue represents a list.
type ListValue []Value

// Type returns the type of the value.
func (v BoolValue) Type() ValueType {
	return BOOL
}

// Type returns the type of the value.
func (v NumberValue) Type() ValueType {
	return NUMBER
}

// Type returns the type of the value.
func (v StringValue) Type() ValueType {
	return STRING
}

// Type returns the type of the value.
func (v ListValue) Type() ValueType {
	return LIST
}
