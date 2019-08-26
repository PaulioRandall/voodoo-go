package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

func TestRule_1_consequence(t *testing.T) {
	tr := tree.New()

	tk := token.OfType(token.TT_ID)

	exp := tree.Copy(tr)
	exp.Left = dummyTree(tree.KD_ID, tk, nil, nil)
	rule_1_consequence(tr, tk)

	assertTree(t, exp, tr, `Trunk`)
}
