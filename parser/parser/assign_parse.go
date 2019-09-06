package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/expr/assign"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// matchAssign returns true if the next part of the statement is an assignment.
func matchAssign(p *Parser) bool {
	for i, tk := range p.t[p.i:] {
		k := tk.Kind()

		if i%2 == 0 {
			if k != token.TT_ID && k != token.TT_VOID {
				return false
			}
		} else {
			if tk.Kind() != token.TT_DELIM {
				if k == token.TT_ASSIGN {
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
		} else if k == token.TT_ASSIGN {
			break
		}
	}

	return dst
}

// parseAssignSrc parses the source expressions, e.g. `[...<-, x+1, y*2, z]`
func parseAssignSrc(p *Parser) ([]expr.Expr, perror.Perror) {
	src := []expr.Expr{}

	for ; p.i < p.size; p.i = p.i + 1 {
		ex, e := nextAssignSrc(p)
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

// nextAssignSrc parses the source expressions that produce the values to
// assign.
func nextAssignSrc(p *Parser) (expr.Expr, perror.Perror) {
	switch {
	case matchOperand(p):
		return parseOperand(p)
	case p.t[p.i].Kind() == token.TT_NEWLINE:
		p.i++
		return nil, nil
	default:
		return nil, noMatch(p)
	}
}
