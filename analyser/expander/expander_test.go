package expander

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummyTok(s string, t token.TokenType) token.Token {
	return token.Token{
		Val:  s,
		Type: t,
	}
}

func TestIndexOf(t *testing.T) {
	in := []token.Token{
		dummyTok(`a`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}

	assert.Equal(t, 0, indexOf(in, 0, token.IDENTIFIER_EXPLICIT))
	assert.Equal(t, 1, indexOf(in, 0, token.PAREN_CURVY_OPEN))
	assert.Equal(t, 2, indexOf(in, 1, token.IDENTIFIER_EXPLICIT))
	assert.Equal(t, -1, indexOf(in, 2, token.PAREN_CURVY_OPEN))
}

func TestRIndexOf(t *testing.T) {
	in := []token.Token{
		dummyTok(`a`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}
	last := len(in) - 1

	assert.Equal(t, 2, rIndexOf(in, last, token.IDENTIFIER_EXPLICIT))
	assert.Equal(t, 1, rIndexOf(in, last, token.PAREN_CURVY_OPEN))
	assert.Equal(t, 2, rIndexOf(in, 2, token.IDENTIFIER_EXPLICIT))
	assert.Equal(t, 0, rIndexOf(in, 1, token.IDENTIFIER_EXPLICIT))
	assert.Equal(t, -1, rIndexOf(in, last, token.LITERAL_STRING))
}

func TestContainsType(t *testing.T) {
	in := []token.Token{}
	assert.False(t, containsType(in, token.IDENTIFIER_EXPLICIT))
	in = append(in, dummyTok(`(`, token.IDENTIFIER_EXPLICIT))
	assert.True(t, containsType(in, token.IDENTIFIER_EXPLICIT))

	in = []token.Token{}
	assert.False(t, containsType(in, token.IDENTIFIER_EXPLICIT, token.KEYWORD_FUNC))
	in = append(in, dummyTok(`)`, token.PAREN_CURVY_CLOSE))
	in = append(in, dummyTok(`func`, token.KEYWORD_FUNC))
	assert.False(t, containsType(in, token.IDENTIFIER_EXPLICIT))
	assert.True(t, containsType(in, token.KEYWORD_FUNC))
	in = append(in, dummyTok(`語`, token.IDENTIFIER_EXPLICIT))
	assert.True(t, containsType(in, token.IDENTIFIER_EXPLICIT, token.KEYWORD_FUNC))
	assert.False(t, containsType(in, token.PAREN_CURVY_OPEN))
}

func TestFindParen_1(t *testing.T) {
	in := []token.Token{
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}
	a, z := findParen(in)
	assert.Equal(t, 0, a)
	assert.Equal(t, 2, z)

	in = []token.Token{
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, z)

	in = []token.Token{
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, 3, a)
	assert.Equal(t, 4, z)

	in = []token.Token{
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
	}
	a, z = findParen(in)
	assert.Equal(t, 3, a)
	assert.Equal(t, 5, z)
}

func TestFindParen_2(t *testing.T) {
	in := []token.Token{
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
	}
	a, z := findParen(in)
	assert.Equal(t, 0, a)
	assert.Equal(t, -1, z)

	in = []token.Token{
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, -1, a)
	assert.Equal(t, 1, z)

	in = []token.Token{
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}
	a, z = findParen(in)
	assert.Equal(t, 2, a)
	assert.Equal(t, 3, z)
}

func TestExpandParen_1(t *testing.T) {
	in := []token.Token{
		dummyTok(`x`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, token.ASSIGNMENT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}

	exp_outer := []token.Token{
		dummyTok(`x`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, token.ASSIGNMENT),
		dummyTok(`#1`, token.IDENTIFIER_IMPLICIT),
	}

	exp_inner := []token.Token{
		dummyTok(`#1`, token.IDENTIFIER_IMPLICIT),
		dummyTok(`<-`, token.ASSIGNMENT),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
	}

	outer, inner, err := expandParen(in, 1)
	require.Nil(t, err)
	assert.Equal(t, exp_outer, outer)
	assert.Equal(t, exp_inner, inner)
}

func TestExpandParen_2(t *testing.T) {
	in := []token.Token{
		dummyTok(`x`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, token.ASSIGNMENT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`語`, token.IDENTIFIER_EXPLICIT),
	}

	outer, inner, err := expandParen(in, 1)
	assert.Nil(t, outer)
	assert.Nil(t, inner)
	assert.NotNil(t, err)
}

func TestExpandParens(t *testing.T) {
	in := []token.Token{
		dummyTok(`x`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`<-`, token.ASSIGNMENT),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`a`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`*`, token.CALC_MULTIPLY),
		dummyTok(`(`, token.PAREN_CURVY_OPEN),
		dummyTok(`b`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`+`, token.CALC_ADD),
		dummyTok(`c`, token.IDENTIFIER_EXPLICIT),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
		dummyTok(`)`, token.PAREN_CURVY_CLOSE),
	}

	exp := [][]token.Token{
		[]token.Token{
			dummyTok(`#1`, token.IDENTIFIER_IMPLICIT),
			dummyTok(`<-`, token.ASSIGNMENT),
			dummyTok(`b`, token.IDENTIFIER_EXPLICIT),
			dummyTok(`+`, token.CALC_ADD),
			dummyTok(`c`, token.IDENTIFIER_EXPLICIT),
		},
		[]token.Token{
			dummyTok(`#2`, token.IDENTIFIER_IMPLICIT),
			dummyTok(`<-`, token.ASSIGNMENT),
			dummyTok(`a`, token.IDENTIFIER_EXPLICIT),
			dummyTok(`*`, token.CALC_MULTIPLY),
			dummyTok(`#1`, token.IDENTIFIER_IMPLICIT),
		},
		[]token.Token{
			dummyTok(`x`, token.IDENTIFIER_EXPLICIT),
			dummyTok(`<-`, token.ASSIGNMENT),
			dummyTok(`#2`, token.IDENTIFIER_IMPLICIT),
		},
	}

	act, err := ExpandParens(in)
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
	Input     []token.Token
	Expect    [][]token.Token
	ExpectErr AnaError
}

func expBracketsTests() []expBracketsTest {
	return []expBracketsTest{
		expBracketsTest{
			Input: []token.Token{
				token.Token{`(`, 0, 1, 0, token.PAREN_CURVY_OPEN},
				token.Token{`x`, 1, 2, 0, token.IDENTIFIER_EXPLICIT},
				token.Token{`)`, 2, 3, 0, token.PAREN_CURVY_CLOSE},
			},
			Expect: [][]token.Token{
				[]token.Token{
					token.Token{`x`, 1, 2, 0, token.IDENTIFIER_EXPLICIT},
				},
			},
		},

			expBracketsTest{
				Input: []token.Token{
					token.Token{`x`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
					token.Token{`<-`, 1, 3, 0, token.ASSIGNMENT},
					token.Token{`(`, 3, 4, 0, token.PAREN_CURVY_OPEN},
					token.Token{`y`, 4, 5, 0, token.IDENTIFIER_EXPLICIT},
					token.Token{`)`, 5, 6, 0, token.PAREN_CURVY_CLOSE},
				},
				Expect: [][]token.Token{
					[]token.Token{
						token.Token{`x`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
						token.Token{`<-`, 1, 3, 0, token.ASSIGNMENT},
						token.Token{`y`, 4, 5, 0, token.IDENTIFIER_EXPLICIT},
					},
				},
			},

	}
}
*/
