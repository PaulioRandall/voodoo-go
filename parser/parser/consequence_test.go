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

	exp := tree.Copy(tr)
	exp.Left = &tree.Tree{
		Kind:  tree.KD_ID,
		Token: tk,
	}

	require.True(t, rule_1_predicate(tr, tk))
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

	require.True(t, rule_2_predicate(tr, tk))
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

	require.True(t, rule_2_predicate(tr, tk))
	rule_2_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_3a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_ASSIGN)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_ASSIGN
	exp.Token = tk

	require.True(t, rule_3_predicate(tr, tk))
	rule_3_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_3b_consequence(t *testing.T) {
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

	tk := token.OfType(token.TT_ASSIGN)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_ASSIGN
	exp.Token = tk

	require.True(t, rule_3_predicate(tr, tk))
	rule_3_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_4a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_NUMBER)

	exp := tree.Copy(tr)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_OPERAND,
		Token: tk,
	}

	require.True(t, rule_4_predicate(tr, tk))
	rule_4_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_4b_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
		},
	}

	tk := token.OfType(token.TT_NUMBER)

	exp := tree.Copy(tr)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_OPERAND,
		Token: tk,
	}

	require.True(t, rule_4_predicate(tr, tk))
	rule_4_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_5a_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_ID)

	exp := tree.Copy(tr)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_ID,
		Token: tk,
	}

	require.True(t, rule_5_predicate(tr, tk))
	rule_5_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}

func TestRule_5b_consequence(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
		},
	}

	tk := token.OfType(token.TT_ID)

	exp := tree.Copy(tr)
	exp.Right = &tree.Tree{
		Kind:  tree.KD_ID,
		Token: tk,
	}

	require.True(t, rule_5_predicate(tr, tk))
	rule_5_consequence(tr, tk)
	assertTree(t, exp, tr, `Trunk`)
}
