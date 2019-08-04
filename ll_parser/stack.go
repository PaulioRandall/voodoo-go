package ll_parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

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

// Len returns the length of the stack.
func (s *stack) Len() int {
	return s.i
}

// PeekAt returns the item at the top index - offset.
func (s *stack) PeekAt(offset int) interface{} {
	i := s.i - 1 - offset
	if i < 0 || i >= s.size {
		panic(`Index + 1 - offset is out of bounds`)
	}
	return s.a[i]
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

// PopTokens pops n tokens off the top of the stack returning
// them as a new array.
func (s *stack) PopTokens(n int) []token.Token {
	s.panicIfEmpty()
	b := make([]token.Token, n)
	for i := 0; i < n; i++ {
		b[i] = s.Pop().(token.Token)
	}
	return b
}
