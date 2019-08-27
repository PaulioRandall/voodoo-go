package parser_2

import (
	"errors"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// Parse parses the input statement into a parse tree.
func Parse(in []token.Token) (*tree.Tree, error) {
	return parse(nil, in)
}

// parse parses the input statement into a parse tree.
func parse(parent *tree.Tree, in []token.Token) (*tree.Tree, error) {

	if i := indexOf(in, token.TT_ASSIGN); i != -1 {
		return parseAssign(nil, in, i)
	}

	return nil, errors.New("Unrecognised statement")
}
