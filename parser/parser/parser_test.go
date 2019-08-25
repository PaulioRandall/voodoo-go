package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
)

func TestParse_Rule_0(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 0, ``, token.TT_SPACE),
	}
	tree, err := Parse(in)
	assert.Nil(t, tree)
	assert.NotNil(t, err)
}
