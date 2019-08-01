package parser

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/expr/exprs"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parseTest struct {
	TestLine int
	Input    []token.Token
	Expect   ctx.Expression
	Error    fault.Fault
}

type dummy struct {
	Val ctx.Value
	Err fault.Fault
}

func (d dummy) Evaluate(c *ctx.Context) (v ctx.Value, err fault.Fault) {
	return d.Val, d.Err
}

func valDummy(v ctx.Value) ctx.Expression {
	return dummy{
		Val: v,
	}
}

func errDummy(err fault.Fault) ctx.Expression {
	return dummy{
		Err: err,
	}
}

func TestParser(t *testing.T) {
	for _, tc := range makeParseTests() {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> parser_test.go : " + testLine)

		var act ctx.Expression
		var err fault.Fault
		act, err = Parse(tc.Input)

		if tc.Error != nil {
			assert.Nil(t, act)
			require.NotNil(t, err)

		} else {
			assert.Nil(t, err)
			require.NotNil(t, act)
			assert.Equal(t, tc.Expect, act)
		}

	}
}

func makeParseTests() []parseTest {
	return []parseTest{
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
			},
			Expect: exprs.Assignment{
				Identifier: token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				Operator:   token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				Right: exprs.Number{
					Number: token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
				},
			},
		},
	}
}
