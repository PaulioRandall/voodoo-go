package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExprDelimFinder(t *testing.T) {
	in := []token.Token{
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_ID),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_ID),
		token.OfType(token.TT_ADD),
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_SPELL),
		token.OfType(token.TT_CURVED_OPEN),
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_ID),
		token.OfType(token.TT_VALUE_DELIM),
		token.OfType(token.TT_STRING),
		token.OfType(token.TT_CURVED_CLOSE),
	}

	exp := []int{
		1,
		3,
		7,
	}

	finder := exprDelimFinder{}
	act, err := finder.Find(in)

	require.Nil(t, err)
	require.NotNil(t, act)
	assert.Equal(t, exp, act)
}
