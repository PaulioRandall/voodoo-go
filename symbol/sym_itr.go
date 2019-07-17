package symbol


// SymItr provides a way to iterate symbol arrays.
type SymItr struct {
	index  int
	length int
	syms   []Symbol
}

// NewSymItr creates a new symbol iterator.
func NewSymItr(syms []Symbol) *SymItr {
	return &SymItr{
		length: len(syms),
		syms:   syms,
	}
}

// HasNext returns true if there are items still to be iterated.
func (itr *SymItr) HasNext() bool {
	i := itr.index + 1
	if i < itr.length {
		return true
	}
	return false
}

// increment increments the index counter.
func (itr *SymItr) increment() {
	itr.index = itr.index + 1
}

// Skip the next symbol in the array by incrementing the
// iterator index without returning anything.
func (itr *SymItr) Skip() {
	itr.increment()
}

// Next returns the next symbol in the array and increases the
// iterator index. Nil is returned if no symbols remain.
func (itr *SymItr) Next() *Symbol {
	defer itr.increment()
	return itr.Peek()
}

// Peek returns the next symbol in the array without
// incrementing the iterator index.
func (itr *SymItr) Peek() *Symbol {
	if itr.HasNext() {
		i := itr.index + 1
		return &itr.syms[i]
	}
	return nil
}
