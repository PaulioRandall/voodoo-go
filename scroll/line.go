package scroll

import (
	"github.com/PaulioRandall/voodoo-go/shared"
)

// Line represents a line in a scroll.
type Line struct {
	Index int    // Current line index
	Num   int    // Current line number
	Val   string // Current line as a string
}

// String returns the lines value.
func (line Line) String() string {
	return line.Val
}

// Print pretty prints the line.
func (line Line) Println() {
	shared.PrintlnWithLineNum(line.Index, line.Val)
}
