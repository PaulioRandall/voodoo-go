package parser_2

import (
	"errors"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// parseIds parses a value delimitered set of identifiers into a parse tree.
func parseIds(parent *tree.Tree, ids []token.Token) (*tree.Tree, error) {
	if len(ids) == 1 {
		return parseId(parent, ids[0])
	}

	if len(ids)%2 == 0 {
		m := "Expected odd number of ID delimitered tokens"
		return nil, errors.New(m)
	}

	if ids[1].Type != token.TT_VALUE_DELIM {
		m := "Expected VALUE_DELIM token at ids[1]"
		return nil, errors.New(m)
	}

	tr := &tree.Tree{
		Kind:   tree.KD_UNION,
		Token:  ids[1],
		Parent: parent,
	}

	var err error

	tr.Left, err = parseId(tr, ids[0])
	if err != nil {
		return nil, err
	}

	tr.Right, err = parseIds(tr, ids[2:])
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// parseId parses a single identifier into a parse tree.
func parseId(parent *tree.Tree, tk token.Token) (*tree.Tree, error) {
	if tk.Type != token.TT_ID {
		m := "Expected ID token at ids[0]"
		return nil, errors.New(m)
	}

	tr := &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  tk,
		Parent: parent,
	}

	return tr, nil
}
