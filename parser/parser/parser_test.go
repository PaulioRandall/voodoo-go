package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func doTestParse(t *testing.T, in []token.Token, exp *tree.Tree) {
	tr, err := Parse(in)
	require.Nil(t, err)
	require.NotNil(t, tr)
	assertTree(t, exp, tr, `Trunk`)
}

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

// Rule 0:
//  Predicate:   The subject token does not match another rule.
//  Consequence: Abandon parsing and generate an error.
func TestParse_Rule_0(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 0, ``, token.TT_SPACE),
	}
	doTestParseError(t, in)
}

//Rule 1:
//  Predicate:   The left node has no kind
//               AND the subject token has the IDENTIFIER type.
//  Consequence: Place the subject token in the left node
//               AND assign the left node the IDENTIFIER kind.
func TestParse_Rule_1(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
	}
	expLeft := dummyTree(tree.KD_ID, in[0], nil, nil)
	exp := dummyTree(tree.KD_UNDEFINED, token.EMPTY, expLeft, nil)
	doTestParse(t, in, exp)
}

//Rule 2:
//  Predicate:   The left node has the IDENTIFIER or UNION kind
//               AND the subject token has the ASSIGNMENT type.
//  Consequence: Place the subject token in the current node
//               AND assign the current node the ASSIGNMENT kind.
func TestParse_Rule_2(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
	}
	expLeft := dummyTree(tree.KD_ID, in[0], nil, nil)
	exp := dummyTree(tree.KD_ASSIGN, in[1], expLeft, nil)
	doTestParse(t, in, exp)
}

//Rule 3:
//  Predicate:   The left node has the IDENTIFIER kind
//               AND the current node has the ASSIGNMENT kind
//               AND the subject token has the NUMBER type.
//  Consequence: Place the subject token in the right node
//               AND assign the right node the OPERAND kind.
func TestParse_Rule_3(t *testing.T) {
	in := []token.Token{
		token.DummyToken(0, 0, 1, `x`, token.TT_ID),
		token.DummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		token.DummyToken(0, 5, 6, `1`, token.TT_NUMBER),
	}
	expLeft := dummyTree(tree.KD_ID, in[0], nil, nil)
	expRight := dummyTree(tree.KD_OPERAND, in[2], nil, nil)
	exp := dummyTree(tree.KD_ASSIGN, in[1], expLeft, expRight)
	doTestParse(t, in, exp)
}
