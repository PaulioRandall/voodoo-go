package parser

// Push appends the instruction to the top of the
// stack.
func (s *ExeStack) Push(in Exe) {
	s.stack = append(s.stack, in)
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
func (s *ExeStack) Sink(n int) {
	if n < 1 {
		panic("Non-positive input not expected")
	}

	x := len(s.stack) - n
	if x < 0 {
		panic("Not enough items on stack to sink")
	}

	s.stack = append(s.stack[x:], s.stack[:x]...)
}
