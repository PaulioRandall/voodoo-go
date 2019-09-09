package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// max_tokens is the maximum number of tokens in the parsers statement buffer.
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

// reset resets the parsers index so a new line statement can be parsed.
func (p *Parser) reset() {
	p.i = 0
}

// Parse parses the next token by storing it and based on the token type,
// returning an expression tree. If nil is returned then more tokens are
// required before parsing can begin.
func (p *Parser) Parse(tk token.Token) (expr.Expr, perror.Perror) {
	if tk == nil {
		return nil, perror.New(
			p.t[0].Line(),
			p.t[0].Start(),
			[]string{
				`Nil passed to parser`,
			},
		)
	}

	if !analyse(p, tk) {
		return nil, nil
	}

	p.size, p.i = p.i, 0
	defer p.reset()
	return parse(p)
}

// analyse stores the token then checks to see if parsing can begin returning
// true if so.
func analyse(p *Parser, tk token.Token) bool {
	k := tk.Kind()

	if p.i == 0 && k == token.TK_NEWLINE {
		return false
	}

	p.t[p.i] = tk
	p.i++
	return k == token.TK_NEWLINE
}

// parse kicks off the parsing process.
func parse(p *Parser) (expr.Expr, perror.Perror) {
	switch {
	case matchAssign(p, p.i):
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
		p.t[0].Start(),
		[]string{
			`Statement has no matching pattern,`,
			`I don't know how to parse it ¯\_(--)_/¯`,
		},
	)
}
