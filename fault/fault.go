package fault

import (
	"fmt"
	"runtime"
)

// Fault represents an error produced by this program rather than a library. The
// error could be a compiler bug or problem with code being parsed.
type Fault interface {

	// Print prints the fault to logs.
	Print(file string)
}

// ReaderFault represents an error with reading wrapped as a fault.
type ReaderFault string

// Print satisfies the Fault interface.
func (err ReaderFault) Print(file string) {
	fmt.Println(err)
}

// TODO: Move this to the parser pkg

// SyntaxFault represents a generic fault with syntax.
type SyntaxFault struct {
	Index int      // Index where the error actually occurred
	Msgs  []string // Description of the error
}

// Print satisfies the Fault interface.
func (err SyntaxFault) Print(file string) {
	//sc.PrettyPrintError(-1, err.Index, err.Msgs...)
}

// CurrLine returns the line of the caller to this function.
func CurrLine() int {
	_, _, line, ok := runtime.Caller(1)

	if !ok {
		return -1
	}

	return line
}
