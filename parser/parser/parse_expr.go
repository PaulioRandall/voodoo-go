package parser

import (
	//"errors"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// parseExprs parses a value delimitered set of expressions into a parse tree.
func parseExprs(parent *tree.Tree, in []token.Token) (*tree.Tree, error) {

	// NOTE: The function that splits the expressions will have to beaware
	// NOTE: of list and function declarations as well as spell and function
	// NOTE: calls.
	return nil, nil
}
