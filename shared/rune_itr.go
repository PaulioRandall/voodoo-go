package shared

import (
	"strings"
	"unicode"
)

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

// TODO: Delete after refactoring.
func (itr *RuneItr) SetIndex(i int) {
	itr.index = i
}

// increment increments the iterators index.
func (itr *RuneItr) increment() {
	itr.index += 1
}

// Length returns the total number of runes.
func (itr *RuneItr) Length() int {
	return itr.length
}

// Index returns the index of the next rune.
func (itr *RuneItr) Index() int {
	return itr.index
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

// PeekRelRune returns the rune specified by the index calculated
// by offsetting the current index by the input. The offset
// may be negative to return previous runes. If no rune exists
// then the -1 will be returned.
func (itr *RuneItr) PeekRelRune(offset int) rune {
	if itr.HasRelRune(offset) {
		i := itr.index + offset
		return itr.runes[i]
	}
	return -1
}

// NextRune returns the next rune in the array and increments
// the iterators index.
func (itr *RuneItr) NextRune() rune {
	if itr.HasRelRune(0) {
		defer itr.increment()
		return itr.runes[itr.index]
	}

	return -1
}

// PeekRune returns the next rune in the array and but does NOT
// increment the iterators index.
func (itr *RuneItr) PeekRune() rune {
	if itr.HasRelRune(0) {
		return itr.runes[itr.index]
	}

	return -1
}

// HasNext returns true if there are any items left to
// iterate.
func (itr *RuneItr) HasNext() bool {
	if itr.index < itr.length {
		return true
	}
	return false
}

// IsNext returns true if the next rune is equal to the input
// rune. False is also returned if no more runes remain to be
// iterated.
func (itr *RuneItr) IsNext(r rune) bool {
	if itr.HasNext() {
		return itr.PeekRune() == r
	}
	return false
}

// IsNextIn returns true if there is a next rune and the next
// rune appears within the input string.
func (itr *RuneItr) IsNextIn(s string) bool {
	if itr.HasNext() {
		return strings.ContainsRune(s, itr.PeekRune())
	}
	return false
}

// IsNextLetter returns true if the next rune is in the unicode
// category 'L' for letter. False is also returned if no more
// runes remain to iterate.
func (itr *RuneItr) IsNextLetter() bool {
	if itr.HasNext() {
		return unicode.IsLetter(itr.PeekRune())
	}
	return false
}
