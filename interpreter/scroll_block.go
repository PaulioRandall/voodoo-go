
package interpreter

import (
	"fmt"
	"strings"
)

// ScrollBlockExe represents the scroll itself as a block executor.
type ScrollBlockExe struct {
	//Variables map[string]VoodooValue			// Currently used variables
}

// executeLines continues execution of the scroll lines at the root
// of the scroll until an error, an exit scroll command, or the end
// of file is encountered.
func executeLines(scroll *Scroll) {
	for scroll.NextCodeLine() {
	
		firstCol := 1
		snip := Snippet{
			Code: scroll.Code,
			Row: scroll.Number,
			Col: firstCol,
		}
	
		switch {
		case snip.HasAssignOperator():
			onAssignment(scroll)
		}
	}
}

// onAssignment handles a line of scroll that assigns something
// to a variable.
func onAssignment(scroll *Scroll) {
	varNames, _ := assignmentCleave(scroll.Code)
			
	for _, v := range varNames {
		printLineNumber(scroll.Index)
		fmt.Println(v)
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
