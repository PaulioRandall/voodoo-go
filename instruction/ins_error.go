
package instruction

// InsError represents an error while attempting to
// perform an instruction.
type InsError interface {
	error
	
	// Returns the line number of the error.
	Line() int
}

// stdInsError is the standard implementation of InsError.
type stdInsError struct {
	err  string // Error message
	line int    // Line number
}

// NewInsError returns a new initialised InsError.
func NewInsError(err string, line int) InsError {
	e := stdInsError{err, line}
	return InsError(e)
}

// Error satisfies the error interface.
func (l stdInsError) Error() string {
	return l.err
}

// Line satisfies the InsError interface.
func (l stdInsError) Line() int {
	return l.line
}