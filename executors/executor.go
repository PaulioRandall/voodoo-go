
package executors

import (
	sc "github.com/PaulioRandall/voodoo-go/scroll"
)

// Executor executes a block of code that may have it's own set of
// variables and context rules. Executables may be an 'activity', consisting
// of other executables, or an 'operation', 1-2 variables with 1-2 operators.
//
// '1 + 1' is an operation
// 'a = 1 + 1' is an activity containing two operations
type Executor interface {

	// Exe continues execution of the scroll lines until an error or
	// the end of the exectable block is encountered. 
	Exe(scroll *sc.Scroll) (exitCode int, err error)
	
	// Vars returns the state of the executor.
	Vars() map[string]sc.VooValue
}

// Activity represents a group of operations such as 'a = 1 + 1'.
// Multi-line blocks of code are always activities.
type Activity interface {
	Executor
}
