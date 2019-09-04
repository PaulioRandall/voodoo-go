package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/assign"
	"github.com/PaulioRandall/voodoo-go/parser/expr/operand"
	//"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummy(t string, k token.Kind) token.Token {
	return token.New(t, 0, 0, 0, k)
}

func mock(in []token.Token) *Parser {
	return &Parser{
		size: len(in),
		t:    in,
	}
}

func doParse(t *testing.T, p *Parser, in token.Token, exp expr.Expr, err bool) {
	act, e := p.Parse(in)
	if !err && e != nil {
		defer t.Log(e.Errors()[0])
	}
	require.Equal(t, err, e != nil, `perror != nil`)
	require.Equal(t, exp, act, `exp.(Expr) != act.(Expr)`)
}

func TestParse_1(t *testing.T) {
	in := []token.Token{
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`1`, token.TT_NUMBER),
		dummy("\n", token.TT_NEWLINE),
	}

	exp := assign.New(
		in[1],
		[]expr.Expr{
			operand.New(in[2]),
		},
		[]token.Token{in[0]},
	)

	p := New()

	doParse(t, p, in[0], nil, false)
	doParse(t, p, in[1], nil, false)
	doParse(t, p, in[2], nil, false)
	doParse(t, p, in[3], exp, false)
}
