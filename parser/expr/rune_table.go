package expr

import (
	"strings"
)

// RuneTable represents a series of lines.
type RuneTable [][]rune

// NewRuneTable creates a new rune table.
func NewRuneTable(height, width int) RuneTable {
	rt := make([][]rune, height)
	for i, _ := range rt {
		rt[i] = make([]rune, width)
	}
	return rt
}

// Imprint overwrites the specified runes with the specified string.
func (rt RuneTable) Imprint(line, col int, s string) {
	ra := [][]rune(rt)
	for i, ru := range s {
		ra[line][col+i] = ru
	}
}

// Margin returns the number of cells remaining after applying the imprint. A
// negative number indicates Imprint() will fail with the absolute of that
// number indicating how much by.
func (rt RuneTable) Margin(col int, s string) int {
	ra := [][]rune(rt)
	return len(ra[0]) - col - len(s)
}

// Filler fills in any slot within the rune table that has the zero value with
// the rune specified.
func (rt *RuneTable) Filler(ru rune) {
	rt.Map(func(_, _ int, curr rune) rune {
		if curr == 0 {
			return ru
		}
		return curr
	})
}

// Map iterates the entire rune table from left to right, top to bottom applying
// the function to each cell. The index of the cell along with the current rune
// is passed into the function and the result is placed into the cell.
func (rt RuneTable) Map(f func(i, j int, ru rune) rune) {
	ra := [][]rune(rt)
	for i, line := range ra {
		for j, ru := range line {
			ra[i][j] = f(i, j, ru)
		}
	}
}

// String returns the rune table as a single string.
func (rt RuneTable) String() string {
	sb := strings.Builder{}
	ra := [][]rune(rt)
	size := len(ra)

	for i, line := range ra {
		for _, ru := range line {
			sb.WriteRune(ru)
		}

		if i+1 < size {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}
