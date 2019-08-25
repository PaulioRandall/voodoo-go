package tree

// Tree represents a parse tree.
type Tree interface {

	// Trunk returns the root node of the tree.
	Trunk() *Node
}

// New creates a new Tree implementation.
func New() Tree {
	return tree_1{}
}

// tree_1 is a private implementation of the Tree interface.
type tree_1 struct {
	trunk *Node
}

// Trunk satisfies the Tree interface.
func (tree tree_1) Trunk() *Node {
	return tree.trunk
}
