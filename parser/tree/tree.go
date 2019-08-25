package tree

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Tree represents a parse tree.
type Tree struct {
	Kind  TreeKind
	Token token.Token
	Left  *Tree
	Right *Tree
}

// New creates a new tree.
func New() *Tree {
	return &Tree{}
}
