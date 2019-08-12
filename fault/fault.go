package fault

import (
	"runtime"
)

// Fault represents an error produced by this program rather than a library. The
// error could be a compiler bug or problem with code being parsed.
type Fault interface {

	// Print prints the fault to logs.
	Print(file string)
}

// CurrLine returns the line of the caller to this function.
func CurrLine() int {
	_, _, line, ok := runtime.Caller(1)

	if !ok {
		return -1
	}

	return line
}
