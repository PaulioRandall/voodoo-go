package shared

import (
	"strings"
)

// Rune representing UTF-8/ASCII NUL (Decimal: 0)
//const NUL_RUNE := 0

// RuneItr represents an iterator of strings.
type RuneItr struct {
	index  int    // Index of the next rune
	length int    // Total number of runes
	str    string // String representation of the rune array
	runes  []rune // Runes to iterate
}

// NewRuneItr creates a new iterator of the runes within
// the input string.
func NewRuneItr(str string) *RuneItr {
	runes := []rune(str)
	return &RuneItr{
		length: len(runes),
		str:    str,
		runes:  runes,
	}
}

// Length returns the total number of runes.
func (itr *RuneItr) Length() int {
	return itr.length
}

// HasRuneRelTo returns true if the index calculated by
// offsetting the current index by the input references
// a rune within the bounds of the rune array.
func (itr *RuneItr) HasRelRune(offset int) bool {
	i := itr.index + offset
	if i >= 0 && i < itr.length {
		return true
	}
	return false
}

// RelRune returns the rune specified by the index calculated
// by offsetting the current index by the input. The offset
// may be negative to return previous runes. If no rune exists
// then the -1 will be returned.
func (itr *RuneItr) RelRune(offset int) rune {
	if itr.HasRelRune(offset) {
		i := itr.index + offset
		return itr.runes[i]
	}
	return -1
}

// ##################################################
// ##################################################
// NEXT: Test the 2 functions above
// NEXT: Test using non-ascii rune strings too
// ##################################################
// ##################################################

// NextIndex returns the index of the next rune.
func (itr *RuneItr) NextIndex() int {
	return itr.index
}

// increment increments the index counter.
func (itr *RuneItr) increment() {
	itr.index = itr.index + 1
}

// Skip the next rune by incrementing the iterator index
// without returning anything.
func (itr *RuneItr) Skip() {
	i := itr.index
	itr.bugIfOutOfBounds(i)
	itr.increment()
}

// HasNext returns true if there are runes still to be iterated.
func (itr *RuneItr) HasNext() bool {
	if itr.index < itr.length {
		return true
	}
	return false
}

// Next returns the next rune and increases the iterator index.
func (itr *RuneItr) Next() rune {
	defer itr.increment()
	return itr.Peek()
}

// Peek returns the next rune without incrementing the iterator
// index.
func (itr *RuneItr) Peek() rune {
	i := itr.index
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// TODO: Replace with HasRel(int)
// HasAsatte returns true if there are at least two more runes
// still to be iterated.
func (itr *RuneItr) HasAsatte() bool {
	i := itr.index + 1
	if i < itr.length {
		return true
	}
	return false
}

// TODO: Replace with PeekRel(int)
// PeekAsatte returns the rune after the next rune without
// incrementing the iterator index.
func (itr *RuneItr) PeekAsatte() rune {
	i := itr.index + 1
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// TODO: Add IndexRel(int)

// TODO: remove
// HasPrev returns true if at least call to Next() has occurred.
func (itr *RuneItr) HasPrev() bool {
	if itr.index > 0 {
		return true
	}
	return false
}

// TODO: remove
// PeekPrev returns the previous rune without decrementing
// the iterator index.
func (itr *RuneItr) PeekPrev() rune {
	i := itr.index - 1
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// TODO: remove
// HasOtotoi returns true if at least two calls to Next() have
// been made.
func (itr *RuneItr) HasOtotoi() bool {
	i := itr.index - 2
	if i > 0 {
		return true
	}
	return false
}

// TODO: remove
// PeekOtotoi returns the rune before the previous one without
// decrementing the iterator index.
func (itr *RuneItr) PeekOtotoi() rune {
	i := itr.index - 2
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// NextIsIn returns true if there is a next rune and the next
// rune is in the input string.
func (itr *RuneItr) NextIsIn(s string) bool {
	if itr.HasNext() {
		return strings.ContainsRune(s, itr.Peek())
	}
	return false
}

// bugIfOutOfBounds will print error message and exit compilation if there
// are no items left in the string.
func (itr *RuneItr) bugIfOutOfBounds(i int) {
	if i < 0 || i >= itr.length {
		CompilerBug(-1, "Iterator call to Next(), Peek(), or AsattePeek() but no items remain")
	}
}
