package symbol

// TokItr provides a way to iterate token arrays.
type TokItr struct {
	index  int
	length int
	array  []Token
}

// increment adds one to the iterators array index
// counter.
func (itr *TokItr) increment() {
	itr.index += 1
}

// NewTokItr creates a new Token iterator.
func NewTokItr(ts []Token) *TokItr {
	return &TokItr{
		length: len(ts),
		array:  ts,
	}
}

// Index returns the iterators current counter value.
func (itr *TokItr) Index() int {
	return itr.index
}

// Length returns the total length of the iterators array.
func (itr *TokItr) Length() int {
	return itr.length
}

// Copy returns a copy of the token iterator.
func (itr *TokItr) Copy() *TokItr {
	return &TokItr{
		index:  itr.index,
		length: itr.length,
		array:  itr.array,
	}
}

// HasNext returns true if there are items remaining to be
// iterated.
func (itr *TokItr) HasNext() bool {
	return itr.index < itr.length
}

// NextTok returns the next token if there is one else it
// returns nil.
func (itr *TokItr) NextTok() *Token {
	if itr.HasNext() {
		defer itr.increment()
		return &itr.array[itr.index]
	}
	return nil
}

// IndexOf returns the index of the last token with the
// specified type within the remaining, uniterated, token array.
// -1 is returned if no match could be found.
func (itr *TokItr) IndexOf(t SymbolType) int {
	for i := itr.index; i < itr.length; i++ {
		if itr.array[i].Type == t {
			return i
		}
	}
	return -1
}

// RIndexOf returns the index of the last token with the
// specified type within the remaining, uniterated, token array.
// -1 is returned if no match could be found.
func (itr *TokItr) RIndexOf(t SymbolType) int {
	i := itr.length
	for i > itr.index {
		i--
		if itr.array[i].Type == t {
			return i
		}
	}
	return -1
}

// MoveTo moves the iterators counter to the specified index.
func (itr *TokItr) MoveTo(i int) {
	itr.index = i
}
