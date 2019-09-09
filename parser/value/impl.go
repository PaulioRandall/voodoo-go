package value

import (
	"strconv"
	"strings"
)

// value implements the Value interface.
type value struct {
	k Kind
	v interface{}
}

// Kind satisfies the Value interface.
func (v value) Kind() Kind {
	return v.k
}

// SameKind satisfies the Value interface.
func (v value) SameKind(other Value) bool {
	return v.k == other.Kind()
}

// Bool satisfies the Value interface.
func (v value) Bool() (bool, bool) {
	if v.k != VK_BOOL {
		return false, false
	}
	return v.v.(bool), true
}

// Num satisfies the Value interface.
func (v value) Num() (float64, bool) {
	if v.k != VK_NUMBER {
		return 0, false
	}
	return v.v.(float64), true
}

// Str satisfies the Value interface.
func (v value) Str() (string, bool) {
	if v.k != VK_STRING {
		return ``, false
	}
	return v.v.(string), true
}

// Tuple satisfies the Value interface.
func (v value) Tuple() ([]Value, bool) {
	if v.k != VK_TUPLE {
		return nil, false
	}
	return v.v.([]Value), true
}

// String satisfies the Value interface.
func (v value) String() string {
	switch v.k {
	case VK_BOOL:
		b, _ := v.v.(bool)
		return `( Bool ) -> ` + strconv.FormatBool(b)
	case VK_NUMBER:
		n, _ := v.v.(float64)
		return `(Number) -> ` + strconv.FormatFloat(n, byte('g'), -1, 64)
	case VK_STRING:
		s, _ := v.v.(string)
		return `(String) -> ` + "`" + s + "`"
	case VK_TUPLE:
		t, _ := v.v.([]Value)
		return tupleString(t)
	default:
		return `(UNDEFINED)`
	}
}

// tupleString returns the string representation of a tuple.
func tupleString(v []Value) string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	sb.WriteRune('\n')

	for _, val := range []Value(v) {
		sb.WriteString(val.String())
		sb.WriteString(",\n")
	}

	sb.WriteRune(']')
	return ` (Tuple ) -> ` + sb.String()
}
