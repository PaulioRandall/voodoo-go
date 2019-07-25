package analyser

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummyTok(s string, t symbol.SymbolType) symbol.Token {
	return symbol.Token{
		Val:  s,
		Type: t,
	}
}

func TestIndexOf(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`a`, symbol.IDENTIFIER),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}

	assert.Equal(t, 0, indexOf(in, 0, symbol.IDENTIFIER))
	assert.Equal(t, 1, indexOf(in, 0, symbol.CURVED_BRACE_OPEN))
	assert.Equal(t, 2, indexOf(in, 1, symbol.IDENTIFIER))
	assert.Equal(t, -1, indexOf(in, 2, symbol.CURVED_BRACE_OPEN))
}

func TestRIndexOf(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`a`, symbol.IDENTIFIER),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	last := len(in) - 1

	assert.Equal(t, 2, rIndexOf(in, last, symbol.IDENTIFIER))
	assert.Equal(t, 1, rIndexOf(in, last, symbol.CURVED_BRACE_OPEN))
	assert.Equal(t, 2, rIndexOf(in, 2, symbol.IDENTIFIER))
	assert.Equal(t, 0, rIndexOf(in, 1, symbol.IDENTIFIER))
	assert.Equal(t, -1, rIndexOf(in, last, symbol.STRING))
}

func TestExpandBrackets(t *testing.T) {
	for i, tc := range expBracketsTests() {
		t.Log("expandBrackets() test case: " + strconv.Itoa(i+1))
		out, err := expandBrackets(tc.Input)

		if tc.ExpectErr != nil {
			assert.NotNil(t, err)
			require.Nil(t, out)
			assert.NotEmpty(t, err.Error())
			assert.Equal(t, tc.ExpectErr.Line(), err.Line())
			assert.Equal(t, tc.ExpectErr.Col(), err.Col())
		} else {
			assert.Nil(t, err)
			require.NotNil(t, out)
			assert.Equal(t, tc.Expect, out)
		}
	}
}

type expBracketsTest struct {
	Input     []symbol.Token
	Expect    [][]symbol.Token
	ExpectErr AnaError
}

func expBracketsTests() []expBracketsTest {
	return []expBracketsTest{
		expBracketsTest{
			Input: []symbol.Token{
				symbol.Token{`(`, 0, 1, 0, symbol.CURVED_BRACE_OPEN},
				symbol.Token{`x`, 1, 2, 0, symbol.IDENTIFIER},
				symbol.Token{`)`, 2, 3, 0, symbol.CURVED_BRACE_CLOSE},
			},
			Expect: [][]symbol.Token{
				[]symbol.Token{
					symbol.Token{`x`, 1, 2, 0, symbol.IDENTIFIER},
				},
			},
		},
	}
}
