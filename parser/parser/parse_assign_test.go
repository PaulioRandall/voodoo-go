package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/require"
)

func TestParseAssign_1(t *testing.T) {
	in := []token.Token{
		token.OfType(token.TT_ID),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_ID),
		token.OfType(token.TT_ASSIGN),
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_NUMBER),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: in[3],
	}

	exp.Left = &tree.Tree{
		Kind:   tree.KD_UNION,
		Token:  in[1],
		Parent: exp,
	}
	exp.Left.Left = &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  in[0],
		Parent: exp.Left,
	}
	exp.Left.Right = &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  in[2],
		Parent: exp.Left,
	}

	exp.Right = &tree.Tree{
		Kind:   tree.KD_UNION,
		Token:  in[5],
		Parent: exp,
	}
	exp.Right.Left = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  in[4],
		Parent: exp.Right,
	}
	exp.Right.Right = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  in[6],
		Parent: exp.Right,
	}

	act, err := parseAssign(nil, in, 3)
	require.Nil(t, err)
	require.NotNil(t, act)
	assertTree(t, exp, act, "Trunk")
}
