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
