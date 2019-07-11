
package scroll

import (
	"fmt"
	"strings"
	
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// Line represents a line in a scroll.
type Line struct {
	Index int								// Current line index
	Num int									// Current line number
	Val string							// Current line as a string
}

// Scroll represents a scroll.
type Scroll struct {
	File string							// File path to the scroll
	Lines []string					// Raw lines from the scroll
	Length int							// Length of the scroll
}

// NewScroll creates a new Scroll instance.
func NewScroll(file string, lines []string) *Scroll {
	return &Scroll{
		File: file,
		Lines: lines,
		Length: len(lines),
	}
}

// Print prints the line
func (line *Line) Print() {
	printlnWithLineNum(line.Index, line.Val)
}

// Next returns the line after the input line returning nil
// if there are no more lines left. If nil is supplied then
// the first line is returned. Whitespace, 'ain't nobody
// got time for that'; all leading and trailing whitespace is
// trimmed from the value of the returned line.
func (scroll *Scroll) Next(prev *Line) *Line {
	if prev == nil {
		return scroll.getLine(0)
	}

	i := prev.Index + 1
	if scroll.IsEndOfScroll(i) {
		return nil
	}
	
	return scroll.getLine(i)
}

// getLine returns the line specified by the index.
func (scroll *Scroll) getLine(i int) *Line {
	v := scroll.Lines[i]
	return &Line{
		Index: i,
		Num: i + 1,
		Val: strings.TrimSpace(v),
	}
}

// IsEndOfScroll returns true if the the end of the scroll has been
// reached.
func (scroll *Scroll) IsEndOfScroll(index int) bool {
	if index < 0 {
		sh.CompilerBug(index + 1, "How can a line index be negative?!")
	} else if index >= scroll.Length {
		return true
	}
	return false
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
		printlnWithLineNum(i, v)
	}
}

// printlnWithLineNum prints the line number then the argument.
func printlnWithLineNum(i int, s string) {
	printLineNumber(i)
	fmt.Println(s)
}

// printLineNumber prints the line number but does not add
// a new line character to the end.
func printLineNumber(index int) {
	num := index + 1
	out := fmt.Sprintf("%-3d: ", num)
	fmt.Print(out)
}
