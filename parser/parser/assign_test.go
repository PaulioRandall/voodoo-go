package parser

import (
	"testing"

	//"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	//"github.com/PaulioRandall/voodoo-go/utils"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestMatchAssign_1(t *testing.T) {
	p := mock([]token.Token{
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`1`, token.TK_NUMBER),
	})

	m := matchAssign(p, 0)
	assert.True(t, m)
}

func TestMatchAssign_2(t *testing.T) {
	p := mock([]token.Token{
		dummy(`1`, token.TK_NUMBER),
	})

	m := matchAssign(p, 0)
	assert.False(t, m)
}

func TestMatchAssign_3(t *testing.T) {
	p := mock([]token.Token{
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`_`, token.TK_VOID),
	})

	m := matchAssign(p, 0)
	assert.True(t, m)
}

func TestMatchAssign_4(t *testing.T) {
	p := mock([]token.Token{
		dummy(`_`, token.TK_VOID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`4`, token.TK_NUMBER),
	})

	m := matchAssign(p, 0)
	assert.True(t, m)
}

func TestMatchAssign_5(t *testing.T) {
	p := mock([]token.Token{
		dummy(`x`, token.TK_ID),
		dummy(`,`, token.TK_DELIM),
		dummy(`y`, token.TK_VOID),
		dummy(`,`, token.TK_DELIM),
		dummy(`z`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`4`, token.TK_NUMBER),
		dummy(`,`, token.TK_DELIM),
		dummy(`Dragonfly`, token.TK_STRING),
		dummy(`,`, token.TK_DELIM),
		dummy(`_`, token.TK_VOID),
	})

	m := matchAssign(p, 0)
	assert.True(t, m)
}
