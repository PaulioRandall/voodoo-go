package scanner

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSymbol scans one or two runes returning one of:
// - comparison operator
// - arithmetic operator
// - logical operator
// - opening or closing brace or bracket
// - value or key-value separator
// - void identifier
// - number range generator
func scanSymbol(r *Runer) token.Token {

	ru1, ru2, err := r.LookAhead()
	if err != nil {
		start := r.Col() + 1
		return errorToToken(r, start, err)
	}

	switch {
	case cmpPair(ru1, ru2, '<', '-'):
		return onMatch(r, token.TT_ASSIGN, 2)
	case cmpPair(ru1, ru2, '<', '='):
		return onMatch(r, token.TT_CMP_LT_OR_EQ, 2)
	case cmp(ru1, '<'):
		return onMatch(r, token.TT_CMP_LT, 1)
	case cmpPair(ru1, ru2, '>', '='):
		return onMatch(r, token.TT_CMP_MT_OR_EQ, 2)
	case cmp(ru1, '>'):
		return onMatch(r, token.TT_CMP_MT, 1)
	case cmpPair(ru1, ru2, '=', '='):
		return onMatch(r, token.TT_CMP_EQ, 2)
	case cmpPair(ru1, ru2, '!', '='):
		return onMatch(r, token.TT_CMP_NOT_EQ, 2)
	case cmpPair(ru1, ru2, '=', '>'):
		return onMatch(r, token.TT_MATCH, 2)
	case cmp(ru1, '!'):
		return onMatch(r, token.TT_NOT, 1)
	case cmpPair(ru1, ru2, '|', '|'):
		return onMatch(r, token.TT_OR, 2)
	case cmpPair(ru1, ru2, '&', '&'):
		return onMatch(r, token.TT_AND, 2)
	case cmp(ru1, '+'):
		return onMatch(r, token.TT_ADD, 1)
	case cmp(ru1, '-'):
		return onMatch(r, token.TT_SUBTRACT, 1)
	case cmp(ru1, '*'):
		return onMatch(r, token.TT_MULTIPLY, 1)
	case cmp(ru1, '/'):
		return onMatch(r, token.TT_DIVIDE, 1)
	case cmp(ru1, '%'):
		return onMatch(r, token.TT_MODULO, 1)
	case cmp(ru1, '('):
		return onMatch(r, token.TT_CURVY_OPEN, 1)
	case cmp(ru1, ')'):
		return onMatch(r, token.TT_CURVY_CLOSE, 1)
	case cmp(ru1, '['):
		return onMatch(r, token.TT_SQUARE_OPEN, 1)
	case cmp(ru1, ']'):
		return onMatch(r, token.TT_SQUARE_CLOSE, 1)
	case cmp(ru1, ','):
		return onMatch(r, token.TT_VALUE_DELIM, 1)
	case cmp(ru1, '_'):
		return onMatch(r, token.TT_VOID, 1)
	}

	return unknownNonTerminal(ru1, r.Line(), r.Col()+2)
}

// cmpPair compares the first terminal with the third and the second with the
// last, if both sets match true is returned.
func cmpPair(a1, b1, a2, b2 rune) bool {
	return a1 == a2 && b1 == b2
}

// cmp compares the two terminals.
func cmp(a, b rune) bool {
	return a == b
}

// onMatch creates the new token when a symbol match is found.
func onMatch(r *Runer, t token.TokenType, count int) token.Token {

	ru, _ := r.ReadRune()
	s := string(ru)

	if count == 2 {
		ru, _ = r.ReadRune()
		s += string(ru)
	}

	tk := token.Token{
		Val:   s,
		Start: r.Col() - count + 1,
		End:   r.Col() + 1,
		Type:  t,
	}

	return tk
}

// unknownNonTerminal creates a fault for when a symbol is not known.
func unknownNonTerminal(ru rune, line, col int) token.Token {
	return token.Token{
		Line:  line,
		Start: col - 1,
		End:   col,
		Type:  token.TT_ERROR_UPSTREAM,
		Errors: []string{
			"I don't know what this symbol means '" + string(ru) + "'",
		},
	}
}
