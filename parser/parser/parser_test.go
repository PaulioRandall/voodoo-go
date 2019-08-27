package parser

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func printTree(tr *tree.Tree, node string) {
	printTreeIndented(tr, 0, node)
}

func printTreeIndented(tr *tree.Tree, indent int, node string) {
	pre := strings.Repeat(" ", indent)
	preMore := strings.Repeat(" ", indent+2)

	fmt.Print(pre + node + ": ")

	if tr == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(pre + "{")
	fmt.Println(preMore + "HasParent: " + strconv.FormatBool(tr.Parent != nil))
	fmt.Println(preMore + "Kind: " + tree.KindName(tr.Kind))
	fmt.Println(preMore + "Token: " + tr.Token.String())
	printTreeIndented(tr.Left, indent+2, node+".Left")
	printTreeIndented(tr.Left, indent+2, node+".Right")
	fmt.Println(pre + "}")
}

func TestParse_1(t *testing.T) {
	// x <- 1
	in := []token.Token{
		token.OfType(token.TT_ID),
		token.OfType(token.TT_ASSIGN),
		token.OfType(token.TT_NUMBER),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: in[1],
		Left: &tree.Tree{
			Kind:  tree.KD_ID,
			Token: in[0],
		},
		Right: &tree.Tree{
			Kind:  tree.KD_OPERAND,
			Token: in[2],
		},
	}

	act, err := Parse(in)
	printTree(exp, "Trunk")
	printTree(act, "Trunk")
	require.Nil(t, err)
	assertTree(t, exp, act, `Trunk`)
}

func OLD_TestParse_2(t *testing.T) {
	// x, y <- 1, 2
	in := []token.Token{
		token.OfTypeUnique(token.TT_ID),
		token.OfTypeUnique(token.TT_VALUE_DELIM),
		token.OfTypeUnique(token.TT_ID),
		token.OfTypeUnique(token.TT_ASSIGN),
		token.OfTypeUnique(token.TT_NUMBER),
		//token.OfTypeUnique(token.TT_VALUE_DELIM),
		//token.OfTypeUnique(token.TT_NUMBER),
	}

	exp := &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: in[3],
		Left: &tree.Tree{
			Kind:  tree.KD_UNION,
			Token: in[1],
			Left: &tree.Tree{
				Kind:  tree.KD_ID,
				Token: in[0],
			},
			Right: &tree.Tree{
				Kind:  tree.KD_ID,
				Token: in[2],
			},
		},
		//Right: &tree.Tree{
		//Kind:  tree.KD_UNION,
		//Token: in[5],
		//Left: &tree.Tree{
		//Kind:  tree.KD_OPERAND,
		//Token: in[4],
		//},
		//				Kind:  tree.KD_OPERAND,
		//				Token: in[6],
		//			},
		//},
	}

	act, err := Parse(in)
	require.Nil(t, err)
	assertTree(t, exp, act, `Trunk`)
}
