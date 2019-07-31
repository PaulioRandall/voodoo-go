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

// Type satisfies the Value interface.
func (v BoolValue) Type() ValueType {
	return BOOL
}

// Type satisfies the Value interface.
func (v NumberValue) Type() ValueType {
	return NUMBER
}

// Type satisfies the Value interface.
func (v StringValue) Type() ValueType {
	return STRING
}

// Type satisfies the Value interface.
func (v ListValue) Type() ValueType {
	return LIST
}

// HasSameType returns true if both input values have the
// same type.
func HasSameType(a Value, b Value) bool {
	return a.Type() == b.Type()
}
