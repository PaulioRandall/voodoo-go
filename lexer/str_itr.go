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

// bugIfOutOfBounds will print error message and exit compilation if there
// are no items left in the string.
func (itr *StrItr) bugIfOutOfBounds(i int) {
	if i >= itr.length {
		sh.CompilerBug(-1, "Iterator call to Next(), Peek(), or AsattePeek() but no items remain")
	}
}
