package parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// Exe represents an executable instruction.
type Exe struct {
	Token   token.Token
	Params  int // Number of input parameters
	Returns int // Number of output parameters
}

// ExeStack represents a stack of instructions.
type ExeStack struct {
	stack []Exe
}

// Push appends the instruction to the top of the
// stack.
func (e *ExeStack) Push(in Exe) {
	e.stack = append(e.stack, in)
}

// Pop removes and returns the instruction at the
// top of the stack. If false is returned then the
// stack is empty.
func (e *ExeStack) Pop() (Exe, bool) {
	top := len(e.stack) - 1
	if top < 0 {
		return Exe{}, false
	}

	out := e.stack[top]
	e.stack[top] = Exe{}
	e.stack = e.stack[:top]
	return out, true
}

// Peek returns the instruction at the top of the
// stack without removing it. If false is returned
// then the stack is empty.
func (e *ExeStack) Peek() (Exe, bool) {
	top := len(e.stack) - 1
	if top < 0 {
		return Exe{}, false
	}
	return e.stack[top], true
}
