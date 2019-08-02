package parser

// EmptyValStack returns a new initialised but empty
// ExeStack.
func EmptyValStack() *ValStack {
	return &ValStack{
		stack: []Token{},
	}
}

// NewValStack returns a new initialised ValStack.
func NewValStack(s []Token) *ValStack {
	return &ValStack{
		stack: s,
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
// Returns its self.
func (s *ValStack) Push(in Token) *ValStack {
	s.stack = append(s.stack, in)
	return s
}

// Pop removes and returns the value at the top of
// the stack. If false is returned then the stack
// is empty.
func (s *ValStack) Pop() (Token, bool) {
	top := len(s.stack) - 1
	if top < 0 {
		return Token{}, false
	}

	out := s.stack[top]
	s.stack[top] = Token{}
	s.stack = s.stack[:top]
	return out, true
}

// Peek returns the value at the top of the stack
// without removing it. If false is returned then
// the stack is empty.
func (s *ValStack) Peek() (Token, bool) {
	top := len(s.stack) - 1
	if top < 0 {
		return Token{}, false
	}
	return s.stack[top], true
}

// Sink moves the top n items to the bottom of the
// stack maintaining the order of the shifted items.
// Note that this function will panic if n is
// negative or n > len. Returns its self.
func (s *ValStack) Sink(n int) *ValStack {
	if n < 0 {
		panic("Non-negative input not expected")
	}

	x := len(s.stack) - n
	if x < 0 {
		panic("Not enough items on stack to sink")
	}

	s.stack = append(s.stack[x:], s.stack[:x]...)
	return s
}

// Reverse tips the stack upside down so the bottom
// becomes the top and the top becomes the bottom.
// Returns its self.
func (s *ValStack) Reverse() *ValStack {
	for i := len(s.stack)/2 - 1; i >= 0; i-- {
		opp := len(s.stack) - 1 - i
		s.stack[i], s.stack[opp] = s.stack[opp], s.stack[i]
	}
	return s
}
