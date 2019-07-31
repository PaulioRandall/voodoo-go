package parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
)

// Expression represents a node within a parse tree.
type Expression interface {

	// Evaluate executes the expression within the given
	// context.
	Evaluate(*Context) (Value, fault.Fault)
}
