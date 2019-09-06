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
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`1`, token.TT_NUMBER),
	})

	m := matchAssign(p)
	assert.True(t, m)
}

func TestMatchAssign_2(t *testing.T) {
	p := mock([]token.Token{
		dummy(`1`, token.TT_NUMBER),
	})

	m := matchAssign(p)
	assert.False(t, m)
}

func TestMatchAssign_3(t *testing.T) {
	p := mock([]token.Token{
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`_`, token.TT_VOID),
	})

	m := matchAssign(p)
	assert.True(t, m)
}

func TestMatchAssign_4(t *testing.T) {
	p := mock([]token.Token{
		dummy(`_`, token.TT_VOID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`4`, token.TT_NUMBER),
	})

	m := matchAssign(p)
	assert.True(t, m)
}

func TestMatchAssign_5(t *testing.T) {
	p := mock([]token.Token{
		dummy(`_`, token.TT_VOID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`true`, token.TT_BOOL),
	})

	m := matchAssign(p)
	assert.True(t, m)
}
