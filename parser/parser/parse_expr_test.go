package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/require"
)

func NEW_TestParseExprs_1(t *testing.T) {
	in := []token.Token{
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_ADD),
		token.OfType(token.TT_NUMBER),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_OPERATION,
		Token: in[1],
	}
	exp.Left = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  in[0],
		Parent: exp,
	}
	exp.Right = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  in[2],
		Parent: exp,
	}

	act, err := parseExprs(nil, in)
	require.Nil(t, err)
	require.NotNil(t, act)
	assertTree(t, exp, act, "Trunk")
}
