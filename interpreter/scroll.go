
package interpreter

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

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