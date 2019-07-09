
package executors

import (
	"strings"
	
	sc "github.com/PaulioRandall/voodoo-go/scroll"
)

// ScrollActivity represents the scroll itself as an activity.
type ScrollActivity struct {
	vars map[string]sc.VooValue
}

// Exe satisfies the Activity interface.
func (se *ScrollActivity) Exe(scroll *sc.Scroll) (exitCode int, err error) {
	for scroll.NextCodeLine() {
	
		firstCol := 1
		snip := sc.Snippet{
			Code: scroll.Code,
			Row: scroll.Number,
			Col: firstCol,
		}
	
		switch {
		case snip.HasAssignOperator():
			onAssignment(scroll)
		}
	}
	
	return exitCode, err
}

// Vars satisfies the activity interface
func (se *ScrollActivity) Vars() map[string]sc.VooValue {
	return se.vars
}

// onAssignment handles a line of scroll that assigns something
// to a variable.
func onAssignment(scroll *sc.Scroll) {
	varNames, _ := assignmentCleave(scroll.Code)
			
	for _, v := range varNames {
		scroll.PrintlnWithLineNum(v)
	}
			
	// TODO: Parse the value being assigned
	
}

// assignmentCleave splits a scroll line that performs an assignment
// into an array of variable names and the statement or expression.
func assignmentCleave(line string) (varNames []string, statOrExpr string) {
	parts := strings.SplitN(line, "=", 2)
	statOrExpr = parts[1]
	varNames = strings.Split(parts[0], ",")
	for i, v := range varNames {
		varNames[i] = strings.TrimSpace(v)
	}
	return
}
