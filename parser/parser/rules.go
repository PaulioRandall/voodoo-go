package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// predicate represents the conditional part of a parse rule, the conditional
// part of a glorified IF statement.
type predicate func(*tree.Tree, token.Token) bool

// consequence represents the state changing part of a parse rule, the THEN part
// of a glorified IF statement.
type consequence func(*tree.Tree, token.Token) *tree.Tree

// Meta rules: rules for writing rules.
// 1) Predicate invocations must be idempotent, that is the result of two calls
//    with the same input must yeild the same result.
// 2) Predicate invocations must be pure (void of side effects), that is a call
//    must not change the state of the input tree or token; that is what the
//    consequence is for.
// 3) Consequence invocations must not have any conditional logic; that is what
//    the predicate is for.
// 4) Nil must never be returned from a consequence invocation as the value
//    determines the current node. Doing so is an error.
// 5) Don't be afraid to just add new rules if it's unclear if an existing rule
//    can be modified. Rules can be merged later when the duplication becomes
//    visible.

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

// Predicate: The left and right nodes both have or both don't have a kind
//            AND the current node has the IDENTIFIER or UNION kind
//            AND the subject token has the VALUE_DELIM type.
func rule_2_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.Is(tree.KD_ID) ||
		tr.Is(tree.KD_UNION)
	b1 := tr.IsLeft(tree.KD_UNDEFINED)
	b2 := tr.IsRight(tree.KD_UNDEFINED)
	b := b1 == b2
	return a && b &&
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
//            AND the right node has no kind
//            AND the current node has the UNION kind
//            AND the subject token has the IDENTIFIER type.
func rule_3_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return a &&
		tr.IsRight(tree.KD_UNDEFINED) &&
		tr.Is(tree.KD_UNION) &&
		tk.Type == token.TT_ID
}

// Consequence: Place the subject token in the right node
//              AND assign the right node the IDENTIFIER kind.
func rule_3_consequence(tr *tree.Tree, tk token.Token) *tree.Tree {
	tr.SetRight(tk, tree.KD_ID)
	return tr
}

// Predicate: The left and right nodes both have or both don't have a kind
//            AND current node has the IDENTIFIER or UNION kind
//            AND the subject token has the ASSIGNMENT type.
func rule_4_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.Is(tree.KD_ID) ||
		tr.Is(tree.KD_UNION)
	b1 := tr.IsLeft(tree.KD_UNDEFINED)
	b2 := tr.IsRight(tree.KD_UNDEFINED)
	b := b1 == b2
	return a && b &&
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
//            AND the right node has no kind
//            AND the current node has the ASSIGNMENT kind
//            AND the subject token has the NUMBER type.
func rule_5_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return a &&
		tr.IsRight(tree.KD_UNDEFINED) &&
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
