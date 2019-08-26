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
		switch {
		case tr.IsLeft(tree.KD_UNDEFINED) && tk.Type == token.TT_ID:
			tr.SetLeft(tk, tree.KD_ID)
		case tr.IsLeft(tree.KD_ID) && tk.Type == token.TT_ASSIGN:
			tr.Set(tk, tree.KD_ASSIGN)
		case tr.Are(tree.KD_ID, tree.KD_ASSIGN, tree.KD_DONT_CARE) && tk.Type == token.TT_NUMBER:
			tr.SetRight(tk, tree.KD_OPERAND)
		default:
			m := "Token[" + strconv.Itoa(i) + "] does not match any parsing rules"
			return nil, errors.New(m)
		}
	}

	return tr, nil
}
