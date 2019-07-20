package strimmer

import (
  "github.com/PaulioRandall/voodoo-go/lexeme"
)

// Strim normalises an array of lexemes and converts them to tokens
// ready for a syntax analyser. 
//
// Normalising involves:
// -> Removing whitespace lexemes
// -> Removing comment lexemes
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Merging any explicit number sign with it's number
// -> Converting all letters to lowercase (Except string literals)
func Strim(ls []lexeme.Lexeme) ([]Token, StrimError) {
  return nil, nil
}