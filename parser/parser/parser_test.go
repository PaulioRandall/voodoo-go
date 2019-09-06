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

func doTestParseStat(t *testing.T, in []token.Token, exp expr.Expr, err bool) {
	p := New()

	for _, v := range in {
		doParse(t, p, v, nil, false)
	}

	eos := dummy("\n", token.TT_NEWLINE)
	doParse(t, p, eos, exp, err)
}

func TestParse_1(t *testing.T) {
	in := []token.Token{
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`1`, token.TT_NUMBER),
	}

	exp := assign.New(
		in[1],
		[]expr.Expr{
			operand.New(in[2]),
		},
		[]token.Token{in[0]},
	)

	doTestParseStat(t, in, exp, false)
}

func TestParse_2(t *testing.T) {
	in := []token.Token{
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`_`, token.TT_VOID),
	}

	exp := assign.New(
		in[1],
		[]expr.Expr{
			operand.New(in[2]),
		},
		[]token.Token{in[0]},
	)

	doTestParseStat(t, in, exp, false)
}

func TestParse_3(t *testing.T) {
	in := []token.Token{
		dummy(`_`, token.TT_VOID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`2`, token.TT_NUMBER),
	}

	exp := assign.New(
		in[1],
		[]expr.Expr{
			operand.New(in[2]),
		},
		[]token.Token{in[0]},
	)

	doTestParseStat(t, in, exp, false)
}

func TestParse_4(t *testing.T) {
	in := []token.Token{
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`false`, token.TT_BOOL),
	}

	exp := assign.New(
		in[1],
		[]expr.Expr{
			operand.New(in[2]),
		},
		[]token.Token{in[0]},
	)

	doTestParseStat(t, in, exp, false)
}
