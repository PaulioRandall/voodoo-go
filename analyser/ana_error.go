package analyser

// AnaError represents an error returned by the analyser and
// contains additional error information such as the line
// and column number.
type AnaError interface {
	error

	// Line returns the line number.
	Line() int

	// Col returns the column number where the error starts.
	Col() int
}

// stdAnaError is the standard implementation of AnaError.
type stdAnaError struct {
	err  string // Error message
	line int    // Line number
	col  int    // Column number
}

// NewAnaError returns a new initialised AnaError.
func NewAnaError(err string, col int) AnaError {
	e := stdAnaError{err, 0, col}
	return AnaError(e)
}

// Error satisfies the error interface.
func (l stdAnaError) Error() string {
	return l.err
}

// Line satisfies the AnaError interface.
func (l stdAnaError) Line() int {
	return l.line
}

// Col satisfies the AnaError interface.
func (l stdAnaError) Col() int {
	return l.col
}

// ChangeLine creates a copy of the input AnaError with a
// different line number.
func ChangeLine(err AnaError, lineNum int) AnaError {
	return stdAnaError{
		err:  err.Error(),
		line: lineNum,
		col:  err.Col(),
	}
}
