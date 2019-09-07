package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/operand"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// matchOperand returns true if the next part of the statement is an operand.
func matchOperand(p *Parser) bool {
	k := p.t[p.i].Kind()
	return k == token.TK_ID ||
		k == token.TK_BOOL ||
		k == token.TK_NUMBER ||
		k == token.TK_STRING ||
		k == token.TK_VOID
}

// parseOperand parses an operand expression.
func parseOperand(p *Parser) (expr.Expr, perror.Perror) {
	o := operand.New(p.t[p.i])
	p.i++
	return o, nil
}
