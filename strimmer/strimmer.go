package strimmer

import (
  "github.com/PaulioRandall/voodoo-go/lexeme"
)

// Strim normalises an array of lexemes, performs some error
// checking on them, and converts them to tokens ready for
// a syntax analyser.
//
// Normalising involves:
// -> Removing whitespace lexemes
// -> Removing comment lexemes
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Converting all letters to lowercase
func Strim(ls []lexeme.Lexeme) ([]Token, StrimError) {
  return nil, nil
}