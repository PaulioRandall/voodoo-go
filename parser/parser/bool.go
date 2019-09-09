package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/operand"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// matchBool returns true if the next part of the statement is a bool operand.
func matchBool(p *Parser, ip int) bool {
	k := p.t[ip].Kind()
	return k == token.TK_ID ||
		k == token.TK_BOOL
}

// parseBool parses a boolean expression then attempts to match and parse the
// remaineder of the statement knowing what valid tokens must come after a bool.
func parseBool(p *Parser) (expr.Expr, perror.Perror) {
	o := operand.New(p.t[p.i])
	p.i++
	return o, nil
}
