package parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Token
type Token token.Token

// Fault
type Fault fault.Fault

// Exe represents an executable instruction.
type Exe struct {
	Token   Token
	Params  int // Number of input parameters
	Returns int // Number of output parameters
}

// ExeStack represents a stack of instructions.
type ExeStack struct {
	stack []Exe
}

// ValStack represents a stack of values.
type ValStack struct {
	stack []Token
}
