package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/require"
)

func TestRule_1_consequence(t *testing.T) {
	tr := tree.New()

	tk := token.OfType(token.TT_ID)

	exp := &tree.Tree{
		Kind:  tree.KD_ID,
		Token: tk,
	}

	require.True(t, rule_1_predicate(tr, tk))
	tr = rule_1_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_2a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ID,
	}

	tk := token.OfType(token.TT_VALUE_DELIM)

	exp := &tree.Tree{
		Kind:  tree.KD_UNION,
		Token: tk,
		Left:  tr,
	}

	require.True(t, rule_2_predicate(tr, tk))
	tr = rule_2_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_2b_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
	}

	tk := token.OfType(token.TT_VALUE_DELIM)

	exp := &tree.Tree{
		Kind:  tree.KD_UNION,
		Token: tk,
		Left:  tr,
	}

	require.True(t, rule_2_predicate(tr, tk))
	tr = rule_2_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_3a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_ID)

	exp := tree.Copy(tr, nil)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_ID,
		Token: tk,
	}

	require.True(t, rule_3_predicate(tr, tk))
	tr = rule_3_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_3b_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
		},
	}

	tk := token.OfType(token.TT_ID)

	exp := tree.Copy(tr, nil)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_ID,
		Token: tk,
	}

	require.True(t, rule_3_predicate(tr, tk))
	tr = rule_3_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_4a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ID,
	}

	tk := token.OfType(token.TT_ASSIGN)

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: tk,
		Left:  tr,
	}

	require.True(t, rule_4_predicate(tr, tk))
	tr = rule_4_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_4b_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
	}

	tk := token.OfType(token.TT_ASSIGN)

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: tk,
		Left:  tr,
	}

	require.True(t, rule_4_predicate(tr, tk))
	tr = rule_4_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_5_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
		Right: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_ASSIGN)

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: tk,
		Left:  tr,
	}

	require.True(t, rule_5_predicate(tr, tk))
	tr = rule_5_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_6a_consequence(t *testing.T) {
	curr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_NUMBER)

	exp := tree.Copy(curr, nil)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_OPERAND,
		Token: tk,
	}

	require.True(t, rule_6_predicate(curr, tk))
	right := rule_6_consequence(curr, tk)
	assertTree(t, exp, curr, `Trunk`)
	assertTree(t, exp.Right, right, `Trunk.Right`)
}

func TestRule_6b_consequence(t *testing.T) {
	curr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
		},
	}

	tk := token.OfType(token.TT_NUMBER)

	exp := tree.Copy(curr, nil)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_OPERAND,
		Token: tk,
	}

	require.True(t, rule_6_predicate(curr, tk))
	right := rule_6_consequence(curr, tk)
	assertTree(t, exp, curr, `Trunk`)
	assertTree(t, exp.Right, right, `Trunk.Right`)
}

/*
func TestRule_7a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Left: &tree.Tree{
			Kind: tree.KD_OPERAND,
		},
	}

	tk := token.OfType(token.TT_VALUE_DELIM)

	exp := tree.Copy(tr, nil)
	exp.Kind = tree.KD_UNION
	exp.Token = tk

	require.True(t, rule_7_predicate(tr, tk))
	tr = rule_7_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_7b_consequence(t *testing.T) {
	tr := &tree.Tree{
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
		},
	}

	tk := token.OfType(token.TT_VALUE_DELIM)

	exp := tree.Copy(tr, nil)
	exp.Kind = tree.KD_UNION
	exp.Token = tk

	require.True(t, rule_7_predicate(tr, tk))
	tr = rule_7_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}
*/
