package analyser

import (
	//"strconv"
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
		dummyTok(`a`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}

	assert.Equal(t, 0, indexOf(in, 0, symbol.IDENTIFIER_EXPLICIT))
	assert.Equal(t, 1, indexOf(in, 0, symbol.CURVED_BRACE_OPEN))
	assert.Equal(t, 2, indexOf(in, 1, symbol.IDENTIFIER_EXPLICIT))
	assert.Equal(t, -1, indexOf(in, 2, symbol.CURVED_BRACE_OPEN))
}

func TestRIndexOf(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`a`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	last := len(in) - 1

	assert.Equal(t, 2, rIndexOf(in, last, symbol.IDENTIFIER_EXPLICIT))
	assert.Equal(t, 1, rIndexOf(in, last, symbol.CURVED_BRACE_OPEN))
	assert.Equal(t, 2, rIndexOf(in, 2, symbol.IDENTIFIER_EXPLICIT))
	assert.Equal(t, 0, rIndexOf(in, 1, symbol.IDENTIFIER_EXPLICIT))
	assert.Equal(t, -1, rIndexOf(in, last, symbol.LITERAL_STRING))
}

func TestContainsType(t *testing.T) {
	in := []symbol.Token{}
	assert.False(t, containsType(in, symbol.IDENTIFIER_EXPLICIT))
	in = append(in, dummyTok(`(`, symbol.IDENTIFIER_EXPLICIT))
	assert.True(t, containsType(in, symbol.IDENTIFIER_EXPLICIT))

	in = []symbol.Token{}
	assert.False(t, containsType(in, symbol.IDENTIFIER_EXPLICIT, symbol.KEYWORD_SPELL))
	in = append(in, dummyTok(`)`, symbol.CURVED_BRACE_CLOSE))
	in = append(in, dummyTok(`spell`, symbol.KEYWORD_SPELL))
	assert.False(t, containsType(in, symbol.IDENTIFIER_EXPLICIT))
	assert.True(t, containsType(in, symbol.KEYWORD_SPELL))
	in = append(in, dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT))
	assert.True(t, containsType(in, symbol.IDENTIFIER_EXPLICIT, symbol.KEYWORD_SPELL))
	assert.False(t, containsType(in, symbol.CURVED_BRACE_OPEN))
}

func TestFindParen_1(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	a, z := findParen(in)
	assert.Equal(t, 0, a)
	assert.Equal(t, 2, z)

	in = []symbol.Token{
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, z)

	in = []symbol.Token{
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, 3, a)
	assert.Equal(t, 4, z)

	in = []symbol.Token{
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
	}
	a, z = findParen(in)
	assert.Equal(t, 3, a)
	assert.Equal(t, 5, z)
}

func TestFindParen_2(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
	}
	a, z := findParen(in)
	assert.Equal(t, 0, a)
	assert.Equal(t, -1, z)

	in = []symbol.Token{
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, -1, a)
	assert.Equal(t, 1, z)

	in = []symbol.Token{
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, 2, a)
	assert.Equal(t, 3, z)
}

func TestExpandParen_1(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`x`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, symbol.ASSIGNMENT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}

	exp_outer := []symbol.Token{
		dummyTok(`x`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, symbol.ASSIGNMENT),
		dummyTok(`#1`, symbol.IDENTIFIER_IMPLICIT),
	}

	exp_inner := []symbol.Token{
		dummyTok(`#1`, symbol.IDENTIFIER_IMPLICIT),
		dummyTok(`<-`, symbol.ASSIGNMENT),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
	}

	outer, inner, err := expandParen(in, 1)
	require.Nil(t, err)
	assert.Equal(t, exp_outer, outer)
	assert.Equal(t, exp_inner, inner)
}

func TestExpandParen_2(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`x`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, symbol.ASSIGNMENT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`語`, symbol.IDENTIFIER_EXPLICIT),
	}

	outer, inner, err := expandParen(in, 1)
	assert.Nil(t, outer)
	assert.Nil(t, inner)
	assert.NotNil(t, err)
}

func TestExpandParens(t *testing.T) {
	in := []symbol.Token{
		dummyTok(`x`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, symbol.ASSIGNMENT),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`a`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`*`, symbol.MULTIPLY),
		dummyTok(`(`, symbol.CURVED_BRACE_OPEN),
		dummyTok(`b`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`+`, symbol.ADD),
		dummyTok(`c`, symbol.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
		dummyTok(`)`, symbol.CURVED_BRACE_CLOSE),
	}

	exp := [][]symbol.Token{
		[]symbol.Token{
			dummyTok(`#1`, symbol.IDENTIFIER_IMPLICIT),
			dummyTok(`<-`, symbol.ASSIGNMENT),
			dummyTok(`b`, symbol.IDENTIFIER_EXPLICIT),
			dummyTok(`+`, symbol.ADD),
			dummyTok(`c`, symbol.IDENTIFIER_EXPLICIT),
		},
		[]symbol.Token{
			dummyTok(`#2`, symbol.IDENTIFIER_IMPLICIT),
			dummyTok(`<-`, symbol.ASSIGNMENT),
			dummyTok(`a`, symbol.IDENTIFIER_EXPLICIT),
			dummyTok(`*`, symbol.MULTIPLY),
			dummyTok(`#1`, symbol.IDENTIFIER_IMPLICIT),
		},
		[]symbol.Token{
			dummyTok(`x`, symbol.IDENTIFIER_EXPLICIT),
			dummyTok(`<-`, symbol.ASSIGNMENT),
			dummyTok(`#2`, symbol.IDENTIFIER_IMPLICIT),
		},
	}

	act, err := expandParens(in)
	require.Nil(t, err)
	assert.Equal(t, exp, act)
}

/*
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
				symbol.Token{`x`, 1, 2, 0, symbol.IDENTIFIER_EXPLICIT},
				symbol.Token{`)`, 2, 3, 0, symbol.CURVED_BRACE_CLOSE},
			},
			Expect: [][]symbol.Token{
				[]symbol.Token{
					symbol.Token{`x`, 1, 2, 0, symbol.IDENTIFIER_EXPLICIT},
				},
			},
		},

			expBracketsTest{
				Input: []symbol.Token{
					symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
					symbol.Token{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
					symbol.Token{`(`, 3, 4, 0, symbol.CURVED_BRACE_OPEN},
					symbol.Token{`y`, 4, 5, 0, symbol.IDENTIFIER_EXPLICIT},
					symbol.Token{`)`, 5, 6, 0, symbol.CURVED_BRACE_CLOSE},
				},
				Expect: [][]symbol.Token{
					[]symbol.Token{
						symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
						symbol.Token{`<-`, 1, 3, 0, symbol.ASSIGNMENT},
						symbol.Token{`y`, 4, 5, 0, symbol.IDENTIFIER_EXPLICIT},
					},
				},
			},

	}
}
*/
