package value

// Value represents a variables current value.
type Value interface {

	// Number returns the value as a number.
	Num() (float64, bool)

	// String returns the human readable string representation of the value.
	String() string
}
