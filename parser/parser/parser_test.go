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

/*
Predicate:   The left node has no kind
             AND the subject token has the IDENTIFIER type.
Consequence: Place the subject token in the left node
             AND assign the left node the IDENTIFIER kind.
*/
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
