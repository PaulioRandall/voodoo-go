package lexer

// LexError represents an error returned by the lexer and
// contains additional error information such as the line
// and column number.
type LexError interface {
	error
}

// stdLexError is the standard implementation of LexError.
type stdLexError struct {
	Err  string // Error message
	Line int    // Line number
	Col  int    // Column number
}

// NewLexError returns a new initialised LexError.
func NewLexError(err string, line, col int) LexError {
	return LexError(stdLexError{err, line, col})
}

// Error satisfies the error interface.
func (l stdLexError) Error() string {
	return l.Err
}
