package scroll

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/shared"
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
	lines, err := shared.ReadLines(path)
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
func (sc *Scroll) Next(prev *Line) *Line {
	if prev == nil {
		return sc.getLine(0)
	}

	i := prev.Index + 1
	if sc.IsEndOfScroll(i) {
		return nil
	}

	return sc.getLine(i)
}

// getLine returns the line specified by the index.
func (sc *Scroll) getLine(i int) *Line {
	v := sc.Lines[i]
	return &Line{
		Index: i,
		Num:   i + 1,
		Val:   v,
	}
}

// IsEndOfScroll returns true if the the end of the scroll has been
// reached.
func (sc *Scroll) IsEndOfScroll(index int) bool {
	if index < 0 {
		panic("How can a line index be negative?!")
	} else if index >= sc.Length {
		return true
	}
	return false
}

// PrintlnLines prints all lines within the specified range.
func (sc *Scroll) PrintlnLines(from int, to int) {
	switch {
	case from < 0, from > to, to > sc.Length:
		e := fmt.Sprintf("Invalid line range: from %d to %d", from, to)
		panic(e)
	}

	lines := sc.Lines[from:to]
	for i, v := range lines {
		shared.PrintlnWithLineNum(i, v)
	}
}
