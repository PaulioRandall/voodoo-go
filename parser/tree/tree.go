package tree

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Tree represents a parse tree.
type Tree struct {
	Kind   Kind
	Token  token.Token
	Parent *Tree
	Left   *Tree
	Right  *Tree
}

// New creates a new uninitialised tree.
func New() *Tree {
	return &Tree{}
}

// Copy performs a deep copy of the input tree. If the input is nil the output
// will be nil.
func Copy(tr *Tree, parent *Tree) *Tree {
	if tr == nil {
		return nil
	}

	cp := &Tree{
		Kind:   tr.Kind,
		Token:  tr.Token,
		Parent: parent,
	}

	cp.Left = Copy(tr.Left, cp)
	cp.Right = Copy(tr.Right, cp)
	return cp
}
