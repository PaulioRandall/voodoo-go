package parser

import (
	"errors"
	"strconv"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// DONE: Make a simple parser for `x <- 1`
// NEXT: Modify the parser to handle `x, y <- 1, 2`

// Parse parses the input statement into a parse tree.
func Parse(in []token.Token) (*tree.Tree, error) {
	tr := tree.New()

	for i, tk := range in {
		ok := parseToken(tr, tk)
		if !ok {
			m := "Token[" + strconv.Itoa(i) + "] does not match any parsing rules"
			return nil, errors.New(m)
		}
	}

	return tr, nil
}

// parseToken applies the first matching parse rule --with token as subject--
// to the tree.
func parseToken(tr *tree.Tree, tk token.Token) bool {
	switch findRule(tr, tk) {
	case 1:
		tr.SetLeft(tk, tree.KD_ID)
	case 2:
		tr.Set(tk, tree.KD_UNION)
	case 3:
		tr.Set(tk, tree.KD_ASSIGN)
	case 4:
		tr.SetRight(tk, tree.KD_OPERAND)
	default:
		return false
	}

	return true
}
