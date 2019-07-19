package lexeme

// SymItr provides a way to iterate symbol arrays.
type SymItr struct {
	index  int
	length int
	array  []Symbol
}

// NewSymItr creates a new symbol iterator.
func NewSymItr(s []Symbol) *SymItr {
	return &SymItr{
		length: len(s),
		array:   s,
	}
}

// Length returns the total length of the iterators array.
func (itr *SymItr) Length() int {
	return itr.length
}

// HasNext returns true if there are any symbols yet to be
// iterated.
func (itr *SymItr) HasNext() bool {
	if itr.index < itr.length {
		return true
	}
	return false
}
