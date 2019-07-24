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

// Length returns the total length of the iterators array.
func (itr *TokItr) Length() int {
	return itr.length
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

// IndexOf returns the first index of the token with the
// specified within the remaining, uniterated, token array.
// -1 is returned if no match could be found.
func (itr *TokItr) IndexOf(t SymbolType) int {
	for i := itr.index; i < itr.length; i++ {
		if itr.array[i].Type == t {
			return i
		}
	}
	return -1
}
