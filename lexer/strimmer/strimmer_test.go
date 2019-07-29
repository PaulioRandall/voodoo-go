package strimmer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type strimTest struct {
	TestLine   int
	Input      []token.Token
	ExpectToks []token.Token
}

func TestStrim(t *testing.T) {
	for _, tc := range strimTests() {
		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> strimmer_test.go : " + testLine)

		ts := Strim(tc.Input)
		require.NotNil(t, ts)
		assert.Equal(t, tc.ExpectToks, ts)
	}
}

func strimTests() []strimTest {
	return []strimTest{
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_IMPLICIT},
				token.Token{` `, 1, 2, 0, token.WHITESPACE},
				token.Token{`<-`, 2, 4, 0, token.ASSIGNMENT},
				token.Token{` `, 4, 5, 0, token.WHITESPACE},
				token.Token{`1`, 5, 6, 0, token.LITERAL_NUMBER},
			},
			ExpectToks: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_IMPLICIT},
				token.Token{`<-`, 2, 4, 0, token.ASSIGNMENT},
				token.Token{`1`, 5, 6, 0, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`// 'There's a snake in my boot'`, 0, 31, 0, token.COMMENT},
			},
			ExpectToks: []token.Token{},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_IMPLICIT},
				token.Token{` `, 1, 2, 0, token.WHITESPACE},
				token.Token{`<-`, 2, 4, 0, token.ASSIGNMENT},
				token.Token{` `, 4, 5, 0, token.WHITESPACE},
				token.Token{`2`, 5, 6, 0, token.LITERAL_NUMBER},
				token.Token{` `, 6, 7, 0, token.WHITESPACE},
				token.Token{`// 'There's a snake in my boot'`, 7, 38, 0, token.COMMENT},
			},
			ExpectToks: []token.Token{
				token.Token{`x`, 0, 1, 0, token.IDENTIFIER_IMPLICIT},
				token.Token{`<-`, 2, 4, 0, token.ASSIGNMENT},
				token.Token{`2`, 5, 6, 0, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`"Howdy partner"`, 5, 20, 0, token.LITERAL_STRING},
			},
			ExpectToks: []token.Token{
				token.Token{`Howdy partner`, 5, 20, 0, token.LITERAL_STRING},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`123_456`, 0, 7, 0, token.LITERAL_NUMBER},
			},
			ExpectToks: []token.Token{
				token.Token{`123456`, 0, 7, 0, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`1__2__3__.__4__5__6__`, 0, 21, 0, token.LITERAL_NUMBER},
			},
			ExpectToks: []token.Token{
				token.Token{`123.456`, 0, 21, 0, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`func`, 0, 6, 0, token.KEYWORD_FUNC},
				token.Token{` `, 6, 7, 0, token.WHITESPACE},
				token.Token{`END`, 7, 10, 0, token.KEYWORD_DONE},
				token.Token{` `, 10, 11, 0, token.WHITESPACE},
				token.Token{`@PrInTlN`, 11, 19, 0, token.SPELL},
				token.Token{`語`, 19, 20, 0, token.IDENTIFIER_IMPLICIT},
			},
			ExpectToks: []token.Token{
				token.Token{`func`, 0, 6, 0, token.KEYWORD_FUNC},
				token.Token{`end`, 7, 10, 0, token.KEYWORD_DONE},
				token.Token{`@println`, 11, 19, 0, token.SPELL},
				token.Token{`語`, 19, 20, 0, token.IDENTIFIER_IMPLICIT},
			},
		},
	}
}
