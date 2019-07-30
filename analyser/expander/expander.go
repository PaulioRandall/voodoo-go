package expander

import (
	"strconv"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// ExpandParens expands all expression parentheses into separate assignment
// statements. These are then compiled into a correctly ordered set of token
// arrays.
func ExpandParens(outer []token.Token) ([][]token.Token, fault.Fault) {
	r := [][]token.Token{}
	var inner []token.Token
	var err fault.Fault

	for id := 1; outer != nil; id++ {
		outer, inner, err = expandParen(outer, id)

		if err != nil {
			return nil, err
		}

		r = append(r, inner)
	}

	return r, nil
}

// expandParen finds one set of parenthesss that do not contain parenthesss
// themselves and extracts it, removing the parenthesss. An implicit
// identifier is inserted into the original (outer) token array to represent
// the result of the extracted expression (inner). The inner is prefixed with
// two tokens that assign to the implicit identifier.
//
// This function essential breaks up a large expressions into two smaller
// ones as defined by the coder. Blocks of parentheses `()` are used by the
// coder determine how the larger expression should be broken up. By
// plugging the `outer` value back into the function the expression can be
// broken up further until `outer` is nil. At which point there are no
// parentheses left to expand.
func expandParen(in []token.Token, nextTempId int) (outer []token.Token, inner []token.Token, err fault.Fault) {

	a, z := findParen(in)
	if a == -1 && z == -1 {
		inner = in
		return
	}

	err = checkParenIndexes(in, a, z)
	if err != nil {
		return
	}

	outer, inner = sliceOutParen(in, a, z, nextTempId)
	return
}

// sliceOutParen slices out the inner expression from the outer leaving an implicit
// identifier in its place. The inner is prefixed with two tokens that assign the
// the implicit identifier with the value of the inner expression.
func sliceOutParen(in []token.Token, a, z, nextTempId int) (outer []token.Token, inner []token.Token) {
	id := newImplicitIdToken(nextTempId)

	inner = []token.Token{
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
func checkParenIndexes(in []token.Token, a, z int) (err fault.Fault) {
	if a == -1 {
		m := "Didn't expect to find a closing parenthesis without a corresponding opening one"
		err = fault.Paren(m)
		err = fault.SetTo(err, in[z].Start)
	} else if z == -1 {
		m := "Didn't expect to find an opening parenthesis without a corresponding closing one"
		err = fault.Paren(m)
		err = fault.SetFrom(err, in[a].Start)
	}

	return
}

// newImplicitIdToken returns a new and unique implicit identifier token.
func newImplicitIdToken(id int) token.Token {
	return token.Token{
		Val:  `#` + strconv.Itoa(id),
		Type: token.IDENTIFIER_IMPLICIT,
	}
}

// newAssignToken returns a new assignment token.
func newAssignToken() token.Token {
	return token.Token{
		Val:  `<-`,
		Type: token.ASSIGNMENT,
	}
}

// findParen finds a pair of matching parenthesis that do not contain
// parentheses themselves.
func findParen(in []token.Token) (a int, z int) {
	l := len(in) - 1
	a = rIndexOf(in, l, token.PAREN_CURVY_OPEN)
	if a == -1 {
		z = indexOf(in, 0, token.PAREN_CURVY_CLOSE)
	} else {
		z = indexOf(in, a, token.PAREN_CURVY_CLOSE)
	}
	return
}

// containsType returns true if the token array contains a token with
// one of the specified symbol types.
func containsType(a []token.Token, t ...token.TokenType) bool {
	for _, v := range a {
		for _, ty := range t {
			if v.Type == ty {
				return true
			}
		}
	}
	return false
}

// indexOf returns the next index of the symbol with the specified type
// of -1 if no matching token is found.
func indexOf(a []token.Token, start int, t token.TokenType) int {
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
func rIndexOf(a []token.Token, start int, t token.TokenType) int {
	for i := start; i > -1; i-- {
		if a[i].Type == t {
			return i
		}
	}

	return -1
}
