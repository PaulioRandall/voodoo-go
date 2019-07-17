package lexer

// LexError represents an error returned by the lexer and
// contains additional error information such as the line
// and column number.
type LexError interface {
	error

	// Line sets the line number.
	Line(lineNum int) LexError
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
func (l stdLexError) Line(lineNum int) LexError {
	l.line = lineNum
	return LexError(l)
}
