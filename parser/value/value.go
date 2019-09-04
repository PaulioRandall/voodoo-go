package value

import (
	"strconv"
)

// Value represents a variables current value.
type Value interface {

	// Number returns the value as a number.
	Num() (float64, bool)

	// String returns the human readable string representation of the value.
	String() string
}

// num_value represnts a number Value.
type num_value float64

// NewNumber returns a new number value.
func NewNumber(n float64) Value {
	return num_value(n)
}

// Num satisfies the Value interface.
func (v num_value) Num() (float64, bool) {
	return float64(v), true
}

// String satisfies the Value interface.
func (v num_value) String() string {
	return `(Number) ` + strconv.FormatFloat(float64(v), byte('g'), -1, 64)
}
