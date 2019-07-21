package symbol

// LexItr provides a way to iterate lexeme arrays.
type LexItr struct {
	index  int
	length int
	array  []Lexeme
}

// increment adds one to the iterators array index
// counter.
func (itr *LexItr) increment() {
	itr.index += 1
}

// NewLexItr creates a new Lexeme iterator.
func NewLexItr(ls []Lexeme) *LexItr {
	return &LexItr{
		length: len(ls),
		array:  ls,
	}
}

// Length returns the total length of the iterators array.
func (itr *LexItr) Length() int {
	return itr.length
}

// HasNext returns true if there are any lexemes yet to be
// iterated.
func (itr *LexItr) HasNext() bool {
	if itr.index < itr.length {
		return true
	}
	return false
}

// NextLex returns the next lexeme in the array and increments
// the iterators array counter.
func (itr *LexItr) NextLex() *Lexeme {
	if itr.HasNext() {
		defer itr.increment()
		return &itr.array[itr.index]
	}
	return nil
}
