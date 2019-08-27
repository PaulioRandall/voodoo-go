package parser

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

	err := checkOperands(in)
	if err != nil {
		return nil, err
	}

	tr := &tree.Tree{
		Kind:   tree.KD_UNION,
		Token:  in[1],
		Parent: parent,
	}

	err = parseChildOperands(tr, in[0], in[2:])
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// checkOperands checks the operand token slice has the expected state.
func checkOperands(in []token.Token) error {
	if len(in)%2 == 0 {
		m := "Expected odd number of operand delimitered tokens"
		return errors.New(m)
	}

	if in[1].Type != token.TT_VALUE_DELIM {
		m := "Expected VALUE_DELIM token at in[1]"
		return errors.New(m)
	}

	return nil
}

// parseChildOperands parses the remaining operands.
func parseChildOperands(parent *tree.Tree, left token.Token, right []token.Token) (err error) {
	parent.Left, err = parseOperand(parent, left)
	if err != nil {
		return
	}

	parent.Right, err = parseOperands(parent, right)
	return
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
