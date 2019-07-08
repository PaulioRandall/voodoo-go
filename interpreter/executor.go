
package interpreter

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// LoadScroll reads the lines of the scroll and creates a
// new Scroll instance for it.
func LoadScroll(path string) (scroll *Scroll, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	lines, err := scanLines(file)
	if err == nil {
		scroll = newScroll(path, lines)
	}
	
	return
}

// scanLines reads in the lines of an opened file.
func scanLines(file *os.File) ([]string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Execute runs the voodoo scroll.
func Execute(scroll *Scroll, scrollArgs []string) (exitCode int, err error) {
	
	scroll.JumpToLine(1) // Ignore first line
	
	for scroll.NextCodeLine() {
	
		// TODO: Create a file and func that handles lines of ordinary scroll
		// TODO: code, i.e. those at the root of a scroll, spell or in a 'for' block but not
		// TODO: those within a 'when' block.
		
		isAssignment, err := Snippet(scroll.Code).isAssignment()
		if err != nil {
			compilerBug(scroll, "bad regex used for finding assignment operator")
		}
		
		if isAssignment {
			varNames, _ := assignmentCleave(scroll.Code)
			
			for _, v := range varNames {
				printLineNumber(scroll.Index)
				fmt.Println(v)
			}
			
			// TODO: Parse the value being assigned
		}
	}
	
	return 1, nil
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
