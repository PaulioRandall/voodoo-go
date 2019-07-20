package strimmer

import (
  "github.com/PaulioRandall/voodoo-go/lexeme"
)

// Strim normalises an array of lexemes and converts them to tokens
// ready for the syntax analyser.
//
// Normalising involves:
// -> Removing whitespace lexemes
// -> Removing comment lexemes
// -> Removing quote marks from string literals
// -> Removing underscores from numbers
// -> Merging any explicit number sign with it's number
// -> Converting all letters to lowercase (Except string literals)
func Strim(ls []lexeme.Lexeme) ([]Token, StrimError) {

	ts := []Token{}

  for _, l := range ls {
    switch {
    case l.Type == lexeme.WHITESPACE:
      continue
    }
    
    t := Token(l)
    ts = append(ts, t)
  }

  return ts, nil
}