package analyser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/operation"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// Analyse parses an array of tokens into a set of instructions.
// Analysis involves:
// 1: Expanding parenthesis out adding new identifiers as needed.
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
func Analyse(in []symbol.Token) (operation.InstructionSet, fault.Fault) {
	return nil, nil
}
