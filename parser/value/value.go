package value

// Value represents a variables current value.
type Value interface {

	// Bool returns the value as a boolean.
	Bool() (bool, bool)

	// Number returns the value as a number.
	Num() (float64, bool)

	// String returns the value as a string.
	Str() (string, bool)

	// String returns the human readable string representation of the value.
	String() string
}
