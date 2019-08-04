package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSymbol scans one or two rune symbols returning one of:
// - comparison operator
// - arithmetic operator
// - logical operator
// - opening or closing brace or bracket
// - value or key-value separator
// - void identifier
// - number range generator
//
// Assumes that the input is not empty.
func scanSymbol(in []rune, col int) (*token.Token, []rune, fault.Fault) {

	switch {
	case startsWith(in, `<-`):
		return onMatchedSymbol(in, token.ASSIGNMENT, 2, col)
	case startsWith(in, `<=`):
		return onMatchedSymbol(in, token.CMP_LESS_THAN_OR_EQUAL, 2, col)
	case startsWith(in, `<`):
		return onMatchedSymbol(in, token.CMP_LESS_THAN, 1, col)
	case startsWith(in, `>=`):
		return onMatchedSymbol(in, token.CMP_GREATER_THAN_OR_EQUAL, 2, col)
	case startsWith(in, `>`):
		return onMatchedSymbol(in, token.CMP_GREATER_THAN, 1, col)
	case startsWith(in, `==`):
		return onMatchedSymbol(in, token.CMP_EQUAL, 2, col)
	case startsWith(in, `!=`):
		return onMatchedSymbol(in, token.CMP_NOT_EQUAL, 2, col)
	case startsWith(in, `=>`):
		return onMatchedSymbol(in, token.LOGICAL_MATCH, 2, col)
	case startsWith(in, `!`):
		return onMatchedSymbol(in, token.LOGICAL_NOT, 1, col)
	case startsWith(in, `||`):
		return onMatchedSymbol(in, token.LOGICAL_OR, 2, col)
	case startsWith(in, `&&`):
		return onMatchedSymbol(in, token.LOGICAL_AND, 2, col)
	case startsWith(in, `+`):
		return onMatchedSymbol(in, token.CALC_ADD, 1, col)
	case startsWith(in, `-`):
		return onMatchedSymbol(in, token.CALC_SUBTRACT, 1, col)
	case startsWith(in, `*`):
		return onMatchedSymbol(in, token.CALC_MULTIPLY, 1, col)
	case startsWith(in, `/`):
		return onMatchedSymbol(in, token.CALC_DIVIDE, 1, col)
	case startsWith(in, `%`):
		return onMatchedSymbol(in, token.CALC_MODULO, 1, col)
	case startsWith(in, `(`):
		return onMatchedSymbol(in, token.PAREN_CURVY_OPEN, 1, col)
	case startsWith(in, `)`):
		return onMatchedSymbol(in, token.PAREN_CURVY_CLOSE, 1, col)
	case startsWith(in, `[`):
		return onMatchedSymbol(in, token.PAREN_SQUARE_OPEN, 1, col)
	case startsWith(in, `]`):
		return onMatchedSymbol(in, token.PAREN_SQUARE_CLOSE, 1, col)
	case startsWith(in, `,`):
		return onMatchedSymbol(in, token.SEPARATOR_VALUE, 1, col)
	case startsWith(in, `_`):
		return onMatchedSymbol(in, token.VOID, 1, col)
	default:
		return nil, nil, unknownSymbol(in[0], col+1)
	}
}

// onMatchedSymbol creates the new token when a symbol match is found.
func onMatchedSymbol(in []rune, t token.TokenType, count, col int) (*token.Token, []rune, fault.Fault) {
	tk := &token.Token{
		Val:   string(in[:count]),
		Start: col,
		Type:  t,
	}

	return tk, in[count:], nil
}

// unknownSymbol creates a fault for when a symbol is not known.
func unknownSymbol(r rune, i int) fault.Fault {
	return fault.SyntaxFault{
		Index: i,
		Msgs: []string{
			"I don't know what this symbol means '" + string(r) + "'",
		},
	}
}
