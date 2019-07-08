
package interpreter

// ValueType represents the type of a voodoo value.
type ValueType int

// Declaration of Value types.
const (
	BoolType ValueType = iota + 1
	NumType
	StrType
	ListType
	ObjType
	FuncType
)

// KeyValuePair represents with a key value pair.
type KeyValuePair struct {
	Key VoodooValue
	Value VoodooValue
}

// VoodooValue represents a value within the scroll.
type VoodooValue struct {
	ValueType ValueType
	BoolValue bool
	NumValue float64
	StrValue string
	ListValue []VoodooValue
	ObjValue []KeyValuePair
	FuncValue []string
}
