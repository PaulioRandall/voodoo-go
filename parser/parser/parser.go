package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

const max_tokens = 256

// Parser represents a structure used to parse tokens into expression trees.
type Parser struct {
	i    int
	size int
	t    []token.Token
}

// New returns a new Parser.
func New() *Parser {
	return &Parser{
		t: make([]token.Token, max_tokens),
	}
}

// Parse parses the next token by storing it and based on the token type,
// returning an expression tree. If nil is returned then more tokens are
// required before parsing can begin.
func (p *Parser) Parse(tk token.Token) (expr.Expr, perror.Perror) {
	if !analyse(p, tk) {
		return nil, nil
	}

	p.size, p.i = p.i, 0
	return parse(p)
}

// analyse stores the token then checks to see if parsing can begin returning
// true if so.
func analyse(p *Parser, tk token.Token) bool {
	p.t[p.i] = tk
	p.i++
	return tk.Kind() == token.TT_NEWLINE
}

// parse kicks off the parsing process.
func parse(p *Parser) (expr.Expr, perror.Perror) {
	switch {
	case matchAssign(p):
		return parseAssign(p)
	default:
		return nil, noMatch(p)
	}
}

// noMatch creates a new Perror that details a failure to match the tokens to a
// valid token pattern.
func noMatch(p *Parser) perror.Perror {
	return perror.New(
		p.t[0].Line(),
		0,
		[]string{
			`No matching pattern found for this statement`,
		},
	)
}
