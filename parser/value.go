package parser

// ValueType represents a type of value
type ValueType int

const (
	UNDEFINED ValueType = iota
	BOOL
	NUMBER
	STRING
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
	}

	return `undefined`
}

// Value represents a value within a scroll. Not all values
// are explicitly declared within the code.
type Value interface {

	// Type returns the type of the value.
	ValueType() ValueType
}

// BoolValue represents a boolean value
type BoolValue bool

// ValueType returns the type of the value.
func (v BoolValue) ValueType() ValueType {
	return BOOL
}

// NumberValue represents a number value
type NumberValue float64

// ValueType returns the type of the value.
func (v NumberValue) ValueType() ValueType {
	return NUMBER
}

// StringValue represents a string value
type StringValue string

// ValueType returns the type of the value.
func (v StringValue) ValueType() ValueType {
	return STRING
}
