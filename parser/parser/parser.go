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
		switch findRule(tr, tk) {
		case 1:
			tr.SetLeft(tk, tree.KD_ID)
		case 2:
			tr.Set(tk, tree.KD_UNION)
		case 3:
			tr.SetRight(tk, tree.KD_OPERAND)
		case 4:
			tr.Set(tk, tree.KD_ASSIGN)
		default:
			m := "Token[" + strconv.Itoa(i) + "] does not match any parsing rules"
			return nil, errors.New(m)
		}
	}

	return tr, nil
}
