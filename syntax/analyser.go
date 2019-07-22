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
//    returned if a brace doesn't have an opposing partner.
//    Braces are removed along the way with temporary variables
//    being created and inserted where needed.
//    E.g...
//    Input:
//      #0 = [x, <-, (, a, +, 12.3, -, (, c, *, d, ), )] = 13 tokens
//    Turns into:
//      #0 = [x, <-, #1]
//      #1 = [a, +, 12.3, -, (, c, *, d, )]
//    Turns into:
//      #0 = [x, <-, #1]
//      #1 = [a, +, 12.3, -, #2]
//      #2 = [c, *, d]
//    Ordering is reversed:
//      [#2, #1, #0]
// 2: Each token array is then split into even smaller ones such
//    that each contains exactly one instruction, error checking
//    where possible.
//    E.g... 
//    Input:
//      #2 = [c, *, d]
//      #1 = [a, +, 12.3, -, #2]
//      #0 = [x, <-, #1]
//    Turns into:
//      #2.1 = [c, *, d]
//      #1.2 = [a, +, 12.3]
//      #1.1 = [#1.2, -, #2.1]
//      #0 = [x, <-, #1.1]
// 3: Finally each token array is converted into its
//    instruction counter part. Number and boolean literals
//    are converted into Go float64 and bools accordingly.
//    E.g...
//    Input:
//      #2.1 = [c, *, d]
//      #1.2 = [a, +, 12.3]
//      #1.1 = [#1.2, -, #2.1]
//      #0 = [x, <-, #1.1]
//    Turns into:
//      InstructionSet{
//        Multiply{
//          Left: c,
//          Right: d,
//          As: #2.1,
//        },
//        Add{
//          Left: a,
//          Right: 12.3,
//          As: #1.2,        
//        },
//        Subtract{
//          Left: #1.2,
//          Right: #2.1,
//          As: #1.1,        
//        },
//        Assign{
//          Val: #2.1,
//          As: x,        
//        },
//      }
func Analyse(ts []symbol.Token) (instruct.InstructionSet, instruct.InsError) {
	return nil, nil
}
