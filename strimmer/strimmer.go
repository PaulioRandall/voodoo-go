package strimmer

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

// Strim normalises an array of lexemes and converts them to tokens
// ready for the syntax analyser. It assumes each lexeme is correct
// and valid even if together they do not form a valid statement.
//
// Normalising involves:
// -> Removing whitespace lexemes
// -> Removing comment lexemes
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Merging any explicit number sign with it's number
// -> Converting all letters to lowercase (Except string literals)
// -> Merging the numbers with their sign (if one has been declared)
func Strim(ls []lexeme.Lexeme) ([]Token, StrimError) {

	ts := []Token{}

	for _, l := range ls {
		switch {
		case l.Type == lexeme.WHITESPACE:
			continue
		case l.Type == lexeme.COMMENT:
			continue
		case l.Type == lexeme.STRING:
			penultimate := len(l.Val) - 1
			l.Val = l.Val[1:penultimate]
		case l.Type == lexeme.NUMBER:
			l.Val = strings.ReplaceAll(l.Val, `_`, ``)
		}

		t := Token(l)
		ts = append(ts, t)
	}

	return ts, nil
}
