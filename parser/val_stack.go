package parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// ValStack represents a stack of values.
type ValStack struct {
	stack []token.Token
}

// Push appends the value to the top of the stack.
func (e *ValStack) Push(in token.Token) {
	e.stack = append(e.stack, in)
}

// Pop removes and returns the value at the top of
// the stack. If false is returned then the stack
// is empty.
func (e *ValStack) Pop() (token.Token, bool) {
	top := len(e.stack) - 1
	if top < 0 {
		return token.Token{}, false
	}

	out := e.stack[top]
	e.stack[top] = token.Token{}
	e.stack = e.stack[:top]
	return out, true
}

// Peek returns the value at the top of the stack
// without removing it. If false is returned then
// the stack is empty.
func (e *ValStack) Peek() (token.Token, bool) {
	top := len(e.stack) - 1
	if top < 0 {
		return token.Token{}, false
	}
	return e.stack[top], true
}
