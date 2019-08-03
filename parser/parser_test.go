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
	Stat     Statement
	Error    Fault
}

func TestParser(t *testing.T) {
	for _, tc := range makeParseTests() {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> parser_test.go : " + testLine)

		var stat Statement
		var err Fault
		stat, err = Parse(tc.Input)

		if tc.Error != nil {
			assert.Nil(t, stat)
			require.NotNil(t, err)

		} else {
			assert.Nil(t, err)
			require.NotNil(t, stat)
			assert.Equal(t, tc.Stat, stat)
		}
	}
}

func makeParseTests() []parseTest {
	return []parseTest{
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []Token{
				// x <- 1
				Token{`x`, 0, 0, token.IDENTIFIER},
				Token{`<-`, 0, 0, token.ASSIGNMENT},
				Token{`1`, 0, 0, token.LITERAL_NUMBER},
			},
			Stat: Assignment{
				Left: List{
					Tokens: []Token{
						Token{`x`, 0, 0, token.IDENTIFIER},
					},
				},
				Operator: Token{`<-`, 0, 0, token.ASSIGNMENT},
				Right: Join{
					Exprs: []Expression{
						Value{
							Token{`1`, 0, 0, token.LITERAL_NUMBER},
						},
					},
				},
			},
		},
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []Token{
				// x <- 1 + 2
				Token{`x`, 0, 0, token.IDENTIFIER},
				Token{`<-`, 0, 0, token.ASSIGNMENT},
				Token{`1`, 0, 0, token.LITERAL_NUMBER},
				Token{`+`, 0, 0, token.CALC_ADD},
				Token{`2`, 0, 0, token.LITERAL_NUMBER},
			},
			Stat: Assignment{
				Left: List{
					Tokens: []Token{
						Token{`x`, 0, 0, token.IDENTIFIER},
					},
				},
				Operator: Token{`<-`, 0, 0, token.ASSIGNMENT},
				Right: Join{
					Exprs: []Expression{
						Operation{
							Left: Value{
								Token{`1`, 0, 0, token.LITERAL_NUMBER},
							},
							Operator: Token{`+`, 0, 0, token.CALC_ADD},
							Right: Value{
								Token{`2`, 0, 0, token.LITERAL_NUMBER},
							},
						},
					},
				},
			},
		},
		/*
			parseTest{
				TestLine: fault.CurrLine(),
				Input: []Token{
					// x <- 1 + 3 - 2
					Token{`x`, 0, 1, token.IDENTIFIER},
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
					newExe(2, 0, Token{`<-`, 2, 4, token.ASSIGNMENT}),
				},
				Vals: []Token{
					Token{`1`, 5, 6, token.LITERAL_NUMBER},
					Token{`3`, 9, 10, token.LITERAL_NUMBER},
					Token{`2`, 13, 14, token.LITERAL_NUMBER},
					Token{`x`, 0, 1, token.IDENTIFIER},
				},
			},
			parseTest{
				TestLine: fault.CurrLine(),
				Input: []Token{
					// x, y <- 1, 2
					Token{`x`, 0, 1, token.IDENTIFIER},
					Token{`,`, 1, 2, token.SEPARATOR_VALUE},
					Token{`y`, 3, 4, token.IDENTIFIER},
					Token{`<-`, 5, 7, token.ASSIGNMENT},
					Token{`1`, 8, 9, token.LITERAL_NUMBER},
					Token{`,`, 9, 10, token.SEPARATOR_VALUE},
					Token{`2`, 11, 12, token.LITERAL_NUMBER},
				},
				Exes: []Exe{
					newExe(4, 0, Token{`<-`, 5, 7, token.ASSIGNMENT}),
				},
				Vals: []Token{
					Token{`1`, 8, 9, token.LITERAL_NUMBER},
					Token{`2`, 11, 12, token.LITERAL_NUMBER},
					Token{`x`, 0, 1, token.IDENTIFIER},
					Token{`y`, 3, 4, token.IDENTIFIER},
				},
			},
		*/
	}
}
