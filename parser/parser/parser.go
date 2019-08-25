package parser

import (
	"errors"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// TODO: Make a simple parser for `x <- 1`
// TODO: It requires:
// TODO: - a tree structure with nodes

// Parse parses the input statement into a parse tree.
func Parse(in []token.Token) (*tree.Tree, error) {
	tr := tree.New()

	for i, tk := range in {
		switch {
		case tr.Left == nil && tk.Type == token.TT_ID:
			tr.Left = tree.New()
			tr.Left.Token = tk
			tr.Left.Kind = tree.KD_ID
		default:
			m := "Token `" + in[i].Val + "` does not match any parsing rules"
			return nil, errors.New(m)
		}
	}

	return tr, nil
}
