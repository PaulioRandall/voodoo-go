
package executors

import (
	sc "github.com/PaulioRandall/voodoo-go/scroll"
)

// OperatorType represents the type of an operation.
type OperatorType string

const (
	BoolOpType OperatorType = "bool"
	MathOpType OperatorType = "math"
	AssignOpType OperatorType = "assign"
	TruthyOpType OperatorType = "truthy"
)

// Operation represents a simple operation such as '1 + 1'.
type Operation struct {
	OpType OperatorType
	Op sc.Operator
	Snip sc.Snippet
}

// Exe satisfies the Executor interface.
func (op *Operation) Exe(scroll *sc.Scroll) (exitCode int, err error) {
	// TODO: Move stuff from snippet 
	return exitCode, err
}

// Vars satisfies the Executor interface.
func (op *Operation) Vars() map[string]sc.VooValue {
	return nil
}
