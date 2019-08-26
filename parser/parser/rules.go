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
type consequence func(*tree.Tree, token.Token)

//  Predicate: The left node has no kind
//             AND the subject token has the IDENTIFIER type.
func rule_1_predicate(tr *tree.Tree, tk token.Token) bool {
	return tr.IsLeft(tree.KD_UNDEFINED) &&
		tk.Type == token.TT_ID
}

//  Consequence: Place the subject token in the left node
//               AND assign the left node the IDENTIFIER kind.
func rule_1_consequence(tr *tree.Tree, tk token.Token) {
	tr.SetLeft(tk, tree.KD_ID)
}

//  Predicate: The left node has the IDENTIFIER or UNION kind
//             AND the subject token has the VALUE_DELIM type.
func rule_2_predicate(tr *tree.Tree, tk token.Token) bool {
	ok := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return ok &&
		tk.Type == token.TT_VALUE_DELIM
}

//  Consequence: Place the subject token in the current node
//               AND assign the current node the UNION kind.
func rule_2_consequence(tr *tree.Tree, tk token.Token) {
	tr.Set(tk, tree.KD_UNION)
}

//  Predicate: The left node has the IDENTIFIER or UNION kind
//             AND the subject token has the ASSIGNMENT type.
func rule_3_predicate(tr *tree.Tree, tk token.Token) bool {
	ok := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return ok &&
		tk.Type == token.TT_ASSIGN
}

//  Consequence: Place the subject token in the current node
//               AND assign the current node the ASSIGNMENT kind.
func rule_3_consequence(tr *tree.Tree, tk token.Token) {
	tr.Set(tk, tree.KD_ASSIGN)
}

//  Predicate: The left node has the IDENTIFIER or UNION kind
//             AND the current node has the ASSIGNMENT kind
//             AND the subject token has the NUMBER type.
func rule_4_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return a &&
		tr.Is(tree.KD_ASSIGN) &&
		tk.Type == token.TT_NUMBER
}

//  Consequence: Place the subject token in the right node
//               AND assign the right node the OPERAND kind.
func rule_4_consequence(tr *tree.Tree, tk token.Token) {
	tr.SetRight(tk, tree.KD_OPERAND)
}

//  Predicate: The left node has the IDENTIFIER or UNION kind
//             AND the current node has the UNION kind
//             AND the subject token has the IDENTIFIER type.
func rule_5_predicate(tr *tree.Tree, tk token.Token) bool {
	a := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return a &&
		tr.Is(tree.KD_UNION) &&
		tk.Type == token.TT_ID
}

//  Consequence: Place the subject token in the right node
//               AND assign the right node the IDENTIFIER kind.
func rule_5_consequence(tr *tree.Tree, tk token.Token) {
	tr.SetRight(tk, tree.KD_ID)
}
