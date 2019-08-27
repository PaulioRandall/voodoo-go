package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// parseAssign parses an assignment expression.
func parseAssign(parent *tree.Tree, in []token.Token, i int) (*tree.Tree, error) {

	tr := &tree.Tree{
		Kind:   tree.KD_ASSIGN,
		Token:  in[i],
		Parent: parent,
	}

	err := parseAssignChildren(tr, in[:i], in[i+1:])
	if err != nil {
		return nil, err
	}

	return tr, nil
}

func parseAssignChildren(parent *tree.Tree, left, right []token.Token) (err error) {
	parent.Left, err = parseIds(parent, left)
	if err != nil {
		return err
	}

	parent.Right, err = parseOperands(parent, right)
	return
}
