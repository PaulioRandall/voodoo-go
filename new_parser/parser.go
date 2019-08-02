package new_parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses a token array into a statement.
func Parse(in []Token) (*Statement, Fault) {
	//_, in := splitAssignment(in)

	return nil, nil
}

// parseAssignment parses the assignment part of the
// token array to produce an expression
func parseAssignment(in []Token) (*Expression, Fault) {
	return nil, nil
}

// splitAssignment returns the assignment part of the
// token array or nil if there is no assignment part.
func splitAssignment(in []Token) ([]Token, []Token) {
	for i, tk := range in {
		if tk.Type == token.ASSIGNMENT {
			return in[:i], in[i:]
		}

		if tk.Type == token.LOGICAL_MATCH {
			return nil, in
		}
	}

	return nil, in
}
