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

	eos := dummy("\n", token.TK_NEWLINE)
	doParse(t, p, eos, exp, err)
}

func TestParse_1(t *testing.T) {
	in := []token.Token{
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`1`, token.TK_NUMBER),
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
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`_`, token.TK_VOID),
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
		dummy(`_`, token.TK_VOID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`2`, token.TK_NUMBER),
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
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`false`, token.TK_BOOL),
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

func TestParse_5(t *testing.T) {
	in := []token.Token{
		dummy(`x`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`dragonfly`, token.TK_STRING),
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

func TestParse_6(t *testing.T) {
	in := []token.Token{
		dummy(`x`, token.TK_ID),
		dummy(`,`, token.TK_DELIM),
		dummy(`y`, token.TK_VOID),
		dummy(`,`, token.TK_DELIM),
		dummy(`z`, token.TK_ID),
		dummy(`<-`, token.TK_ASSIGN),
		dummy(`4`, token.TK_NUMBER),
		dummy(`,`, token.TK_DELIM),
		dummy(`Dragonfly`, token.TK_STRING),
		dummy(`,`, token.TK_DELIM),
		dummy(`_`, token.TK_VOID),
	}

	exp := assign.New(
		in[5],
		[]expr.Expr{
			operand.New(in[6]),
			operand.New(in[8]),
			operand.New(in[10]),
		},
		[]token.Token{
			in[0],
			in[2],
			in[4],
		},
	)

	doTestParseStat(t, in, exp, false)
}
