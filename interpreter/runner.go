
package interpreter

import (
	"fmt"
	"os"
	"time"
	"strings"
	"regexp"
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

// Execute runs the voodoo scroll.
func Execute(scroll *Scroll, scrollArgs []string) (exitCode int, err error) {
	
	scroll.JumpToLine(1) // Ignore first line
	
	for scroll.NextCodeLine() {
		
		// Note: The assignment operator '=' must have whitespace on both sides
		// of it. It is only used for assignment, therefore, if the operator
		// exists in a line it must be assigning something to a variable.
		isAssignment, err := regexp.MatchString(`\s=\s`, scroll.Code)
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

// compilerBug writes a compiler bug to output then exits the program
// with code 1.
func compilerBug(scroll *Scroll, msg string) {
	fmt.Println("[COMPILER BUG]")
	info := fmt.Sprintf("\t...when parsing line '%d' of '%s'", scroll.Number, scroll.File)
	fmt.Println(info)
	fmt.Print("\t..." + msg)
	os.Exit(1)
}

/******************************************************************************
	github.com/PaulioRandall/go-cookies/cookies
******************************************************************************/
/*
// TrimWhitespace removes all white space from a string.
func StripWhitespace(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		if !unicode.IsSpace(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
*/

// StopWatch represents a process timer with a few common operations.
type StopWatch struct {
	Started time.Time
	Stopped time.Time
}

// Start starts the stop watch overwritting any previously start time.
func (sw *StopWatch) Start() {
	sw.Started = time.Now().UTC()
}

// Stop stops the stop watch overwriting any previous stop time.
func (sw *StopWatch) Stop() {
	sw.Stopped = time.Now().UTC()
}

// Lap restarts the stop watch but returns a copy of the stop watch with the
// previous laps times.
func (sw *StopWatch) Lap() StopWatch {
	sw.Stop()
	r := StopWatch{
		Started: sw.Started,
		Stopped: sw.Stopped,
	}
	sw.Started = sw.Stopped
	return r
}

// Elapsed returns the elapsed time between the started and stopped times.
func (sw *StopWatch) Elapsed() time.Duration {
	return sw.Stopped.Sub(sw.Started)
}

// PrintElapsed prints the elapsed time as returned by Elapsed(). 'radix' may be
// passed to print the result in a more appropriate manner.
//
// e.g. PrintElapsed(time.Millisecond) will print as milliseconds.
//
// e.g. PrintElapsed(10 * 1000 * 1000) will print using a custom radix,
// microseconds * 10 in this case.
func (sw *StopWatch) PrintElapsed(radix time.Duration) {
	t := sw.Elapsed()
	f := float64(t) / float64(radix)

	switch radix {
	case time.Nanosecond:
		fmt.Printf("%.3f ns\n", f)
	case time.Microsecond:
		fmt.Printf("%.3f us\n", f)
	case time.Millisecond:
		fmt.Printf("%.3f ms\n", f)
	case time.Second:
		fmt.Printf("%.3f s\n", f)
	case time.Minute:
		fmt.Printf("%.3f m\n", f)
	case time.Hour:
		fmt.Printf("%.3f hr\n", f)
	default:
		fmt.Printf("%.3f\n", f)
	}
}