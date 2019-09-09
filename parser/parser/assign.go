package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/assign"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// matchAssign returns true if the next part of the statement is an assignment.
func matchAssign(p *Parser, ip int) bool {
	for i, tk := range p.t[ip:] {
		k := tk.Kind()

		if i%2 == 0 {
			if k != token.TK_ID && k != token.TK_VOID {
				return false
			}
		} else {
			if tk.Kind() != token.TK_DELIM {
				if k == token.TK_ASSIGN {
					return i+1 < p.size
				}
				return false
			}
		}
	}

	return false
}

// parseAssign parses an assignment expression.
func parseAssign(p *Parser) (expr.Expr, perror.Perror) {
	dst := parseAssignDst(p)
	t := p.t[p.i]
	p.i++

	src, e := parseAssignSrc(p)
	if e != nil {
		return nil, e
	}

	return assign.New(t, src, dst), nil
}

// parseAssignDst parses the destination IDs, e.g. `[x, y, z, <-, ...ignore]`
func parseAssignDst(p *Parser) []token.Token {
	dst := []token.Token{}

	for even := true; p.i < p.size; even, p.i = !even, p.i+1 {
		k := p.t[p.i].Kind()
		if even {
			dst = append(dst, p.t[p.i])
		} else if k == token.TK_ASSIGN {
			break
		}
	}

	return dst
}

// parseAssignSrc parses the source expressions, e.g. `[...<-, x+1, y*2, z]`
func parseAssignSrc(p *Parser) ([]expr.Expr, perror.Perror) {
	src := []expr.Expr{}

	// TODO: This needs to be changed so it checks that a comma appears
	// TODO: between each expression.
	for ; p.i < p.size; p.i++ {
		ex, e := parseNextSrc(p)
		if e != nil {
			return nil, e
		}
		if ex == nil {
			break
		}
		src = append(src, ex)
	}
	return src, nil
}

// parseNextSrc parses a source expression whos result will be assigned to a
// variable.
func parseNextSrc(p *Parser) (expr.Expr, perror.Perror) {
	switch {
	case matchNumber(p, p.i):
		return parseNumber(p)
	case matchBool(p, p.i):
		return parseBool(p)
	case matchOperand(p, p.i):
		return parseOperandOnly(p)
	case p.t[p.i].Kind() == token.TK_NEWLINE:
		p.i++
		return nil, nil
	default:
		return nil, noMatch(p)
	}
}
