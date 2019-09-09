package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/operand"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// matchNumber returns true if the next part of the statement is a number
// or ID operand.
func matchNumber(p *Parser, ip int) bool {
	k := p.t[ip].Kind()
	return k == token.TK_ID ||
		k == token.TK_NUMBER
}

// parseNumber parses an number expression then attempts match and parse the
// remaining statement knowing what valid tokens appear after a number.
func parseNumber(p *Parser) (expr.Expr, perror.Perror) {
	o := operand.New(p.t[p.i])
	p.i++
	return afterNumber(p, o)
}

// afterNumber parses the expression that follows a number.
func afterNumber(p *Parser, nu expr.Expr) (expr.Expr, perror.Perror) {
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
