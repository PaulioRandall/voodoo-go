package lexer

// LexError represents an error returned by the lexer and
// contains additional error information such as the line
// and column number.
type LexError interface {
	error

	// Line returns the line number.
	Line() int

	// Col returns the column number where the error starts.
	Col() int
}

// stdLexError is the standard implementation of LexError.
type stdLexError struct {
	err  string // Error message
	line int    // Line number
	col  int    // Column number
}

// NewLexError returns a new initialised LexError.
func NewLexError(err string, col int) LexError {
	e := stdLexError{err, 0, col}
	return LexError(e)
}

// Error satisfies the error interface.
func (l stdLexError) Error() string {
	return l.err
}

// Line satisfies the LexError interface.
func (l stdLexError) Line() int {
	return l.line
}

// Col satisfies the LexError interface.
func (l stdLexError) Col() int {
	return l.col
}

// ChangeLine creates a copy of the input LexError with a
// different line number.
func ChangeLine(err LexError, lineNum int) LexError {
	return stdLexError{
		err:  err.Error(),
		line: lineNum,
		col:  err.Col(),
	}
}
