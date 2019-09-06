package value

// str_value represnts a string literal Value.
type str_value string

// String returns a new string value.
func String(s string) Value {
	return str_value(s)
}

// Bool satisfies the Value interface.
func (v str_value) Bool() (bool, bool) {
	return false, false
}

// Num satisfies the Value interface.
func (v str_value) Num() (float64, bool) {
	return 0, false
}

// Str satisfies the Value interface.
func (v str_value) Str() (string, bool) {
	return string(v), true
}

// Tuple satisfies the Value interface.
func (v str_value) Tuple() ([]Value, bool) {
	return nil, false
}

// String satisfies the Value interface.
func (v str_value) String() string {
	return string(v) + ` (String)`
}
