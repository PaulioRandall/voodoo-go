package ll_parser

// stack represents a stack of expressions and tokens
// used while parsing a statement.
type stack struct {
	a    []interface{}
	size int
	i    int
}

// newStack creates a new stack with the specified size.
func newStack(size int) *stack {
	return &stack{
		a:    make([]interface{}, size),
		size: size,
	}
}

// PanicIfFull panics if the stack is full.
func (s *stack) panicIfFull() {
	if s.i+1 >= s.size {
		panic(`Expression stack is full, might need to allocate more space`)
	}
}

// panicIfEmpty panics if the stack is empty.
func (s *stack) panicIfEmpty() {
	if s.i-1 < 0 {
		panic(`Expression stack is empty`)
	}
}

// Push pushes an item onto the top of the stack.
func (s *stack) Push(item interface{}) {
	s.panicIfFull()
	s.a[s.i] = item
	s.i++
}

// Pop pops an item off the top of the stack.
func (s *stack) Pop() interface{} {
	s.panicIfEmpty()
	s.i--
	item := s.a[s.i]
	s.a[s.i] = nil
	return item
}
