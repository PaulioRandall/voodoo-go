package value

import (
	"strings"
)

// tuple_value represnts a tuple.
type tuple_value []Value

// Tuple returns a new tuple value.
func Tuple(t ...Value) Value {
	return tuple_value(t)
}

// Bool satisfies the Value interface.
func (v tuple_value) Bool() (bool, bool) {
	return false, false
}

// Num satisfies the Value interface.
func (v tuple_value) Num() (float64, bool) {
	return 0, false
}

// Str satisfies the Value interface.
func (v tuple_value) Str() (string, bool) {
	return ``, false
}

// Tuple satisfies the Value interface.
func (v tuple_value) Tuple() ([]Value, bool) {
	return []Value(v), true
}

// String satisfies the Value interface.
func (v tuple_value) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	sb.WriteRune('\n')

	for _, val := range []Value(v) {
		sb.WriteString(val.String())
		sb.WriteString(",\n")
	}

	sb.WriteRune(']')
	return sb.String() + ` (Tuple)`
}
