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
	var ok bool

	for i, tk := range in {
		tr, ok = parseToken(tr, tk)
		if !ok {
			m := "Token[" + strconv.Itoa(i) + "] does not match any parsing rules"
			return nil, errors.New(m)
		}
	}

	return tr, nil
}

// parseToken applies the first matching parse rule, with token as subject, to
// the tree.
func parseToken(tr *tree.Tree, tk token.Token) (*tree.Tree, bool) {
	switch {
	case rule_1_predicate(tr, tk):
		tr = rule_1_consequence(tr, tk)
	case rule_2_predicate(tr, tk):
		tr = rule_2_consequence(tr, tk)
	case rule_3_predicate(tr, tk):
		tr = rule_3_consequence(tr, tk)
	case rule_4_predicate(tr, tk):
		tr = rule_4_consequence(tr, tk)
	case rule_5_predicate(tr, tk):
		tr = rule_5_consequence(tr, tk)
	case rule_6_predicate(tr, tk):
		tr = rule_6_consequence(tr, tk)
	default:
		return tr, false
	}

	return tr, true
}
