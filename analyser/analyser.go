package analyser

import (
	"github.com/PaulioRandall/voodoo-go/operation"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// Analyse parses an array of tokens into a set of instructions.
// Analysis involves:
// 1: Expanding precedence brackets out and placing them at the
//    front of the array of tokens so they are executed first.
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
func Analyse(ts []symbol.Token) (operation.InstructionSet, AnaError) {

	return nil, nil
}

// expandBrackets scans the token array for precedence brackets,
// the curvy ones, then splitting the token array into smaller
// ones until no brackets are left, an error is returned if a
// brace doesn't have an opposing partner. Braces are removed
// along the way with temporary variables being created and
// inserted where needed. E.g...
//   Input:
//     #0 = [x, <-, (, a, +, 12.3, -, (, c, *, d, ), )]
//   Turns into:
//     #0 = [x, <-, #1]
//     #1 = [a, +, 12.3, -, (, c, *, d, )]
//   Turns into:
//     #0 = [x, <-, #1]
//     #1 = [a, +, 12.3, -, #2]
//     #2 = [c, *, d]
//   Ordering is always reversed:
//     [#2, #1, #0]
func expandBrackets(ts []symbol.Token) ([][]symbol.Token, AnaError) {
	r := [][]symbol.Token{}

loop:
	for i := 0; i < len(ts); i++ {
		o := indexOf(ts, i, symbol.CURVED_BRACE_OPEN)
		c := 1

		switch {
		case o == -1 && c == -1:
			break loop

		case o != -1 && c == -1:
			m := "Missing closing brace to corresponding opening one"
			err := NewAnaError(m, o)
			return nil, err

		case o == -1 && c != -1:
			m := "Didn't expect to find a closing brace without a corresponding opening one"
			err := NewAnaError(m, c)
			return nil, err
		}

	}

	return r, nil
}

// indexOf returns the next index of the symbol with the specified type
// of -1 if no matching token is found.
func indexOf(a []symbol.Token, start int, t symbol.SymbolType) int {
	for i, tk := range a {
		if i < start {
			continue
		}

		if tk.Type == t {
			return i
		}
	}

	return -1
}
