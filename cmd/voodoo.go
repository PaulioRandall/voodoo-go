//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"runtime"
	"bufio"
	"strings"
)

// main is the entry point for this script. It wraps the standard Go format,
// build, test, run, and install operations specifically for this project.
func main() {
	clearTerminal()

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

// clearTerminal clears the terminal.
func clearTerminal() {
	p := runtime.GOOS
	switch p {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		panic("Platform '" + p + "' not currently supported")
	}
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
	
	for scroll.MoveToNextCodeLine() {
		scroll.PrintCode()
	}
}

/******************************************************************************
	github.com/PaulioRandall/voodoo-go/cmd/scroll
******************************************************************************/

// Scroll represents a scroll and it's current working state.
type Scroll struct {
	Lines []string						// Raw lines from the scroll
	Length int							// Length of the scroll
	Index int								// Current line index
	Number int							// Current line number
	Code string							// Code from current line 
	Comment string				// Comment from current line
}

// LoadScroll reads the lines of the scroll and creates a
// new Scroll instance for it.
func LoadScroll(path string) (scroll Scroll, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	err = scanner.Err()
	if err == nil {
		scroll = Scroll{
			Lines: lines,
			Length: len(lines),
		}
	}
	
	return
}

// MoveToNextCodeLine moves the line index counter to
// the next line that has executable code. True is returned
// if there are lines still to be executed.
func (scroll *Scroll) MoveToNextCodeLine() bool {
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

// prepareLine finds and trims the line indicated by the
// current index then sets it as the 'CurrentLine' .
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
	inLiteral := false
	prev := ""
	
	for i, r := range line {
		curr := string(r)
		
		switch {
		case curr == "\"" && prev != "\\":
			inLiteral = !inLiteral
		case inLiteral:
			// Do nothing!
		case curr == "/" && prev == "/":
			return i - 1
		}
		
		prev = curr
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

// PrintComment prints the comment in the current line.
func (scroll *Scroll) PrintComment() {
	scroll.PrintCommentAt(scroll.Index)
}

// PrintCommentAt prints the comment of the specified line.
func (scroll *Scroll) PrintCommentAt(index int) {
	comment := scroll.Comment
	printNumberedLine(index, comment)
}

// PrintCode prints the code in the current line.
func (scroll *Scroll) PrintCode() {
	scroll.PrintCodeAt(scroll.Index)
}

// PrintCodeAt prints the code of the specified line.
func (scroll *Scroll) PrintCodeAt(index int) {
	code := scroll.Code
	printNumberedLine(index, code)
}

// PrintLine prints the current line.
func (scroll *Scroll) PrintLine() {
	scroll.PrintLineAt(scroll.Index)
}

// PrintLineAt prints the specified line.
func (scroll *Scroll) PrintLineAt(index int) {
	line := scroll.Lines[index]
	printNumberedLine(index, line)
}

// printNumberedLine prints the line number then the line
// contents.
func printNumberedLine(index int, line string) {
	num := index + 1
	out := fmt.Sprintf("%-3d: %v", num, line)
	fmt.Println(out)
}

// PrintLines prints all lines within the specified range.
func (scroll *Scroll) PrintLines(from int, to int) {
	switch {
	case from < 0, to < 0, from > to:
		e := fmt.Sprintf("Invalid line range: from %d to %d", from, to)
		panic(e)
	case to > scroll.Length:
		to = scroll.Length
	}
	
	lines := scroll.Lines[from:to]
	for i, v := range lines {
		printNumberedLine(i, v)
	}
}

// increment increments the line index counter by one.
func (scroll *Scroll) increment() {
	next := scroll.Index + 1
	scroll.JumpToLine(next)
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