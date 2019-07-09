
package executors

import (
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// Executor executes a block of code that may have it's own set of
// variables and context rules.
type Executor interface {
	
	// ExeLine executes a line of code returning an exit code and the executor
	// that should execute the next line.
	ExeLine(scroll *sc.Scroll, line string) (sh.ExitCode, Executor, sh.ExeError)
}
