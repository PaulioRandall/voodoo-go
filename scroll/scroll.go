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
