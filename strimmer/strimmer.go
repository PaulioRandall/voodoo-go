package strimmer

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

// Strim normalises an array of lexemes and converts them to tokens
// ready for the syntax analyser. It assumes each lexeme is correct
// and valid even if together they do not form a valid statement;
// validation will happen later.
//
// Normalising involves:
// -> Removing whitespace lexemes
// -> Removing comment lexemes
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Converting all letters to lowercase (Except string literals)
func Strim(in []symbol.Lexeme) []symbol.Token {

	out := []symbol.Token{}

	for _, l := range in {
		switch {
		case l.Type == symbol.WHITESPACE:
			continue
		case l.Type == symbol.COMMENT:
			continue
		case l.Type == symbol.LITERAL_STRING:
			penultimate := len(l.Val) - 1
			l.Val = l.Val[1:penultimate]
		case l.Type == symbol.LITERAL_NUMBER:
			l.Val = strings.ReplaceAll(l.Val, `_`, ``)
		case l.Type > symbol.ALPHABETIC_START && l.Type < symbol.ALPHABETIC_END:
			l.Val = strings.ToLower(l.Val)
		}

		t := symbol.Token(l)
		out = append(out, t)
	}

	return out
}
