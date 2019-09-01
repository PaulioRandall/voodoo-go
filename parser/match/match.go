package match

import (
	"errors"

	"github.com/PaulioRandall/voodoo-go/parser/stat"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Match creates a statement from a token array by matching the pattern of token
// kinds to the space of all allowable patterns.
// E.g. `x <- 1` matches `ID, ASSIGN, EXPRESSION` pattern.
func Match(in []token.Token) (stat.Statement, error) {
	if i := indexOf(in, token.TT_ASSIGN); i != -1 {
		s := stat.New(stat.SK_EXPRESSION, in[0:i], in[i+1:len(in)])
		return s, nil
	}

	return stat.Empty(), errors.New(`No matching statement pattern`)
}

// indexOf returns the index of token with the specified kind or -1 if no
// token matching the kind could be found.
func indexOf(in []token.Token, k token.Kind) int {
	for i, v := range in {
		if v.Kind() == k {
			return i
		}
	}
	return -1
}
