package parser

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/tree"
)

// findRule finds the first matching rule number.
func findRule(tr *tree.Tree, tk token.Token) int {
	switch {
	case rule_1(tr, tk):
		//  Predicate:   The left node has no kind
		//               AND the subject token has the IDENTIFIER type.
		//  Consequence: Place the subject token in the left node
		//               AND assign the left node the IDENTIFIER kind.
		return 1
	case rule_2(tr, tk):
		//  Predicate:   The left node has the IDENTIFIER kind
		//               AND the subject token has the VALUE_DELIM type.
		//  Consequence: Place the subject token in the current node
		//               AND assign the current node the UNION kind.
		return 2
	case rule_3(tr, tk):
		//  Predicate:   The left node has the IDENTIFIER kind
		//               AND the current node has the ASSIGNMENT kind
		//               AND the subject token has the NUMBER type.
		//  Consequence: Place the subject token in the right node
		//               AND assign the right node the OPERAND kind.
		return 3
	case rule_4(tr, tk):
		//  Predicate:   The left node has the IDENTIFIER or UNION kind
		//               AND the subject token has the ASSIGNMENT type.
		//  Consequence: Place the subject token in the current node
		//               AND assign the current node the ASSIGNMENT kind.
		return 4
	default:
		//  Predicate:   The subject token does not match any other rule.
		//  Consequence: Abandon parsing and generate an error.
		return 0
	}
}

// rule_1 returns true if the input satisfies Rule 1.
func rule_1(tr *tree.Tree, tk token.Token) bool {
	return tr.IsLeft(tree.KD_UNDEFINED) &&
		tk.Type == token.TT_ID
}

// rule_2 returns true if the input satisfies Rule 2.
func rule_2(tr *tree.Tree, tk token.Token) bool {
	return tr.IsLeft(tree.KD_ID) &&
		tk.Type == token.TT_VALUE_DELIM
}

// rule_3 returns true if the input satisfies Rule 3.
func rule_3(tr *tree.Tree, tk token.Token) bool {
	return tr.IsLeft(tree.KD_ID) &&
		tr.Is(tree.KD_ASSIGN) &&
		tk.Type == token.TT_NUMBER
}

// rule_4 returns true if the input satisfies Rule 4.
func rule_4(tr *tree.Tree, tk token.Token) bool {
	ok := tr.IsLeft(tree.KD_ID) ||
		tr.IsLeft(tree.KD_UNION)
	return ok && tk.Type == token.TT_ASSIGN
}
