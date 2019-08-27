package parser_2

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/require"
)

func TestParseIds(t *testing.T) {
	ids := []token.Token{
		token.OfType(token.TT_ID),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_ID),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_ID),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_UNION,
		Token: ids[1],
	}
	exp.Left = &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  ids[0],
		Parent: exp,
	}
	exp.Right = &tree.Tree{
		Kind:   tree.KD_UNION,
		Token:  ids[3],
		Parent: exp,
	}
	exp.Right.Left = &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  ids[2],
		Parent: exp.Right,
	}
	exp.Right.Right = &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  ids[4],
		Parent: exp.Right,
	}

	act, err := parseIds(nil, ids)
	require.Nil(t, err)
	require.NotNil(t, act)
	assertTree(t, exp, act, "Trunk")
}

func TestParseId(t *testing.T) {
	id := token.OfType(token.TT_ID)

	parent := &tree.Tree{
		Kind: tree.KD_ASSIGN,
	}
	exp := &tree.Tree{
		Kind:   tree.KD_ID,
		Token:  id,
		Parent: parent,
	}

	act, err := parseId(parent, id)
	require.Nil(t, err)
	require.NotNil(t, act)
	assertTree(t, exp, act, "Trunk")
}
