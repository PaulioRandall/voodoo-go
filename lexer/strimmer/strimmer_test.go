package strimmer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrim(t *testing.T) {
	for _, tc := range strimTests() {
		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> strimmer_test.go : " + testLine)

		ts := Strim(tc.Input)
		require.NotNil(t, ts)
		assert.Equal(t, tc.ExpectToks, ts)
	}
}

type strimTest struct {
	TestLine   int
	Input      []symbol.Token
	ExpectToks []symbol.Token
}

func strimTests() []strimTest {
	return []strimTest{
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_IMPLICIT},
				symbol.Token{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Token{`1`, 5, 6, 0, symbol.LITERAL_NUMBER},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_IMPLICIT},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{`1`, 5, 6, 0, symbol.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []symbol.Token{
				symbol.Token{`// 'There's a snake in my boot'`, 0, 31, 0, symbol.COMMENT},
			},
			ExpectToks: []symbol.Token{},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_IMPLICIT},
				symbol.Token{` `, 1, 2, 0, symbol.WHITESPACE},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{` `, 4, 5, 0, symbol.WHITESPACE},
				symbol.Token{`2`, 5, 6, 0, symbol.LITERAL_NUMBER},
				symbol.Token{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Token{`// 'There's a snake in my boot'`, 7, 38, 0, symbol.COMMENT},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`x`, 0, 1, 0, symbol.IDENTIFIER_IMPLICIT},
				symbol.Token{`<-`, 2, 4, 0, symbol.ASSIGNMENT},
				symbol.Token{`2`, 5, 6, 0, symbol.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []symbol.Token{
				symbol.Token{`"Howdy partner"`, 5, 20, 0, symbol.LITERAL_STRING},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`Howdy partner`, 5, 20, 0, symbol.LITERAL_STRING},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []symbol.Token{
				symbol.Token{`123_456`, 0, 7, 0, symbol.LITERAL_NUMBER},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`123456`, 0, 7, 0, symbol.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []symbol.Token{
				symbol.Token{`1__2__3__.__4__5__6__`, 0, 21, 0, symbol.LITERAL_NUMBER},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`123.456`, 0, 21, 0, symbol.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []symbol.Token{
				symbol.Token{`func`, 0, 6, 0, symbol.KEYWORD_FUNC},
				symbol.Token{` `, 6, 7, 0, symbol.WHITESPACE},
				symbol.Token{`END`, 7, 10, 0, symbol.KEYWORD_END},
				symbol.Token{` `, 10, 11, 0, symbol.WHITESPACE},
				symbol.Token{`@PrInTlN`, 11, 19, 0, symbol.SOURCERY},
				symbol.Token{`語`, 19, 20, 0, symbol.IDENTIFIER_IMPLICIT},
			},
			ExpectToks: []symbol.Token{
				symbol.Token{`func`, 0, 6, 0, symbol.KEYWORD_FUNC},
				symbol.Token{`end`, 7, 10, 0, symbol.KEYWORD_END},
				symbol.Token{`@println`, 11, 19, 0, symbol.SOURCERY},
				symbol.Token{`語`, 19, 20, 0, symbol.IDENTIFIER_IMPLICIT},
			},
		},
	}
}
