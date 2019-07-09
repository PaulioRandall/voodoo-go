
package executors

import (
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	sh "github.com/PaulioRandall/voodoo-go/shared"
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

// ExeLine satisfies the Executor interface.
func (op *Operation) ExeLine(scroll *sc.Scroll, line string) (sh.ExitCode, Executor, sh.ExeError) {
	return sh.CatchAllErr, nil, sh.NewError(sh.CatchAllErr, "Not yet implemented")
}
