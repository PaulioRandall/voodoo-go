package fault

import (
	"runtime"

	"github.com/PaulioRandall/voodoo-go/scroll"
)

// Fault represents an error produced by this program
// rather than a library. The error could be due to a bug
// or could detail a problem with code being parsed.
type Fault interface {

	// Print prints the fault to console with the line number
	// of the scroll where the error originated.
	Print(sc *scroll.Scroll, line int)
}

// SyntaxFault represents a generic fault with syntax.
// If different forms of error logging are required then
// they must implement the Fault interface.
type SyntaxFault struct {
	Index int      // Index where the error actually occurred
	Msgs  []string // Description of the error
}

// Print satisfies the Fault interface.
func (err SyntaxFault) Print(sc *scroll.Scroll, line int) {
	sc.PrettyPrintError(line, err.Index, err.Msgs...)
}

// CurrLine returns the line of the caller to this function.
func CurrLine() int {
	_, _, line, ok := runtime.Caller(1)

	if !ok {
		return -1
	}

	return line
}
