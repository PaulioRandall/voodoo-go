package new_parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses a token array into a parse tree.
func Parse(in []Token) (*Expression, Fault) {

	return nil, nil
}

// containsAssignment returns true if the input token array results
// in an assignment but not a match.
func containsAssignment(in []Token) bool {
	for _, tk := range in {
		if tk.Type == token.ASSIGNMENT {
			return true
		}

		if tk.Type == token.LOGICAL_MATCH {
			return false
		}
	}

	return false
}
