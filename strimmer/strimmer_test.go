package strimmer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
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
	Input      []lexeme.Lexeme
	ExpectToks []lexeme.Token
}

func strimTests() []strimTest {
	return []strimTest{
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				lexeme.Lexeme{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`1`, 5, 6, 0, lexeme.NUMBER},
			},
			ExpectToks: []lexeme.Token{
				lexeme.Token{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				lexeme.Token{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Token{`1`, 5, 6, 0, lexeme.NUMBER},
			},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`// 'There's a snake in my boot'`, 0, 31, 0, lexeme.COMMENT},
			},
			ExpectToks: []lexeme.Token{},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				lexeme.Lexeme{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`2`, 5, 6, 0, lexeme.NUMBER},
				lexeme.Lexeme{` `, 6, 7, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`// 'There's a snake in my boot'`, 7, 38, 0, lexeme.COMMENT},
			},
			ExpectToks: []lexeme.Token{
				lexeme.Token{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				lexeme.Token{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Token{`2`, 5, 6, 0, lexeme.NUMBER},
			},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`"Howdy partner"`, 5, 20, 0, lexeme.STRING},
			},
			ExpectToks: []lexeme.Token{
				lexeme.Token{`Howdy partner`, 5, 20, 0, lexeme.STRING},
			},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`123_456`, 0, 7, 0, lexeme.NUMBER},
			},
			ExpectToks: []lexeme.Token{
				lexeme.Token{`123456`, 0, 7, 0, lexeme.NUMBER},
			},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`1__2__3__.__4__5__6__`, 0, 21, 0, lexeme.NUMBER},
			},
			ExpectToks: []lexeme.Token{
				lexeme.Token{`123.456`, 0, 21, 0, lexeme.NUMBER},
			},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`scroll`, 0, 6, 0, lexeme.KEYWORD_SCROLL},
				lexeme.Lexeme{` `, 6, 7, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`END`, 7, 10, 0, lexeme.KEYWORD_END},
				lexeme.Lexeme{` `, 10, 11, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`@PrInTlN`, 11, 19, 0, lexeme.SOURCERY},
				lexeme.Lexeme{`語`, 19, 20, 0, lexeme.IDENTIFIER},
			},
			ExpectToks: []lexeme.Token{
				lexeme.Token{`scroll`, 0, 6, 0, lexeme.KEYWORD_SCROLL},
				lexeme.Token{`end`, 7, 10, 0, lexeme.KEYWORD_END},
				lexeme.Token{`@println`, 11, 19, 0, lexeme.SOURCERY},
				lexeme.Token{`語`, 19, 20, 0, lexeme.IDENTIFIER},
			},
		},
	}
}
