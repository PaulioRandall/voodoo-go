package analyser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func dummyTok(s string, t symbol.SymbolType) symbol.Token {
	return symbol.Token{
		Val:  s,
		Type: t,
	}
}

func TestFindBracketPair_1(t *testing.T) {
	ls := []symbol.Token{
		dummyTok(`a`, symbol.IDENTIFIER),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	itr := symbol.NewTokItr(ls)

	o, c := findBracketPair(itr)
	assert.Equal(t, 1, o)
	assert.Equal(t, 3, c)
}

func TestFindBracketPair_2(t *testing.T) {
	ls := []symbol.Token{
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER),
	}
	itr := symbol.NewTokItr(ls)

	o, c := findBracketPair(itr)
	assert.Equal(t, 0, o)
	assert.Equal(t, -1, c)
}

func TestFindBracketPair_3(t *testing.T) {
	ls := []symbol.Token{
		dummyTok(`語`, symbol.IDENTIFIER),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	itr := symbol.NewTokItr(ls)

	o, c := findBracketPair(itr)
	assert.Equal(t, -1, o)
	assert.Equal(t, 1, c)
}

func TestFindBracketPair_4(t *testing.T) {
	ls := []symbol.Token{
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	itr := symbol.NewTokItr(ls)

	o, c := findBracketPair(itr)
	assert.Equal(t, 0, o)
	assert.Equal(t, 4, c)
}

func TestFindBracketPair_5(t *testing.T) {
	ls := []symbol.Token{
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	itr := symbol.NewTokItr(ls)

	o, c := findBracketPair(itr)
	assert.Equal(t, 0, o)
	assert.Equal(t, -1, c)
}

func TestFindBracketPair_6(t *testing.T) {
	ls := []symbol.Token{
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	itr := symbol.NewTokItr(ls)

	o, c := findBracketPair(itr)
	assert.Equal(t, 0, o)
	assert.Equal(t, 2, c)
}
