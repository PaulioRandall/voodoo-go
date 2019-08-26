package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// predicate represents the predicate of a parse rule. It returns true if the
// input satisfies the rules conditions. It is paired with a consequence
// function which is invoked if the predicate returns true. Note that the
// return value depends upon the state of the tree and the specific token so any
// changes to either input renders the previous response invalid.
type predicate func(*tree.Tree, token.Token) bool

// consequence modifies the input trees state based on the parse rule which it
// implements. A consequence should only be invoked if the paired predicate
// returns true and the tree and the token are not modified in the meantime.
type consequence func(*tree.Tree, token.Token) *tree.Tree

// Predicate: The left, current, and right node have no kind
//            AND the subject token has the IDENTIFIER type.
func rule_1_predicate(tr *tree.Tree, tk token.Token) bool {
	return tr.IsLeft(tree.KD_UNDEFINED) &&
		tr.Is(tree.KD_UNDEFINED) &&
		tr.IsRight(tree.KD_UNDEFINED) &&
		tk.Type == token.TT_ID
}

// Consequence: Place the subject token in the current node
//              AND assign the current node the IDENTIFIER kind.
func rule_1_consequence(tr *tree.Tree, tk token.Token) *tree.Tree {
	tr.Set(tk, tree.KD_ID)
	return tr
}

// Predicate: The current node has the IDENTIFIER or UNION kind
//            AND the subject token has the VALUE_DELIM type.
func rule_2_predicate(tr *tree.Tree, tk token.Token) bool {
	ok := tr.Is(tree.KD_ID) ||
		tr.Is(tree.KD_UNION)
	return ok &&
		tk.Type == token.TT_VALUE_DELIM
}

// Consequence: Create a new node
//              AND place the subject token in the new node
//              AND assign the new node the UNION kind
//              AND place the current node as the new nodes left
//              AND set the new node as the current
func rule_2_consequence(tr *tree.Tree, tk token.Token) *tree.Tree {
	return &tree.Tree{
		Kind:  tree.KD_UNION,
		Token: tk,
		Left:  tr,
	}
}

// Predicate: The left node has the IDENTIFIER or UNION kind
//            AND the current node has the UNION kind
//            AND the subject token has the IDENTIFIER type.
func rule_3_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return a &&
		tr.Is(tree.KD_UNION) &&
		tk.Type == token.TT_ID
}

// Consequence: Place the subject token in the right node
//              AND assign the right node the IDENTIFIER kind.
func rule_3_consequence(tr *tree.Tree, tk token.Token) *tree.Tree {
	tr.SetRight(tk, tree.KD_ID)
	return tr
}

// Predicate: The current node has the IDENTIFIER or UNION kind
//            AND the subject token has the ASSIGNMENT type.
func rule_4_predicate(tr *tree.Tree, tk token.Token) bool {
	ok := tr.Is(tree.KD_ID) ||
		tr.Is(tree.KD_UNION)
	return ok &&
		tk.Type == token.TT_ASSIGN
}

// Consequence: Create a new node
//              AND place the subject token in the new node
//              AND assign the new node the ASSIGNMENT kind
//              AND place the current node as the new nodes left
//              AND set the new node as the current
func rule_4_consequence(tr *tree.Tree, tk token.Token) *tree.Tree {
	return &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: tk,
		Left:  tr,
	}
}

// Predicate: The left node has the IDENTIFIER or UNION kind
//            AND the current node has the ASSIGNMENT kind
//            AND the subject token has the NUMBER type.
func rule_5_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return a &&
		tr.Is(tree.KD_ASSIGN) &&
		tk.Type == token.TT_NUMBER
}

// Consequence: Place the subject token in the right node
//              AND assign the right node the OPERAND kind.
func rule_5_consequence(tr *tree.Tree, tk token.Token) *tree.Tree {
	tr.SetRight(tk, tree.KD_OPERAND)
	return tr
}

// Predicate: The left and right nodes have a kind
//            AND the current node has the UNION kind
//            AND the subject token has the ASSIGNMENT type.
func rule_6_predicate(tr *tree.Tree, tk token.Token) bool {
	return !tr.IsLeft(tree.KD_UNDEFINED) &&
		!tr.IsRight(tree.KD_UNDEFINED) &&
		tr.Is(tree.KD_UNION) &&
		tk.Type == token.TT_ASSIGN
}

// Consequence: Create a new node
//              AND place the subject token in the new node
//              AND assign the new node the ASSIGNMENT kind
//              AND place the current node as the new nodes left
//              AND set the new node as the current
func rule_6_consequence(tr *tree.Tree, tk token.Token) *tree.Tree {
	return &tree.Tree{
		Kind:  tree.KD_ASSIGN,
		Token: tk,
		Left:  tr,
	}
}
