package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
)

func TestMatchArithmetic_1(t *testing.T) {
	p := mock([]token.Token{
		dummy(`+`, token.TK_ADD),
		dummy(`2`, token.TK_NUMBER),
	})

	m := matchArithmetic(p, 0)
	assert.True(t, m)
}

func TestMatchArithmetic_2(t *testing.T) {
	p := mock([]token.Token{
		dummy(`x`, token.TK_ID),
		dummy(`*`, token.TK_MULTIPLY),
		dummy(`y`, token.TK_ID),
	})

	m := matchArithmetic(p, 1)
	assert.True(t, m)
}
