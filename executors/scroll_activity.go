
package executors

import (
	"strings"
	
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// ScrollActivity represents the scroll itself as an activity.
type ScrollActivity struct {
	vars map[string]sc.VooValue
}

// NewScrollActivity returns a new scroll activity.
func NewScrollActivity() *ScrollActivity {
	return &ScrollActivity{
		vars: make(map[string]sc.VooValue),
	}
}

// ExeLine satisfies the Executor interface.
func (sa *ScrollActivity) ExeLine(scroll *sc.Scroll, line string) (sh.ExitCode, Executor, sh.ExeError) {
	exitCode := sh.OK
	next := Executor(sa)
	var err sh.ExeError = nil
	
	firstCol := 1
	snip := sc.Snippet{
		Code: line,
		Row: scroll.Number,
		Col: firstCol,
	}
	
	switch {
	case snip.HasAssignOperator():
		onAssignment(scroll)
	}
	
	return exitCode, next, err
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
