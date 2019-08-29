package scanner

// ScanError represents an error found while scanning text.
type ScanError interface {

	// Line returns the line of the error.
	Line() int

	// Index returns the index of the error.
	Index() int

	// Errors returns an array of errors messages.
	Errors() []string
}

// scanErr is an implementation of ScanError.
type scanErr struct {
	l int
	i int
	e []string
}

// Line satisfies the ScanError interface.
func (err scanErr) Line() int {
	return err.l
}

// Index satisfies the ScanError interface.
func (err scanErr) Index() int {
	return err.i
}

// Errors satisfies the ScanError interface.
func (err scanErr) Errors() []string {
	return err.e
}
