package parser_2

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/require"
)

func TestParseOperands(t *testing.T) {
	ids := []token.Token{
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_NUMBER),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_UNION,
		Token: ids[1],
	}
	exp.Left = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  ids[0],
		Parent: exp,
	}
	exp.Right = &tree.Tree{
		Kind:   tree.KD_UNION,
		Token:  ids[3],
		Parent: exp,
	}
	exp.Right.Left = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  ids[2],
		Parent: exp.Right,
	}
	exp.Right.Right = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  ids[4],
		Parent: exp.Right,
	}

	act, err := parseOperands(nil, ids)
	require.Nil(t, err)
	require.NotNil(t, act)
	assertTree(t, exp, act, "Trunk")
}

func TestParseOperand(t *testing.T) {
	id := token.OfType(token.TT_ID)

	parent := &tree.Tree{
		Kind: tree.KD_ASSIGN,
	}
	exp := &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  id,
		Parent: parent,
	}

	act, err := parseOperand(parent, id)
	require.Nil(t, err)
	require.NotNil(t, act)
	assertTree(t, exp, act, "Trunk")
}
