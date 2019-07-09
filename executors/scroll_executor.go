
package executors

import (
	"strings"
	"fmt"
	
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	sh "github.com/PaulioRandall/voodoo-go/shared"
)


type TokenType string

const (
	Keyword TokenType = "keyword"
	VarNames TokenType = "variable-names"
	Void TokenType = "void"
)

// ScrollExecutor represents the scroll itself as an activity.
type ScrollExecutor struct {
	vars map[string]sc.VooValue
}

// NewScrollExecutor returns a new scroll executor.
func NewScrollExecutor() *ScrollExecutor {
	return &ScrollExecutor{
		vars: make(map[string]sc.VooValue),
	}
}

// ExeLine satisfies the Executor interface.
func (sa *ScrollExecutor) Exe(scroll *sc.Scroll, line sc.Statement) (sh.ExitCode, Executor, sh.ExeError) {
	exitCode := sh.OK
	next := Executor(sa)
	var err sh.ExeError = nil
	
	// TODO: Process each rune one at a time
	
	switch {
	case line.IsAssignment():
		onAssignment(scroll)
	}
	
	return exitCode, next, err
}

// handle
func handle(stat sc.Statement) sc.Statement {	
	s := string(stat.Val)
	first := rune(s[0])
	
	if isLetter(first) {
		handleVariableNames(s)
	}
	
	// TODO: Return a new one statement that starts where
	// TODO: the current one finished processing.
	return stat
}

// handleVariableNames handles 
func handleVariableNames(s string) sh.ExeError {
	names := strings.Split(s, `,`)
	for i, v := range names {
		v = strings.TrimSpace(v)
		names[i] = v
		
		if !isLetter(rune(v[0])) {
			m := fmt.Sprintf("Expected first rune of variable name to be a letter '%v'", v)
			return sh.NewError(1, m)
		}
		
		for _, r := range v[1:] {
			if !isVariableRune(r) {
				m := fmt.Sprintf("Unexpected rune in variable name '%v'", v)
				return sh.NewError(1, m)
			}
		}
	
		fmt.Println("Variable name: " + v)
	}
	
	return nil
}

// isVariableRune returns true if the rune is allowed within a
// variable name.
func isVariableRune(r rune) bool {
	if isLetter(r) || isNumber(r) || isUnderscore(r) {
		return true
	}
	return false
}

// isLetter returns true if the rune is a letter.
func isLetter(r rune) bool {
	// (UTF-8) 65-90 == A-Z, 97-122 == a-z
	return (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
}

// isNumber returns true if the rune is a number.
func isNumber(r rune) bool {
	// (UTF-8) 48-57 == 0-9
	return r >= 48 && r <= 57
}

// isUnderscore returns true if the rune is an '_'.
func isUnderscore(r rune) bool {
	// (UTF-8) 95 == _
	return r == 95
}

// isAtSign returns true if the rune is an '@'.
func isAtSign(r rune) bool {
	// (UTF-8) 64 == @
	return r == 64
}

// isComma returns true if the rune is an ','.
func isComma(r rune) bool {
	// (UTF-8) 44 == ,
	return r == 44
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
