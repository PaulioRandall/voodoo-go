package value

// Value represents a variables current value.
type Value interface {

	// Kind returns the type of the value.
	Kind() Kind

	// SameKind returns true if the kind of the input value is the same as the
	// receivers kind.
	SameKind(Value) bool

	// Bool returns the value as a boolean.
	Bool() (bool, bool)

	// Number returns the value as a number.
	Num() (float64, bool)

	// String returns the value as a string.
	Str() (string, bool)

	// Tuple returns the value as a list.
	Tuple() ([]Value, bool)

	// String returns the human readable string representation of the value.
	String() string
}

// Bool returns a new bool value.
func Bool(v bool) Value {
	return value{
		k: VK_BOOL,
		v: v,
	}
}

// Number returns a new number value.
func Number(v float64) Value {
	return value{
		k: VK_NUMBER,
		v: v,
	}
}

// String returns a new string value.
func String(v string) Value {
	return value{
		k: VK_STRING,
		v: v,
	}
}

// Tuple returns a new tuple value.
func Tuple(v ...Value) Value {
	return value{
		k: VK_TUPLE,
		v: v,
	}
}
