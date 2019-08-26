package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func doTestParseToken(
	t *testing.T,
	tr *tree.Tree,
	in token.Token,
	expOk bool,
	exp *tree.Tree) {

	ok := parseToken(tr, in)
	if expOk {
		require.True(t, ok, `Expected 'parseToken(..) == true'`)
		assertTree(t, exp, tr, `Trunk`)
	} else {
		require.False(t, ok, `Expected 'parseToken(..) == false'`)
	}
}

func doTestParse(t *testing.T, in []token.Token, exp *tree.Tree) {
	tr, err := Parse(in)
	require.Nil(t, err)
	require.NotNil(t, tr)
	assertTree(t, exp, tr, `Trunk`)
}

// MARK: Can be removed?
func doTestParseError(t *testing.T, in []token.Token) {
	tr, err := Parse(in)
	assert.True(t, err != nil, `Expected 'err != nil'`)
	assert.True(t, tr == nil, `Expected error but 'tree != nil'`)
}

func assertTree(t *testing.T, exp *tree.Tree, act *tree.Tree, node string) {
	if exp == nil {
		assert.Nil(t, act, `Expected '%s == nil'`, node)
		return
	}

	if !token.AssertToken(t, &exp.Token, &act.Token) {
		t.Logf(`%s: Not the expected token`, node)
	}
	assert.Equal(t, exp.Kind, act.Kind, `%s: Not the expected kind`, node)

	assertTree(t, exp.Left, act.Left, node+`.Left`)
	assertTree(t, exp.Right, act.Right, node+`.Right`)
}

func dummyTree(k tree.Kind, t token.Token, l *tree.Tree, r *tree.Tree) *tree.Tree {
	return &tree.Tree{
		Kind:  k,
		Token: t,
		Left:  l,
		Right: r,
	}
}

func TestParse_Rule_0(t *testing.T) {
	tr := tree.New()
	in := token.DummyToken(0, 0, 0, ``, token.TT_SPACE)
	doTestParseToken(t, tr, in, false, nil)
}

func TestParse_Rule_1(t *testing.T) {
	tr := tree.New()

	in := token.DummyToken(0, 0, 1, `x`, token.TT_ID)

	exp := tree.Copy(tr)
	exp.Left = dummyTree(tree.KD_ID, in, nil, nil)

	doTestParseToken(t, tr, in, true, exp)
}

func TestParse_Rule_2a(t *testing.T) {
	tr := &tree.Tree{
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	in := token.DummyToken(0, 1, 2, `,`, token.TT_VALUE_DELIM)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_UNION
	exp.Token = in

	doTestParseToken(t, tr, in, true, exp)
}

func TestParse_Rule_2b(t *testing.T) {
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

	in := token.DummyToken(0, 3, 4, `,`, token.TT_VALUE_DELIM)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_UNION
	exp.Token = in

	doTestParseToken(t, tr, in, true, exp)
}

func TestParse_Rule_3a(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	in := token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_ASSIGN
	exp.Token = in

	doTestParseToken(t, tr, in, true, exp)
}

func TestParse_Rule_3b(t *testing.T) {
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

	in := token.DummyToken(0, 3, 5, `<-`, token.TT_ASSIGN)

	exp := tree.Copy(tr)
	exp.Kind = tree.KD_ASSIGN
	exp.Token = in

	doTestParseToken(t, tr, in, true, exp)
}

func TestParse_Rule_4a(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
		Left: &tree.Tree{
			Kind: tree.KD_ID,
		},
	}

	in := token.DummyToken(0, 5, 6, `1`, token.TT_NUMBER)

	exp := tree.Copy(tr)
	exp.Right = dummyTree(tree.KD_OPERAND, in, nil, nil)

	doTestParseToken(t, tr, in, true, exp)
}

/*
func TestParse_Rule_4b(t *testing.T) {
	tr := &tree.Tree{
		Kind: tree.KD_ASSIGN,
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

	in := token.DummyToken(0, 5, 6, `1`, token.TT_NUMBER)

	exp := tree.Copy(tr)
	exp.Right = dummyTree(tree.KD_OPERAND, in, nil, nil)

	doTestParseToken(t, tr, in, true, exp)
}
*/
