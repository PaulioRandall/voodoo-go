package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/assert"
)

func TestRule_1_predicate(t *testing.T) {
	tr := tree.New()

	tk := token.OfType(token.TT_ID)
	r := rule_1_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_1_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_2a_predicate(t *testing.T) {
	tr := &tree.Tree{
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_VALUE_DELIM)
	r := rule_2_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_2_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_2b_predicate(t *testing.T) {
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
	r := rule_2_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_2_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_3a_predicate(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_ASSIGN)
	r := rule_3_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_3_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_3b_predicate(t *testing.T) {
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
	r := rule_3_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_3_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_4a_predicate(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_NUMBER)
	r := rule_4_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_4_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_4b_predicate(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
		},
	}

	tk := token.OfType(token.TT_NUMBER)
	r := rule_4_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_4_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_5a_predicate(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	tk := token.OfType(token.TT_ID)
	r := rule_5_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_5_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}

func TestRule_5b_predicate(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_UNION,
		Left: &tree.Tree{
			Kind: tree.KD_UNION,
		},
	}

	tk := token.OfType(token.TT_ID)
	r := rule_5_predicate(tree.Copy(tr), tk)
	assert.True(t, r)

	tk = token.OfType(token.TT_SPACE)
	r = rule_5_predicate(tree.Copy(tr), tk)
	assert.False(t, r)
}
