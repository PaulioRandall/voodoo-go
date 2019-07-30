package fault_new

import (
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
