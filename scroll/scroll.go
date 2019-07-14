package scroll

import (
	"fmt"

	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// Scroll represents a scroll.
type Scroll struct {
	File   string   // File path to the scroll
	Lines  []string // Raw lines from the scroll
	Length int      // Length of the scroll
}

// LoadScroll reads the lines of the scroll and creates a
// new Scroll instance for it.
func LoadScroll(path string) (*Scroll, error) {
	lines, err := sh.ReadLines(path)
	if err != nil {
		return nil, err
	}
	sc := &Scroll{
		File:   path,
		Lines:  lines,
		Length: len(lines),
	}
	return sc, nil
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
		Num:   i + 1,
		Val:   v,
	}
}

// IsEndOfScroll returns true if the the end of the scroll has been
// reached.
func (scroll *Scroll) IsEndOfScroll(index int) bool {
	if index < 0 {
		sh.CompilerBug(index+1, "How can a line index be negative?!")
	} else if index >= scroll.Length {
		return true
	}
	return false
}

// PrintlnLines prints all lines within the specified range.
func (scroll *Scroll) PrintlnLines(from int, to int) {
	switch {
	case from < 0, from > to, to > scroll.Length:
		e := fmt.Sprintf("Invalid line range: from %d to %d", from, to)
		sh.CompilerBug(-1, e)
	}

	lines := scroll.Lines[from:to]
	for i, v := range lines {
		sh.PrintlnWithLineNum(i, v)
	}
}
