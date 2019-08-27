package parser_2

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/require"
)

func TestParse_1(t *testing.T) {
	in := []token.Token{
		token.OfType(token.TT_ID),
		token.OfType(token.TT_ASSIGN),
		token.OfType(token.TT_NUMBER),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: in[1],
	}

	exp.Left = &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  in[0],
		Parent: exp.Left,
	}
	exp.Right = &tree.Tree{
		Kind:   tree.KD_OPERAND,
		Token:  in[2],
		Parent: exp,
	}

	act, err := parse(nil, in)
	require.Nil(t, err)
	require.NotNil(t, act)
	assertTree(t, exp, act, "Trunk")
}
