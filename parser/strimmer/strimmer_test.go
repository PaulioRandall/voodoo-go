package strimmer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type strimTest struct {
	TestLine   int
	Input      []token.Token
	ExpectToks []token.Token
}

// dummyToken creates a new dummy token.
func dummyToken(line, start, end int, val string, t token.TokenType) token.Token {
	return token.Token{
		Line:  line,
		Start: start,
		End:   end,
		Val:   val,
		Type:  t,
	}
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
				dummyToken(0, 0, 1, `x`, token.IDENTIFIER),
				dummyToken(0, 1, 2, ` `, token.WHITESPACE),
				dummyToken(0, 2, 4, `<-`, token.ASSIGNMENT),
				dummyToken(0, 4, 5, ` `, token.WHITESPACE),
				dummyToken(0, 5, 6, `1`, token.LITERAL_NUMBER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 1, `x`, token.IDENTIFIER),
				dummyToken(0, 2, 4, `<-`, token.ASSIGNMENT),
				dummyToken(0, 5, 6, `1`, token.LITERAL_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 31, `// 'There's a snake in my boot'`, token.COMMENT),
			},
			ExpectToks: []token.Token{},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 1, `x`, token.IDENTIFIER),
				dummyToken(0, 1, 2, ` `, token.WHITESPACE),
				dummyToken(0, 2, 4, `<-`, token.ASSIGNMENT),
				dummyToken(0, 4, 5, ` `, token.WHITESPACE),
				dummyToken(0, 5, 6, `2`, token.LITERAL_NUMBER),
				dummyToken(0, 6, 7, ` `, token.WHITESPACE),
				dummyToken(0, 7, 38, `// 'There's a snake in my boot'`, token.COMMENT),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 1, `x`, token.IDENTIFIER),
				dummyToken(0, 2, 4, `<-`, token.ASSIGNMENT),
				dummyToken(0, 5, 6, `2`, token.LITERAL_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 5, 20, `"Howdy partner"`, token.LITERAL_STRING),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 5, 20, `Howdy partner`, token.LITERAL_STRING),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 7, `123_456`, token.LITERAL_NUMBER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 7, `123456`, token.LITERAL_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 21, `1__2__3__.__4__5__6__`, token.LITERAL_NUMBER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 21, `123.456`, token.LITERAL_NUMBER),
			},
		},
		strimTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				dummyToken(0, 0, 6, `func`, token.KEYWORD_FUNC),
				dummyToken(0, 6, 7, ` `, token.WHITESPACE),
				dummyToken(0, 7, 10, `END`, token.KEYWORD_DONE),
				dummyToken(0, 10, 11, ` `, token.WHITESPACE),
				dummyToken(0, 11, 19, `@PrInTlN`, token.SPELL),
				dummyToken(0, 19, 20, `語`, token.IDENTIFIER),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 6, `func`, token.KEYWORD_FUNC),
				dummyToken(0, 7, 10, `end`, token.KEYWORD_DONE),
				dummyToken(0, 11, 19, `@println`, token.SPELL),
				dummyToken(0, 19, 20, `語`, token.IDENTIFIER),
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
				dummyToken(0, 0, 1, `x`, token.IDENTIFIER),
				dummyToken(0, 1, 2, ` `, token.WHITESPACE),
				dummyToken(0, 2, 4, `<-`, token.ASSIGNMENT),
				dummyToken(0, 4, 5, ` `, token.WHITESPACE),
				dummyToken(0, 5, 6, `[`, token.PAREN_SQUARE_OPEN),
				dummyToken(0, 6, 7, "\n", token.NEWLINE),
				dummyToken(0, 0, 2, `  `, token.WHITESPACE),
				dummyToken(0, 2, 3, `1`, token.LITERAL_NUMBER),
				dummyToken(0, 3, 4, "\n", token.NEWLINE),
				dummyToken(0, 0, 2, `  `, token.WHITESPACE),
				dummyToken(0, 2, 3, `2`, token.LITERAL_NUMBER),
				dummyToken(0, 3, 4, `,`, token.VALUE_DELIM),
				dummyToken(0, 4, 5, ` `, token.WHITESPACE),
				dummyToken(0, 5, 6, `3`, token.LITERAL_NUMBER),
				dummyToken(0, 6, 7, "\n", token.NEWLINE),
				dummyToken(0, 0, 1, `]`, token.PAREN_SQUARE_CLOSE),
				dummyToken(0, 1, 2, "\n", token.NEWLINE),
			},
			ExpectToks: []token.Token{
				dummyToken(0, 0, 1, `x`, token.IDENTIFIER),
				dummyToken(0, 2, 4, `<-`, token.ASSIGNMENT),
				dummyToken(0, 5, 6, `[`, token.PAREN_SQUARE_OPEN),
				dummyToken(0, 2, 3, `1`, token.LITERAL_NUMBER),
				dummyToken(0, 3, 4, "\n", token.TT_EOS),
				dummyToken(0, 2, 3, `2`, token.LITERAL_NUMBER),
				dummyToken(0, 3, 4, `,`, token.VALUE_DELIM),
				dummyToken(0, 5, 6, `3`, token.LITERAL_NUMBER),
				dummyToken(0, 6, 7, "\n", token.TT_EOS),
				dummyToken(0, 0, 1, `]`, token.PAREN_SQUARE_CLOSE),
				dummyToken(0, 1, 2, "\n", token.TT_EOS),
			},
		},
	}
}
