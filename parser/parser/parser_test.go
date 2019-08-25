package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse_Rule_0(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 0, ``, token.TT_SPACE),
	}

	tr, err := Parse(in)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
}

//Rule 1:
//  Predicate:   The left node has no kind
//               AND the subject token has the IDENTIFIER type.
//  Consequence: Place the subject token in the left node
//               AND assign the left node the IDENTIFIER kind.
func TestParse_Rule_1(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
	}

	tr, err := Parse(in)
	require.Nil(t, err)
	require.NotNil(t, tr)

	require.NotNil(t, tr.Left)
	assert.Equal(t, tr.Left.Token, in[0])
	assert.Equal(t, tr.Left.Kind, tree.KD_ID)
}

//Rule 2:
//  Predicate:   The left node has the IDENTIFIER kind
//               AND the subject token has the ASSIGNMENT type.
//  Consequence: Place the subject token in the current node
//               AND assign the current node the ASSIGNMENT kind.
func TestParse_Rule_2(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
	}

	tr, err := Parse(in)
	require.Nil(t, err)
	require.NotNil(t, tr)

	require.NotNil(t, tr.Left)
	assert.Equal(t, tr.Left.Token, in[0])
	assert.Equal(t, tr.Left.Kind, tree.KD_ID)

	assert.Equal(t, tr.Token, in[1])
	assert.Equal(t, tr.Kind, tree.KD_ASSIGN)
}

//Rule 3:
//  Predicate:   The left node has the IDENTIFIER kind
//               AND the current node has the ASSIGNMENT kind
//               AND the subject token has the NUMBER type.
//  Consequence: Place the subject token in the right node
//               AND assign the right node the OPERAND kind.
func TestParse_Rule_3(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 5, 6, `1`, token.TT_NUMBER),
	}

	tr, err := Parse(in)
	require.Nil(t, err)
	require.NotNil(t, tr)

	require.NotNil(t, tr.Left)
	assert.Equal(t, tr.Left.Token, in[0])
	assert.Equal(t, tr.Left.Kind, tree.KD_ID)

	assert.Equal(t, tr.Token, in[1])
	assert.Equal(t, tr.Kind, tree.KD_ASSIGN)

	require.NotNil(t, tr.Right)
	assert.Equal(t, tr.Right.Token, in[2])
	assert.Equal(t, tr.Right.Kind, tree.KD_OPERAND)
}
