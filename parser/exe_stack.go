package parser

// EmptyExeStack returns a new initialised but empty
// ExeStack.
func EmptyExeStack() *ExeStack {
	return &ExeStack{
		stack: []Exe{},
	}
}

// NewExeStack returns a new initialised ExeStack.
func NewExeStack(s []Exe) *ExeStack {
	return &ExeStack{
		stack: s,
	}
}

// Array returns a copy of the backing slice.
func (s *ExeStack) Array() []Exe {
	out := make([]Exe, len(s.stack))
	for i, v := range s.stack {
		out[i] = v
	}
	return out
}

// Len returns size of the stack.
func (s *ExeStack) Len() int {
	return len(s.stack)
}

// Push appends the instruction to the top of the
// stack. Returns its self.
func (s *ExeStack) Push(in Exe) *ExeStack {
	s.stack = append(s.stack, in)
	return s
}

// Pop removes and returns the instruction at the
// top of the stack. If false is returned then the
// stack is empty.
func (s *ExeStack) Pop() (Exe, bool) {
	top := len(s.stack) - 1
	if top < 0 {
		return Exe{}, false
	}

	out := s.stack[top]
	s.stack[top] = Exe{}
	s.stack = s.stack[:top]
	return out, true
}

// Peek returns the instruction at the top of the
// stack without removing it. If false is returned
// then the stack is empty.
func (s *ExeStack) Peek() (Exe, bool) {
	top := len(s.stack) - 1
	if top < 0 {
		return Exe{}, false
	}
	return s.stack[top], true
}

// Sink moves the top n items to the bottom of the
// stack maintaining the order of the shifted items.
// Note that this function will panic if n is
// negative or n > len. Returns its self.
func (s *ExeStack) Sink(n int) *ExeStack {
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
func (s *ExeStack) Reverse() *ExeStack {
	for i := len(s.stack)/2 - 1; i >= 0; i-- {
		opp := len(s.stack) - 1 - i
		s.stack[i], s.stack[opp] = s.stack[opp], s.stack[i]
	}
	return s
}
