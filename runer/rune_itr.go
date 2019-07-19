package runer

import (
	"errors"
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
// the iterators index. -1 is returned if no runes remain to
// return.
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

// NextStr returns 'n' runes, as a string, in the array and adds
// 'n' to the iterators index. An error is returned if 'n' <= 0
// or the index + 'n' is out of bounds.
func (itr *RuneItr) NextStr(n int) (string, error) {
	end := itr.index + n
	if n <= 0 || end > itr.length {
		m := "Input is either zero, negative, or puts the iterator index out of bounds"
		return "", errors.New(m)
	}

	r := itr.runes[itr.index:end]
	itr.index = end
	return string(r), nil
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

// IsNextStr returns true if the input string matches the next
// set of runes  to appears within the iterator.
func (itr *RuneItr) IsNextStr(s string) bool {
	for i, ru := range []rune(s) {
		if itr.PeekRelRune(i) != ru {
			return false
		}
	}
	return true
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

// IsRelLetter returns true if the rune specified by the relative
// index is in the unicode category 'L' for letter. The relative
// index is calculated by offsetting the current index by the input.
// The offset may be negative to check previous runes. If no rune
// exists then the -1 will be returned. False is also returned if no
// more runes remain to iterate.
func (itr *RuneItr) IsRelLetter(offset int) bool {
  ru := itr.PeekRelRune(offset)
	return unicode.IsLetter(ru)
}

// IsNextDigit returns true if the next rune is in the unicode
// category 'Nd' for decimal digit. False is also returned if no
// more runes remain to iterate.
func (itr *RuneItr) IsNextDigit() bool {
	if itr.HasNext() {
		return unicode.IsDigit(itr.PeekRune())
	}
	return false
}

// IsNextSpace returns true if the next rune has the unicode
// property 'whitespace'. False is also returned if no more
// runes remain to iterate.
func (itr *RuneItr) IsNextSpace() bool {
	if itr.HasNext() {
		return unicode.IsSpace(itr.PeekRune())
	}
	return false
}

// RemainingStr returns all remaining runes as a string and sets
// iterator index to the index after the last rune.
func (itr *RuneItr) RemainingStr() string {
	r := ``
	if itr.HasNext() {
		r = itr.str[itr.index:]
	}
	itr.index = itr.length
	return r
}
