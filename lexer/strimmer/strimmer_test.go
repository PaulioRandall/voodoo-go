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

// stream takes an array of tokens and pushes them into a token channel.
func stream(in []token.Token, out chan token.Token) {
	defer close(out)
	for _, tk := range in {
		out <- tk
	}
}

// collect collects all tokens coming from a channel into an array.
func collect(in chan token.Token, out chan []token.Token) {
	defer close(out)
	tks := []token.Token{}
	for tk := range in {
		tks = append(tks, tk)
	}
	out <- tks
}

func TestStrim(t *testing.T) {
	for _, tc := range strimTests() {
		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> strimmer_test.go : " + testLine)

		inChan := make(chan token.Token)
		outChan := make(chan token.Token)
		collectChan := make(chan []token.Token)

		go stream(tc.Input, inChan)
		go collect(outChan, collectChan)

		Strim(inChan, outChan)
		tks := <-collectChan

		require.NotNil(t, tks)
		assert.Equal(t, tc.ExpectToks, tks)
	}
}

func strimTests() []strimTest {
	return []strimTest{
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER},
				token.Token{` `, 1, 2, token.WHITESPACE},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
			},
			ExpectToks: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`// 'There's a snake in my boot'`, 0, 31, token.COMMENT},
			},
			ExpectToks: []token.Token{},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER},
				token.Token{` `, 1, 2, token.WHITESPACE},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`2`, 5, 6, token.LITERAL_NUMBER},
				token.Token{` `, 6, 7, token.WHITESPACE},
				token.Token{`// 'There's a snake in my boot'`, 7, 38, token.COMMENT},
			},
			ExpectToks: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{`2`, 5, 6, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`"Howdy partner"`, 5, 20, token.LITERAL_STRING},
			},
			ExpectToks: []token.Token{
				token.Token{`Howdy partner`, 5, 20, token.LITERAL_STRING},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`123_456`, 0, 7, token.LITERAL_NUMBER},
			},
			ExpectToks: []token.Token{
				token.Token{`123456`, 0, 7, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`1__2__3__.__4__5__6__`, 0, 21, token.LITERAL_NUMBER},
			},
			ExpectToks: []token.Token{
				token.Token{`123.456`, 0, 21, token.LITERAL_NUMBER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`func`, 0, 6, token.KEYWORD_FUNC},
				token.Token{` `, 6, 7, token.WHITESPACE},
				token.Token{`END`, 7, 10, token.KEYWORD_DONE},
				token.Token{` `, 10, 11, token.WHITESPACE},
				token.Token{`@PrInTlN`, 11, 19, token.SPELL},
				token.Token{`語`, 19, 20, token.IDENTIFIER},
			},
			ExpectToks: []token.Token{
				token.Token{`func`, 0, 6, token.KEYWORD_FUNC},
				token.Token{`end`, 7, 10, token.KEYWORD_DONE},
				token.Token{`@println`, 11, 19, token.SPELL},
				token.Token{`語`, 19, 20, token.IDENTIFIER},
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				// x <- [
				//   1,
				//   2, 3,
				// ]
				//
				token.Token{`x`, 0, 1, token.IDENTIFIER},
				token.Token{` `, 1, 2, token.WHITESPACE},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`[`, 5, 6, token.PAREN_SQUARE_OPEN},
				token.Token{"\n", 6, 7, token.NEWLINE},
				token.Token{`  `, 0, 2, token.WHITESPACE},
				token.Token{`1`, 2, 3, token.LITERAL_NUMBER},
				token.Token{"\n", 3, 4, token.NEWLINE},
				token.Token{`  `, 0, 2, token.WHITESPACE},
				token.Token{`2`, 2, 3, token.LITERAL_NUMBER},
				token.Token{`,`, 3, 4, token.VALUE_DELIM},
				token.Token{` `, 4, 5, token.WHITESPACE},
				token.Token{`3`, 5, 6, token.LITERAL_NUMBER},
				token.Token{"\n", 6, 7, token.NEWLINE},
				token.Token{`]`, 0, 1, token.PAREN_SQUARE_CLOSE},
				token.Token{"\n", 1, 2, token.NEWLINE},
			},
			ExpectToks: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{`[`, 5, 6, token.PAREN_SQUARE_OPEN},
				token.Token{`1`, 2, 3, token.LITERAL_NUMBER},
				token.Token{"\n", 3, 4, token.END_OF_STATEMENT},
				token.Token{`2`, 2, 3, token.LITERAL_NUMBER},
				token.Token{`,`, 3, 4, token.VALUE_DELIM},
				token.Token{`3`, 5, 6, token.LITERAL_NUMBER},
				token.Token{"\n", 6, 7, token.END_OF_STATEMENT},
				token.Token{`]`, 0, 1, token.PAREN_SQUARE_CLOSE},
				token.Token{"\n", 1, 2, token.END_OF_STATEMENT},
			},
		},
	}
}
