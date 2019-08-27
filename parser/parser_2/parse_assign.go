package parser_2

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

	var err error

	tr.Left, err = parseIds(tr, in[:i])
	if err != nil {
		return nil, err
	}

	tr.Right, err = parseOperands(tr, in[i+1:])
	if err != nil {
		return nil, err
	}

	return tr, nil
}
