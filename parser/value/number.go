package value

import (
	"strconv"
)

// num_value represnts a number Value.
type num_value float64

// Number returns a new number value.
func Number(n float64) Value {
	return num_value(n)
}

// Num satisfies the Value interface.
func (v num_value) Num() (float64, bool) {
	return float64(v), true
}

// String satisfies the Value interface.
func (v num_value) String() string {
	return strconv.FormatFloat(float64(v), byte('g'), -1, 64) + ` (Number)`
}
