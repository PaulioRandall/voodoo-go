package tree

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Tree represents a parse tree.
type Tree struct {
	Kind  Kind
	Token token.Token
	isSet bool
	Left  *Tree
	Right *Tree
}

// New creates a new uninitialised tree.
func New() *Tree {
	return &Tree{}
}

// Set sets the field values of the tree.
func (tr *Tree) Set(tk token.Token, kind Kind) {
	if tr.isSet {
		panic("This tree's token and kind fields have already been set")
	}
	tr.Token = tk
	tr.Kind = kind
	tr.isSet = true
}

// SetLeft sets the field values for the left tree.
func (tr *Tree) SetLeft(tk token.Token, kind Kind) {
	if tr.Left != nil && tr.Left.isSet {
		panic("The left tree's token and kind fields have already been set")
	}
	tr.Left = &Tree{
		Token: tk,
		Kind:  kind,
		isSet: true,
	}
}

// SetRight sets the field values for the right tree.
func (tr *Tree) SetRight(tk token.Token, kind Kind) {
	if tr.Right != nil && tr.Right.isSet {
		panic("The right tree's token and kind fields have already been set")
	}
	tr.Right = &Tree{
		Token: tk,
		Kind:  kind,
		isSet: true,
	}
}

// Is checks to see if the input kind matches the tree's kind.
func (tr *Tree) Is(k Kind) bool {
	return k == KD_DONT_CARE || k == tr.Kind
}

// IsLeft checks to see if the input kind matches the left tree's kind.
func (tr *Tree) IsLeft(l Kind) bool {
	switch {
	case l == KD_DONT_CARE:
		return true
	case tr.Left == nil:
		return l == KD_UNDEFINED
	default:
		return l == tr.Left.Kind
	}
}

// IsRight checks to see if the input kind matches the right tree's kind.
func (tr *Tree) IsRight(r Kind) bool {
	switch {
	case r == KD_DONT_CARE:
		return true
	case tr.Right == nil:
		return r == KD_UNDEFINED
	default:
		return r == tr.Right.Kind
	}
}

// Are checks to see if the input kinds match the tree's field kinds.
func (tr *Tree) Are(l, k, r Kind) bool {
	return tr.IsLeft(l) && tr.Is(k) && tr.IsRight(r)
}
