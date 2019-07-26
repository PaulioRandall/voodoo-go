package analyser

import (
	//"fmt"
	//"strconv"

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
func Analyse(a []symbol.Token) (operation.InstructionSet, AnaError) {
	for {
		var b []symbol.Token
		var err AnaError

		if !containsType(a, symbol.CURVED_BRACE_OPEN, symbol.CURVED_BRACE_CLOSE) {
			break
		}

		a, b, err := expandExpr(a)

		if err != nil {
			return nil, err
		}

		if a == nil || b == nil { // REMOVE
			return nil, nil
		}
		// ... add b to the end of the instruction set
	}

	// ... add remaining to the end of the instruction set
	return nil, nil
}

// expandExpr finds one set of parenthesis that do not contain parenthesis
// themselves and extracts it, removing the parenthesis. An identifier is
// inserted into the original (outer) token array to represent the result
// of the extracted expression (inner). The inner is prefixed with an
// assignment operation to the identifier. The outer is returned as the
// first result, inner second.
func expandExpr(in []symbol.Token) (outer []symbol.Token, inner []symbol.Token, err AnaError) {

	a, z := findParenPair(in)
	err = checkParenIndexes(in, a, z)
	if err != nil {
		return
	}

	id := newTempIdToken()

	inner = []symbol.Token{
		id,
		newAssignToken(),
	}
	inner = append(inner, in[a+1:z]...)

	outer = in[:a]
	outer = append(outer, id)
	outer = append(outer, in[z+1:]...)

	return
}

// checkParenIndexes checks the results of findParenPair() and returns
// a non-nil error if they are invalid.
func checkParenIndexes(in []symbol.Token, a, b int) (err AnaError) {
	switch {
	case a == -1 && b == -1:
		m := "[BUG] This function may only be called if parenthesis are present within the input"
		err = NewAnaError(m, -1)
	case a == -1:
		m := "Didn't expect to find a closing parenthesis without a corresponding opening one"
		err = NewAnaError(m, in[b].Start)
	case b == -1:
		m := "Didn't expect to find an opening parenthesis without a corresponding closing one"
		err = NewAnaError(m, in[a].Start)
	}

	return
}

// newTempIdToken returns a new and unique temporary identifier token.
func newTempIdToken() symbol.Token {
	return symbol.Token{
		Val:  `#1`,
		Type: symbol.TEMP_IDENTIFIER,
	}
}

// newAssignToken returns a new assignment token.
func newAssignToken() symbol.Token {
	return symbol.Token{
		Val:  `<-`,
		Type: symbol.ASSIGNMENT,
	}
}

// findParenPair finds a pair of matching parenthesis that do not contain
// parenthesis themselves.
func findParenPair(in []symbol.Token) (a int, z int) {
	l := len(in) - 1
	a = rIndexOf(in, l, symbol.CURVED_BRACE_OPEN)
	if a == -1 {
		z = indexOf(in, 0, symbol.CURVED_BRACE_CLOSE)
	} else {
		z = indexOf(in, a, symbol.CURVED_BRACE_CLOSE)
	}
	return
}

// containsType returns true if the token array contains a token with
// one of the specified symbol types.
func containsType(a []symbol.Token, t ...symbol.SymbolType) bool {
	for _, v := range a {
		for _, ty := range t {
			if v.Type == ty {
				return true
			}
		}
	}
	return false
}

// expandExpr finds one set of parenthesis that do not contain parenthesis
// themselves and extracts it, removing the parenthesis. An identifier is
// inserted into the original (outer) token array to represent the result
// of the extracted expression (inner). The inner is prefixed with an
// assignment operation to the identifier. The outer is returned as the
// first result, inner second.
/*
func expandParen(a []symbol.Token) (outer []symbol.Token, inner []symbol.Token, err AnaError) {
	tempIds := 0

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
*/

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
