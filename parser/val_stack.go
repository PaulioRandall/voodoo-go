package parser

// NewValStack returns a new initialised ValStack.
func NewValStack() *ValStack {
	return &ValStack{
		stack: []Token{},
	}
}

// Array returns a copy of the backing slice.
func (s *ValStack) Array() []Token {
	out := make([]Token, len(s.stack))
	for i, v := range s.stack {
		out[i] = v
	}
	return out
}

// Len returns size of the stack.
func (s *ValStack) Len() int {
	return len(s.stack)
}

// Push appends the value to the top of the stack.
func (e *ValStack) Push(in Token) {
	e.stack = append(e.stack, in)
}

// Pop removes and returns the value at the top of
// the stack. If false is returned then the stack
// is empty.
func (e *ValStack) Pop() (Token, bool) {
	top := len(e.stack) - 1
	if top < 0 {
		return Token{}, false
	}

	out := e.stack[top]
	e.stack[top] = Token{}
	e.stack = e.stack[:top]
	return out, true
}

// Peek returns the value at the top of the stack
// without removing it. If false is returned then
// the stack is empty.
func (e *ValStack) Peek() (Token, bool) {
	top := len(e.stack) - 1
	if top < 0 {
		return Token{}, false
	}
	return e.stack[top], true
}

// Sink moves the top n items to the bottom of the
// stack maintaining the order of the shifted items.
// Note that this function will panic if the n is
// negative or the n > len.
func (s *ValStack) Sink(n int) {
	if n < 0 {
		panic("Non-negative input not expected")
	}

	x := len(s.stack) - n
	if x < 0 {
		panic("Not enough items on stack to sink")
	}

	s.stack = append(s.stack[x:], s.stack[:x]...)
}
