package syntax

import (
	"github.com/PaulioRandall/voodoo-go/symbol"
	instruct "github.com/PaulioRandall/voodoo-go/instruction"
)

// Analyse parses an array of tokens into a set of instructions.
// Analysis involves:
// 1: Scanning the token array for precedence brackets, the
//    the curvy ones, and splitting the token array into
//    smaller ones until no brackets are left, an error is
//    returned if each brace doesn't have an opposing
//    partner. E.g...
//    1st:
//      #0: `x <- (a + b - (c * d))` - 13 tokesn
//    2nd:
//      #0: `x <- #1` - 3 tokens
//      #1: `a + b - (c * d)` - 9 tokens
//    3rd:
//      #0: `x <- #1` - 3 tokens
//      #1: `a + b - #2` - 5 tokens
//      #2: `c * d` - 3 tokens
func Analyse(ts []symbol.Token) (instruct.InstructionSet, instruct.InsError) {
	return nil, nil
}
