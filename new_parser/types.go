package new_parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Token
type Token token.Token

// Fault
type Fault fault.Fault

// NodeType represents the type of a node.
type NodeType int

// Node represents an operation or value.
type Node struct {
	Type   NodeType
	Tokens []Token
}

// ParseTree represents a statment as a parse tree.
type ParseTree struct {
	Node
}

const (
	UNDEFINED NodeType = iota
)
