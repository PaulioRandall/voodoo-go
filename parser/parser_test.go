package parser

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parseTest struct {
	TestLine int
	Input    []Token
	Exes     []Exe
	Values   []Token
	Error    Fault
}

func newExe(p, r int, tk Token) Exe {
	return Exe{
		Token:   tk,
		Params:  p,
		Returns: r,
	}
}

func TestParser(t *testing.T) {
	for _, tc := range makeParseTests() {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> parser_test.go : " + testLine)

		var exes []Exe
		var values []Token
		var err Fault
		exes, values, err = Parse(tc.Input)

		if tc.Error != nil {
			assert.Nil(t, exes)
			assert.Nil(t, values)
			require.NotNil(t, err)

		} else {
			assert.Nil(t, err)

			require.NotNil(t, exes)
			require.NotNil(t, values)

			assert.Equal(t, tc.Exes, exes)
			assert.Equal(t, tc.Values, values)
		}
	}
}

func makeParseTests() []parseTest {
	return []parseTest{
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []Token{
				Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				Token{`<-`, 2, 4, token.ASSIGNMENT},
				Token{`1`, 5, 6, token.LITERAL_NUMBER},
			},
			Exes: []Exe{
				newExe(2, 1, Token{`<-`, 2, 4, token.ASSIGNMENT}),
			},
			Values: []Token{
				Token{`1`, 5, 6, token.LITERAL_NUMBER},
				Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
			},
		},
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []Token{
				Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				Token{`<-`, 2, 4, token.ASSIGNMENT},
				Token{`1`, 5, 6, token.LITERAL_NUMBER},
				Token{`+`, 7, 8, token.CALC_ADD},
				Token{`2`, 9, 10, token.LITERAL_NUMBER},
			},
			Exes: []Exe{
				newExe(2, 1, Token{`+`, 7, 8, token.CALC_ADD}),
				newExe(2, 1, Token{`<-`, 2, 4, token.ASSIGNMENT}),
			},
			Values: []Token{
				Token{`1`, 5, 6, token.LITERAL_NUMBER},
				Token{`2`, 9, 10, token.LITERAL_NUMBER},
				Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
			},
		},
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []Token{
				Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				Token{`<-`, 2, 4, token.ASSIGNMENT},
				Token{`1`, 5, 6, token.LITERAL_NUMBER},
				Token{`+`, 7, 8, token.CALC_ADD},
				Token{`3`, 9, 10, token.LITERAL_NUMBER},
				Token{`-`, 11, 12, token.CALC_SUBTRACT},
				Token{`2`, 13, 14, token.LITERAL_NUMBER},
			},
			Exes: []Exe{
				newExe(2, 1, Token{`+`, 7, 8, token.CALC_ADD}),
				newExe(2, 1, Token{`-`, 11, 12, token.CALC_SUBTRACT}),
				newExe(2, 1, Token{`<-`, 2, 4, token.ASSIGNMENT}),
			},
			Values: []Token{
				Token{`1`, 5, 6, token.LITERAL_NUMBER},
				Token{`3`, 9, 10, token.LITERAL_NUMBER},
				Token{`2`, 13, 14, token.LITERAL_NUMBER},
				Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
			},
		},
	}
}
