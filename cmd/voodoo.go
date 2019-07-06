//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"os"
	"time"
	"bufio"
	"strings"
	"regexp"
)

// main is the entry point for this script. It wraps the standard Go format,
// build, test, run, and install operations specifically for this project.
func main() {
	stopWatch := StopWatch{}
	stopWatch.Start()
	fmt.Printf("Started\t%v\n\n", stopWatch.Started.UTC())

	// Don't abstract the build workflows! They are more readable and extendable
	// this way.
	option := getArgument()
	switch option {
	case "run":
		voodooRun()

	default:
		badSyntax()
	}

	stopWatch.Stop()
	fmt.Printf("\nDone\t")
	stopWatch.PrintElapsed(time.Microsecond)

	os.Exit(0)
}

// getArgument returns the argument passed that represents the operation to
// perform.
func getArgument() string {
	args := os.Args[1:]
	if len(args) < 2 {
		badSyntax()
	}
	return args[0]
}

// badSyntax prints the scripts syntax to console then exits the application
// with code 1.
func badSyntax() {
	syntax := `syntax options:
1) ./voodoo.exe run [scroll-name]`

	fmt.Println(syntax + "\n")
	os.Exit(1)
}

// voodooRun runs the voodoo scroll.
func voodooRun() {
	scrollPathArgIndex := 2
	scrollPath := os.Args[scrollPathArgIndex]
	//scrollArgs := os.Args[3:]
	
	scroll, err := LoadScroll(scrollPath)
	if err != nil {
		panic(err)
	}
	
	lastIgnoredLine := 2
	scroll.JumpToLine(lastIgnoredLine)
	
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
	github.com/PaulioRandall/voodoo-go/cmd/variable
******************************************************************************/

// TODO: Move this code to it's own package

// ValueType represents the type of a voodoo value.
type ValueType int

// Declaration of Value types.
const (
	BoolType ValueType = iota + 1
	NumType
	StrType
	ListType
	ObjType
	FuncType
)

// KeyValuePair represents with a key value pair.
type KeyValuePair struct {
	Key VoodooValue
	Value VoodooValue
}

// VoodooValue represents a value within the scroll.
type VoodooValue struct {
	ValueType ValueType
	BoolValue bool
	NumValue float64
	StrValue string
	ListValue []VoodooValue
	ObjValue []KeyValuePair
	FuncValue []string
}

/******************************************************************************
	github.com/PaulioRandall/voodoo-go/cmd/scroll
******************************************************************************/

// TODO: Move this code to it's own package

// Scroll represents a scroll and it's current state.
type Scroll struct {
	// Scroll
	File string								// File path to the scroll
	Lines []string						// Raw lines from the scroll
	Length int							// Length of the scroll
	// Line state
	Index int								// Current line index
	Number int							// Current line number
	Code string							// Code from current line 
	Comment string					// Comment from current line
	// Variable state
	Variables map[string]VoodooValue			// Currently used variables
}

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

// newScroll creates a new Scroll instance.
func newScroll(file string, lines []string) *Scroll {
	return &Scroll{
		File: file,
		Lines: lines,
		Length: len(lines),
	}
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

// NextCodeLine moves the line index counter to
// the next line that has executable code. True is returned
// if there are lines still to be executed.
func (scroll *Scroll) NextCodeLine() bool {
	for {
		scroll.increment()
		if scroll.IsEndOfScroll() {
			return false
		}
		
		scroll.prepLine()
		if scroll.IsCodeLine() {
			return true
		}
	}
}

// increment increments the line index counter by one.
func (scroll *Scroll) increment() {
	next := scroll.Index + 1
	scroll.JumpToLine(next)
}

// prepLine finds, splits, trims, and sets the code and comments
// of the line indicated by the current index.
func (scroll *Scroll) prepLine() {
	line := scroll.Lines[scroll.Index]
	code, comment := cleaveLine(line)
	scroll.Code = code
	scroll.Comment = comment
}

// cleaveLine splits the line into the code part and comment part.
// Both parts are trimmed before being returned.
func cleaveLine(line string) (code string, comment string) {
	cleaveIndex := findCleavePoint(line)
	runes := []rune(line)
	
	if cleaveIndex == -1 {
		code = prepLinePart(runes)
		comment = ""
	} else {
		code = prepLinePart(runes[:cleaveIndex])
		comment = prepLinePart(runes[cleaveIndex:])
	}
	
	return
}

// prepLinePart prepares the code or comment part of a
// line for processing by removing redudant whitespace
// and converting it to a string.
func prepLinePart(runes []rune) string {
	str := string(runes)
	return strings.TrimSpace(str)
}

// findCleavePoint finds the rune index where a comment starts
// in a line. -1 is returned if there is no point.
func findCleavePoint(line string) int {
	prevIndex := 0
	prev := ""
	
	return  findIndexInLine(line, func(i int, r rune) int {
		s := string(r)
		
		if (i - 1) == prevIndex && prev == "/" && s == "/" {
			return prevIndex
		}
		
		prevIndex = i
		prev = s
		return -1
	})
}

// findIndexInLine searches a line to find whatever the 'onEachRune'
// function is searching for. Runes within string literals are not
// passed to the function so no special handling is required within
// the supplied function. 'onEachRune' function accepts the index
// as the first value and the rune as the second while a non-negative
// index if the required index has been found.
func findIndexInLine(line string, onEachRune func(int, rune) int) int {
	inLiteral := false
	prev := ""
	
	for i, r := range line {
		s := string(r)
		wasInLiteral := false
		
		if inLiteral && prev != "\\" && s == "\"" {
			inLiteral = false
			wasInLiteral = true
		}
		
		if !inLiteral {
			index := onEachRune(i, r)
			if index > -1 {
				return index
			}
		}
		
		if !inLiteral && !wasInLiteral && s == "\"" {
			inLiteral = true
		}
		
		prev = s
	}
	
	return -1
}

// IsCodeLine returns true if the currrent line contains
// executable code.
func (scroll *Scroll) IsCodeLine() bool {
	if scroll.Code != "" {
		return true
	}
	return false
}

// IsCommentLine returns true if the current line is a
// comment line.
func (scroll *Scroll) IsCommentLine() bool {
	if scroll.Comment != ""{
		return true
	}
	return false
}

// HasMoreLines returns true if the there are lines
// in the scroll still to be executed.
func (scroll *Scroll) HasMoreLines() bool {
	return !scroll.IsEndOfScroll()
}

// IsEndOfScroll returns true if the the end of the
// scroll has been reached.
func (scroll *Scroll) IsEndOfScroll() bool {
	i := scroll.Index
	if i < 0 || i >= scroll.Length {
		return true
	}
	return false
}

// PrintlnComment prints the comment in the current line.
func (scroll *Scroll) PrintlnComment() {
	scroll.PrintlnCommentAt(scroll.Index)
}

// PrintlnCommentAt prints the comment of the specified line.
func (scroll *Scroll) PrintlnCommentAt(index int) {
	comment := scroll.Comment
	printlnNumberedLine(index, comment)
}

// PrintlnCode prints the code in the current line.
func (scroll *Scroll) PrintlnCode() {
	scroll.PrintlnCodeAt(scroll.Index)
}

// PrintlnCodeAt prints the code of the specified line.
func (scroll *Scroll) PrintlnCodeAt(index int) {
	code := scroll.Code
	printlnNumberedLine(index, code)
}

// PrintlnLine prints the current line.
func (scroll *Scroll) PrintlnLine() {
	scroll.PrintlnLineAt(scroll.Index)
}

// PrintlnLineAt prints the specified line.
func (scroll *Scroll) PrintlnLineAt(index int) {
	line := scroll.Lines[index]
	printlnNumberedLine(index, line)
}

// PrintlnLines prints all lines within the specified range.
func (scroll *Scroll) PrintlnLines(from int, to int) {
	switch {
	case from < 0, to < 0, from > to:
		e := fmt.Sprintf("Invalid line range: from %d to %d", from, to)
		panic(e)
	case to > scroll.Length:
		to = scroll.Length
	}
	
	lines := scroll.Lines[from:to]
	for i, v := range lines {
		printlnNumberedLine(i, v)
	}
}

// printlnNumberedLine prints the line number then the line
// contents.
func printlnNumberedLine(index int, line string) {
	printLineNumber(index)
	fmt.Println(line)
}

// printLineNumber prints the line number but does not add
// a new line character to the end.
func printLineNumber(index int) {
	num := index + 1
	out := fmt.Sprintf("%-3d: ", num)
	fmt.Print(out)
}

// JumpToLine sets the next line cursor to the specified line index.
func (scroll *Scroll) JumpToLine(index int) {
	scroll.Index = index
	scroll.Number = index + 1
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