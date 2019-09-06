package value

import (
	"strconv"
)

// bool_value represnts a boolean Value.
type bool_value bool

// Bool returns a new boolean value.
func Bool(b bool) Value {
	return bool_value(b)
}

// Bool satisfies the Value interface.
func (v bool_value) Bool() (bool, bool) {
	return bool(v), true
}

// Num satisfies the Value interface.
func (v bool_value) Num() (float64, bool) {
	return 0, false
}

// String satisfies the Value interface.
func (v bool_value) String() string {
	return strconv.FormatBool(bool(v)) + ` (Bool)`
}
