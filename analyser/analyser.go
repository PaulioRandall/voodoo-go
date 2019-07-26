package analyser

import (
	"fmt"
	"strconv"

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
func expandBrackets(a []symbol.Token) ([][]symbol.Token, AnaError) {
	r := [][]symbol.Token{}
	tempIds := 0
  
  // NEXT: Write some more tests

	for len(a) > 0 {
		if len(a) == 1 && a[0].Type == symbol.TEMP_IDENTIFIER {
			break
		}

		l := len(a) - 1

		o := rIndexOf(a, l, symbol.CURVED_BRACE_OPEN)
		if o == -1 {
			c := indexOf(a, 0, symbol.CURVED_BRACE_CLOSE)

			if c == -1 {
				r = append(r, a)
				break
			} else {
				m := "Missing closing brace to corresponding opening one"
				err := NewAnaError(m, c)
				return nil, err
			}
		}

		c := indexOf(a, o, symbol.CURVED_BRACE_CLOSE)
		if c == -1 {
			m := "Didn't expect to find a closing brace without a corresponding opening one"
			err := NewAnaError(m, c)
			return nil, err
		}

		s := a[o+1 : c]
		r = append(r, s)

		start := a[:o]
		tempIds++
		mid := symbol.Token{
			Val:  `#` + strconv.Itoa(tempIds),
			Type: symbol.TEMP_IDENTIFIER,
		}
		end := a[c+1:]

		a = append(start, mid)
		a = append(a, end...)

		fmt.Println(a)
	}

	return r, nil
}

// indexOf returns the next index of the symbol with the specified type
// of -1 if no matching token is found.
func indexOf(a []symbol.Token, start int, t symbol.SymbolType) int {
	l := len(a)
	for i := start; i < l; i++ {
		if i < start {
			continue
		}

		if a[i].Type == t {
			return i
		}
	}

	return -1
}

// rIndexOf returns the index of the last token with the specified type
// of -1 if no matching token is found. 'start' determines where to start
// searching back from, anything after will not be searched.
func rIndexOf(a []symbol.Token, start int, t symbol.SymbolType) int {
	for i := start; i > -1; i-- {
		if a[i].Type == t {
			return i
		}
	}

	return -1
}
