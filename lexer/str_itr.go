package lexer

import (
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// StrItr provides a way to iterate strings.
type StrItr struct {
	index  int
	length int
	str    string
}

// NewStrItr creates a new string iterator.
func NewStrItr(str string) *StrItr {
	return &StrItr{
		length: len(str),
		str:    str,
	}
}

// PrevIndex returns the index of the previous rune.
func (itr *StrItr) PrevIndex() int {
	return itr.index - 1
}

// NextIndex returns the index of the next rune.
func (itr *StrItr) NextIndex() int {
	return itr.index
}

// increment increments the index counter.
func (itr *StrItr) increment() {
	itr.index = itr.index + 1
}

// Skip the next rune by incrementing the iterator index
// without returning anything.
func (itr *StrItr) Skip() {
	i := itr.index
	itr.bugIfOutOfBounds(i)
	itr.increment()
}

// HasNext returns true if there are runes still to be iterated.
func (itr *StrItr) HasNext() bool {
	if itr.index < itr.length {
		return true
	}
	return false
}

// Next returns the next rune and increases the iterator index.
func (itr *StrItr) Next() rune {
	defer itr.increment()
	return itr.Peek()
}

// Peek returns the next rune without incrementing the iterator
// index.
func (itr *StrItr) Peek() rune {
	i := itr.index
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// HasAsatte returns true if there are at least two more runes
// still to be iterated.
func (itr *StrItr) HasAsatte() bool {
	i := itr.index + 1
	if i < itr.length {
		return true
	}
	return false
}

// PeekAsatte returns the rune after the next rune without
// incrementing the iterator index.
func (itr *StrItr) PeekAsatte() rune {
	i := itr.index + 1
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// HasPrev returns true if at least call to Next() has occurred.
func (itr *StrItr) HasPrev() bool {
	if itr.index > 0 {
		return true
	}
	return false
}

// PeekPrev returns the previous rune without decrementing
// the iterator index.
func (itr *StrItr) PeekPrev() rune {
	i := itr.index - 1
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// HasOtotoi returns true if at least two calls to Next() have
// been made.
func (itr *StrItr) HasOtotoi() bool {
	i := itr.index - 2
	if i > 0 {
		return true
	}
	return false
}

// PeekOtotoi returns the rune before the previous one without
// decrementing the iterator index.
func (itr *StrItr) PeekOtotoi() rune {
	i := itr.index - 2
	itr.bugIfOutOfBounds(i)
	return rune(itr.str[i])
}

// bugIfOutOfBounds will print error message and exit compilation if there
// are no items left in the string.
func (itr *StrItr) bugIfOutOfBounds(i int) {
	if i < 0 || i >= itr.length {
		sh.CompilerBug(-1, "Iterator call to Next(), Peek(), or AsattePeek() but no items remain")
	}
}
