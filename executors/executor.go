
package executors

import (
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// Executor executes a block of code that may have it's own set of
// variables and context rules.
type Executor interface {

	// Exe executes a statement returning an exit code and an executor
	// that will execute the next statement.
	Exe(scroll *sc.Scroll, line sc.Line) (sh.ExitCode, Executor, sh.ExeError)
}
