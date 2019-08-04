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
						NewValue(`1`, token.LITERAL_NUMBER),
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
							Left:     NewValue(`1`, token.LITERAL_NUMBER),
							Operator: Token{`+`, 0, 0, token.CALC_ADD},
							Right:    NewValue(`2`, token.LITERAL_NUMBER),
						},
					},
				},
			},
		},
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []Token{
				// x <- 1 + 3 - 2
				Token{`x`, 0, 0, token.IDENTIFIER},
				Token{`<-`, 0, 0, token.ASSIGNMENT},
				Token{`1`, 0, 0, token.LITERAL_NUMBER},
				Token{`+`, 0, 0, token.CALC_ADD},
				Token{`3`, 0, 0, token.LITERAL_NUMBER},
				Token{`-`, 0, 0, token.CALC_SUBTRACT},
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
							Left: Operation{
								Left:     NewValue(`1`, token.LITERAL_NUMBER),
								Operator: Token{`+`, 0, 0, token.CALC_ADD},
								Right:    NewValue(`3`, token.LITERAL_NUMBER),
							},
							Operator: Token{`-`, 0, 0, token.CALC_SUBTRACT},
							Right:    NewValue(`2`, token.LITERAL_NUMBER),
						},
					},
				},
			},
		},
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []Token{
				// x <- 1 + 2 + 3 + 4
				Token{`x`, 0, 0, token.IDENTIFIER},
				Token{`<-`, 0, 0, token.ASSIGNMENT},
				Token{`1`, 0, 0, token.LITERAL_NUMBER},
				Token{`-`, 0, 0, token.CALC_SUBTRACT},
				Token{`2`, 0, 0, token.LITERAL_NUMBER},
				Token{`+`, 0, 0, token.CALC_ADD},
				Token{`3`, 0, 0, token.LITERAL_NUMBER},
				Token{`*`, 0, 0, token.CALC_MULTIPLY},
				Token{`4`, 0, 0, token.LITERAL_NUMBER},
				Token{`/`, 0, 0, token.CALC_DIVIDE},
				Token{`5`, 0, 0, token.LITERAL_NUMBER},
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
							Left: Operation{
								Left: Operation{
									Left: Operation{
										Left:     NewValue(`1`, token.LITERAL_NUMBER),
										Operator: Token{`-`, 0, 0, token.CALC_SUBTRACT},
										Right:    NewValue(`2`, token.LITERAL_NUMBER),
									},
									Operator: Token{`+`, 0, 0, token.CALC_ADD},
									Right:    NewValue(`3`, token.LITERAL_NUMBER),
								},
								Operator: Token{`*`, 0, 0, token.CALC_MULTIPLY},
								Right:    NewValue(`4`, token.LITERAL_NUMBER),
							},
							Operator: Token{`/`, 0, 0, token.CALC_DIVIDE},
							Right:    NewValue(`5`, token.LITERAL_NUMBER),
						},
					},
				},
			},
		},
	}
}
