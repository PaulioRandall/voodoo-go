package operation

// OpError represents an error while attempting to
// perform an instruction.
type OpError interface {
	error

	// Returns the line number of the error.
	Line() int
}

// stdOpError is the standard implementation of OpError.
type stdOpError struct {
	err  string // Error message
	line int    // Line number
}

// NewOpError returns a new initialised OpError.
func NewOpError(err string, line int) OpError {
	e := stdOpError{err, line}
	return OpError(e)
}

// Error satisfies the error interface.
func (l stdOpError) Error() string {
	return l.err
}

// Line satisfies the OpError interface.
func (l stdOpError) Line() int {
	return l.line
}
