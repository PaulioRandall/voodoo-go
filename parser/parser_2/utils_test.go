package parser_2

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	in := []token.Token{
		token.OfType(token.TT_ID),
		token.OfType(token.TT_ASSIGN),
		token.OfType(token.TT_NUMBER),
		token.OfType(token.TT_ADD),
		token.OfType(token.TT_NUMBER),
	}

	assert.Equal(t, -1, indexOf(in, token.TT_SPACE))
	assert.Equal(t, 0, indexOf(in, token.TT_ID))
	assert.Equal(t, 1, indexOf(in, token.TT_ASSIGN))
	assert.Equal(t, 2, indexOf(in, token.TT_NUMBER))
}
