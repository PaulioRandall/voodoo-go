package new_parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses a token array into a statement.
func Parse(in []Token) (Statement, Fault) {
	//a, in := splitAssignment(in)

	return nil, nil
}

// parseAssignment parses the assignment part of a statment
// to produce an expression for the left side.
func parseAssignment(in []Token) (Expression, Fault) {

	idCount := len(in) / 2
	ids := make([]Token, idCount)

	for i, tk := range in {
		if isEven(i) {
			err := validateIdentifier(tk)
			if err != nil {
				return nil, err
			}

			ids[i/2] = tk
			continue
		}

		err := validateDelimiter(tk)
		if err != nil {
			return nil, err
		}
	}

	out := List{
		Tokens: ids,
	}

	return out, nil
}

// validateDelimiter validates the passed token is a value
// delimiter or assignment returning a fault if not.
func validateDelimiter(tk Token) Fault {
	if tk.Type != token.ASSIGNMENT && tk.Type != token.SEPARATOR_VALUE {
		return ParseFault{
			Msgs: []string{
				`Unexpected token type`,
				`Expected a value delimiter or assignment token`,
			},
		}
	}

	return nil
}

// validateIdentifier validates the passed token is an identifier
// returning a fault if not.
func validateIdentifier(tk Token) Fault {
	if tk.Type != token.IDENTIFIER {
		return ParseFault{
			Msgs: []string{
				`Can't assign to non-identifier`,
				`Expected an identifier`,
			},
		}
	}

	return nil
}

// isEven returns true if the input is odd.
func isEven(i int) bool {
	return i == 0 || i%2 == 0
}

// splitOnAssignment returns the assignment part of the
// token array or nil if there is no assignment part.
func splitOnAssignment(in []Token) ([]Token, []Token) {
	for i, tk := range in {
		if tk.Type == token.ASSIGNMENT {
			i++
			return in[:i], in[i:]
		}

		if tk.Type == token.LOGICAL_MATCH {
			return nil, in
		}
	}

	return nil, in
}
