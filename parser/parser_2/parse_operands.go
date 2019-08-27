package parser_2

import (
	"errors"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// parseOperands parses a value delimitered set of identifiers into a parse tree.
func parseOperands(parent *tree.Tree, in []token.Token) (*tree.Tree, error) {
	if len(in) == 1 {
		return parseOperand(parent, in[0])
	}

	if len(in)%2 == 0 {
		m := "Expected odd number of operand delimitered tokens"
		return nil, errors.New(m)
	}

	if in[1].Type != token.TT_VALUE_DELIM {
		m := "Expected VALUE_DELIM token at in[1]"
		return nil, errors.New(m)
	}

	tr := &tree.Tree{
		Kind:   tree.KD_UNION,
		Token:  in[1],
		Parent: parent,
	}

	var err error

	tr.Left, err = parseOperand(tr, in[0])
	if err != nil {
		return nil, err
	}

	tr.Right, err = parseOperands(tr, in[2:])
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// parseOperand parses a single operand into a parse tree.
func parseOperand(parent *tree.Tree, tk token.Token) (*tree.Tree, error) {
	switch tk.Type {
	case token.TT_ID:
	case token.TT_NUMBER:
	default:
		m := "Expected token from the OPERAND set"
		return nil, errors.New(m)
	}

	tr := &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  tk,
		Parent: parent,
	}

	return tr, nil
}
