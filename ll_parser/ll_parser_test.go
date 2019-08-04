package ll_parser

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parseTest struct {
	TestLine int
	Input    []token.Token
	Stat     ctx.Statement
	Error    fault.Fault
}

func TestParse(t *testing.T) {
	for _, tc := range makeParseTests() {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> parser_test.go : " + testLine)

		var stat ctx.Statement
		var err fault.Fault
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

func NewValue(vStr string, v interface{}, tkType token.TokenType, valType ctx.ValueType) ctx.Value {
	return ctx.Value{
		Token: token.Token{
			Val:  vStr,
			Type: tkType,
		},
		Val:  v,
		Type: valType,
	}
}

func makeParseTests() []parseTest {
	return []parseTest{
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				// x <- 1
				token.Token{`x`, 0, 0, token.IDENTIFIER},
				token.Token{`<-`, 0, 0, token.ASSIGNMENT},
				token.Token{`1`, 0, 0, token.LITERAL_NUMBER},
			},
			Stat: ctx.Assignment{
				Left:     NewValue(`x`, `x`, token.IDENTIFIER, ctx.IDENTIFIER_TYPE),
				Operator: token.Token{`<-`, 0, 0, token.ASSIGNMENT},
				Right:    NewValue(`1`, float64(1), token.LITERAL_NUMBER, ctx.NUMBER_TYPE),
			},
		},
		/*
			parseTest{
				TestLine: fault.CurrLine(),
				Input: []token.Token{
					// x <- 1 + 2
					token.Token{`x`, 0, 0, token.IDENTIFIER},
					token.Token{`<-`, 0, 0, token.ASSIGNMENT},
					token.Token{`1`, 0, 0, token.LITERAL_NUMBER},
					token.Token{`+`, 0, 0, token.CALC_ADD},
					token.Token{`2`, 0, 0, token.LITERAL_NUMBER},
				},
				Stat: exprs.Assignment{
					Left: exprs.List{
						Tokens: []token.Token{
							token.Token{`x`, 0, 0, token.IDENTIFIER},
						},
					},
					Operator: token.Token{`<-`, 0, 0, token.ASSIGNMENT},
					Right: exprs.Join{
						Exprs: []ctx.Expression{
							exprs.Operation{
								Left:     NewValue(`1`, token.LITERAL_NUMBER),
								Operator: token.Token{`+`, 0, 0, token.CALC_ADD},
								Right:    NewValue(`2`, token.LITERAL_NUMBER),
							},
						},
					},
				},
			},
			parseTest{
				TestLine: fault.CurrLine(),
				Input: []token.Token{
					// x <- 1 + 3 - 2
					token.Token{`x`, 0, 0, token.IDENTIFIER},
					token.Token{`<-`, 0, 0, token.ASSIGNMENT},
					token.Token{`1`, 0, 0, token.LITERAL_NUMBER},
					token.Token{`+`, 0, 0, token.CALC_ADD},
					token.Token{`3`, 0, 0, token.LITERAL_NUMBER},
					token.Token{`-`, 0, 0, token.CALC_SUBTRACT},
					token.Token{`2`, 0, 0, token.LITERAL_NUMBER},
				},
				Stat: exprs.Assignment{
					Left: exprs.List{
						Tokens: []token.Token{
							token.Token{`x`, 0, 0, token.IDENTIFIER},
						},
					},
					Operator: token.Token{`<-`, 0, 0, token.ASSIGNMENT},
					Right: exprs.Join{
						Exprs: []ctx.Expression{
							exprs.Operation{
								Left: Operation{
									Left:     NewValue(`1`, token.LITERAL_NUMBER),
									Operator: token.Token{`+`, 0, 0, token.CALC_ADD},
									Right:    NewValue(`3`, token.LITERAL_NUMBER),
								},
								Operator: token.Token{`-`, 0, 0, token.CALC_SUBTRACT},
								Right:    NewValue(`2`, token.LITERAL_NUMBER),
							},
						},
					},
				},
			},
			parseTest{
				TestLine: fault.CurrLine(),
				Input: []token.Token{
					// x <- 1 + 2 + 3 + 4
					token.Token{`x`, 0, 0, token.IDENTIFIER},
					token.Token{`<-`, 0, 0, token.ASSIGNMENT},
					token.Token{`1`, 0, 0, token.LITERAL_NUMBER},
					token.Token{`-`, 0, 0, token.CALC_SUBTRACT},
					token.Token{`2`, 0, 0, token.LITERAL_NUMBER},
					token.Token{`+`, 0, 0, token.CALC_ADD},
					token.Token{`3`, 0, 0, token.LITERAL_NUMBER},
					token.Token{`*`, 0, 0, token.CALC_MULTIPLY},
					token.Token{`4`, 0, 0, token.LITERAL_NUMBER},
					token.Token{`/`, 0, 0, token.CALC_DIVIDE},
					token.Token{`5`, 0, 0, token.LITERAL_NUMBER},
				},
				Stat: exprs.Assignment{
					Left: List{
						Tokens: []token.Token{
							token.Token{`x`, 0, 0, token.IDENTIFIER},
						},
					},
					Operator: token.Token{`<-`, 0, 0, token.ASSIGNMENT},
					Right: exprs.Join{
						Exprs: []ctx.Expression{
							exprs.Operation{
								Left: exprs.Operation{
									Left: exprs.Operation{
										Left: exprs.Operation{
											Left:     NewValue(`1`, token.LITERAL_NUMBER),
											Operator: token.Token{`-`, 0, 0, token.CALC_SUBTRACT},
											Right:    NewValue(`2`, token.LITERAL_NUMBER),
										},
										Operator: token.Token{`+`, 0, 0, token.CALC_ADD},
										Right:    NewValue(`3`, token.LITERAL_NUMBER),
									},
									Operator: token.Token{`*`, 0, 0, token.CALC_MULTIPLY},
									Right:    NewValue(`4`, token.LITERAL_NUMBER),
								},
								Operator: token.Token{`/`, 0, 0, token.CALC_DIVIDE},
								Right:    NewValue(`5`, token.LITERAL_NUMBER),
							},
						},
					},
				},
			},
		*/
	}
}
