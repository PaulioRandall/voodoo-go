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

func TestRule_2a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_VALUE_DELIM)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_UNION
	exp.Token = tk

	rule_2_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_2b_consequence(t *testing.T) {
	tr := &tree.Tree{
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
			Left: &tree.Tree{
				Kind: tree.KD_ID,
			},
			Right: &tree.Tree{
				Kind: tree.KD_ID,
			},
		},
	}

	tk := token.OfType(token.TT_VALUE_DELIM)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_UNION
	exp.Token = tk

	rule_2_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}
