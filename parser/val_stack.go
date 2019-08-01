package parser

// ValStack represents a stack of values.
type ValStack struct {
	stack []Token
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
