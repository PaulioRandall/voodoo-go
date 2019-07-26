package strimmer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrim(t *testing.T) {
	for i, tc := range strimTests() {
		t.Log("Strim() test case: " + strconv.Itoa(i+1))
		ts := Strim(tc.Input)
		require.NotNil(t, ts)
		assert.Equal(t, tc.ExpectToks, ts)
	}
}

type strimTest struct {
	Input      []symbol.Lexeme
	ExpectToks []symbol.Token
}

func strimTests() []strimTest {
	return []strimTest{
		strimTest{
			Input: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`1`, 5, 6, 0, symbol.NUMBER},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{`1`, 5, 6, 0, symbol.NUMBER},
			},
		},
		strimTest{
			Input: []symbol.Lexeme{
				symbol.Lexeme{`// 'There's a snake in my boot'`, 0, 31, 0, symbol.COMMENT},
			},
			ExpectToks: []symbol.Token{},
		},
		strimTest{
			Input: []symbol.Lexeme{
				symbol.Lexeme{`x`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Lexeme{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Lexeme{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Lexeme{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Lexeme{`2`, 5, 6, 0, symbol.NUMBER},
				symbol.Lexeme{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Lexeme{`// 'There's a snake in my boot'`, 7, 38, 0, symbol.COMMENT},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{`2`, 5, 6, 0, symbol.NUMBER},
			},
		},
		strimTest{
			Input: []symbol.Lexeme{
				symbol.Lexeme{`"Howdy partner"`, 5, 20, 0, symbol.STRING},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`Howdy partner`, 5, 20, 0, symbol.STRING},
			},
		},
		strimTest{
			Input: []symbol.Lexeme{
				symbol.Lexeme{`123_456`, 0, 7, 0, symbol.NUMBER},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`123456`, 0, 7, 0, symbol.NUMBER},
			},
		},
		strimTest{
			Input: []symbol.Lexeme{
				symbol.Lexeme{`1__2__3__.__4__5__6__`, 0, 21, 0, symbol.NUMBER},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`123.456`, 0, 21, 0, symbol.NUMBER},
			},
		},
		strimTest{
			Input: []symbol.Lexeme{
				symbol.Lexeme{`spell`, 0, 6, 0, symbol.KEYWORD_SPELL},
				symbol.Lexeme{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Lexeme{`END`, 7, 10, 0, symbol.KEYWORD_END},
				symbol.Lexeme{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Lexeme{`@PrInTlN`, 11, 19, 0, symbol.SOURCERY},
				symbol.Lexeme{`語`, 19, 20, 0, symbol.IDENTIFIER},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`spell`, 0, 6, 0, symbol.KEYWORD_SPELL},
				symbol.Token{`end`, 7, 10, 0, symbol.KEYWORD_END},
				symbol.Token{`@println`, 11, 19, 0, symbol.SOURCERY},
				symbol.Token{`語`, 19, 20, 0, symbol.IDENTIFIER},
			},
		},
	}
}
