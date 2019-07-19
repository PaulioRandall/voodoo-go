package lexeme

// LexItr provides a way to iterate lexeme arrays.
type LexItr struct {
	index  int
	length int
	array  []Lexeme
}

// NewLexItr creates a new Lexeme iterator.
func NewLexItr(ls []Lexeme) *LexItr {
	return &LexItr{
		length: len(ls),
		array:   ls,
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
