package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/operand"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// matchOperand returns true if the next part of the statement is an operand.
func matchOperand(p *Parser, ip int) bool {
	k := p.t[ip].Kind()
	return k == token.TK_ID ||
		k == token.TK_STRING ||
		k == token.TK_VOID
}

// parseOperandOnly parses an operand expression but does not atempt to match
// the expression afterwards.
func parseOperandOnly(p *Parser) (expr.Expr, perror.Perror) {
	o := operand.New(p.t[p.i])
	p.i++
	return o, nil
}
