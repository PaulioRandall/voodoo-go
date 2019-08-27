package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertTree(t *testing.T, exp *tree.Tree, act *tree.Tree, node string) {
	if exp == nil {
		assert.Nil(t, act, `Expected '%s == nil'`, node)
		return
	}

	if !token.AssertToken(t, &exp.Token, &act.Token) {
		t.Logf(`%s: Not the expected token`, node)
	}
	assert.Equal(t, exp.Kind, act.Kind, `%s: Not the expected kind`, node)

	assertTree(t, exp.Left, act.Left, node+`.Left`)
	assertTree(t, exp.Right, act.Right, node+`.Right`)
}

func TestParse_1(t *testing.T) {
	in := []token.Token{
		token.OfType(token.TT_ID),
		token.OfType(token.TT_ASSIGN),
		token.OfType(token.TT_NUMBER),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: in[1],
		Left: &tree.Tree{
			Kind:  tree.KD_ID,
			Token: in[0],
		},
		Right: &tree.Tree{
			Kind:  tree.KD_OPERAND,
			Token: in[2],
		},
	}

	act, err := Parse(in)
	require.Nil(t, err)
	assertTree(t, exp, act, `Trunk`)
}
