package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/arithmetic"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// matchArithmetic returns true if the next part of the statement is some
// arithmetic.
func matchArithmetic(p *Parser, ip int) bool {
	if ip+1 >= p.size || !matchNumber(p, ip+1) {
		return false
	}

	k := p.t[ip].Kind()
	if k == token.TK_ADD ||
		k == token.TK_SUBTRACT ||
		k == token.TK_MULTIPLY ||
		k == token.TK_DIVIDE ||
		k == token.TK_MODULO {
		return true
	}

	return false
}

// parseArithmetic parses an arithmetic operation where the numerator
// expression is preparsed.
func parseArithmetic(p *Parser, nu expr.Expr) (expr.Expr, perror.Perror) {
	t := p.t[p.i]
	p.i++

	de, e := parseOperandOnly(p)
	if e != nil {
		return nil, e
	}

	a := arithmetic.New(t, nu, de)
	return afterArithmetic(p, a)
}

// afterArithmetic parses the expression that follows some arithmetic.
func afterArithmetic(p *Parser, nu expr.Expr) (expr.Expr, perror.Perror) {
	k := p.t[p.i].Kind()
	switch {
	case k == token.TK_DELIM:
		return nu, nil
	case matchArithmetic(p, p.i):
		return parseArithmetic(p, nu)
	case k == token.TK_NEWLINE:
		p.i++
		return nu, nil
	default:
		return nil, noMatch(p)
	}
}
