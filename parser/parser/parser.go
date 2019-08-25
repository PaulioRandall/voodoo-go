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
	for i, _ := range in {
		switch {
		default:
			m := "Token `" + in[i].Val + "` does not match any parsing rules"
			return nil, errors.New(m)
		}
	}

	return nil, nil
}
