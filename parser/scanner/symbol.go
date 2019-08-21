package scanner

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSymbol scans one or two runes returning one of:
// - comparison operator token
// - arithmetic operator token
// - logical operator token
// - opening or closing brace token
// - value separator token
// - void token
func scanSymbol(r *Runer) (*token.Token, ParseToken, *token.Token) {

	ru1, ru2, err := r.LookAhead()
	if err != nil {
		return nil, nil, runerErrorToken(r, err)
	}

	var tk *token.Token

	switch {
	case ru1 == '<' && ru2 == '-':
		tk = symbolToken(r, token.TT_ASSIGN, 2)
	case ru1 == '<' && ru2 == '=':
		tk = symbolToken(r, token.TT_CMP_LT_OR_EQ, 2)
	case ru1 == '<':
		tk = symbolToken(r, token.TT_CMP_LT, 1)
	case ru1 == '>' && ru2 == '=':
		tk = symbolToken(r, token.TT_CMP_MT_OR_EQ, 2)
	case ru1 == '>':
		tk = symbolToken(r, token.TT_CMP_MT, 1)
	case ru1 == '=' && ru2 == '=':
		tk = symbolToken(r, token.TT_CMP_EQ, 2)
	case ru1 == '!' && ru2 == '=':
		tk = symbolToken(r, token.TT_CMP_NOT_EQ, 2)
	case ru1 == '=' && ru2 == '>':
		tk = symbolToken(r, token.TT_MATCH, 2)
	case ru1 == '!':
		tk = symbolToken(r, token.TT_NOT, 1)
	case ru1 == '|' && ru2 == '|':
		tk = symbolToken(r, token.TT_OR, 2)
	case ru1 == '&' && ru2 == '&':
		tk = symbolToken(r, token.TT_AND, 2)
	case ru1 == '+':
		tk = symbolToken(r, token.TT_ADD, 1)
	case ru1 == '-':
		tk = symbolToken(r, token.TT_SUBTRACT, 1)
	case ru1 == '*':
		tk = symbolToken(r, token.TT_MULTIPLY, 1)
	case ru1 == '/':
		tk = symbolToken(r, token.TT_DIVIDE, 1)
	case ru1 == '%':
		tk = symbolToken(r, token.TT_MODULO, 1)
	case ru1 == '(':
		tk = symbolToken(r, token.TT_CURVY_OPEN, 1)
	case ru1 == ')':
		tk = symbolToken(r, token.TT_CURVY_CLOSE, 1)
	case ru1 == '[':
		tk = symbolToken(r, token.TT_SQUARE_OPEN, 1)
	case ru1 == ']':
		tk = symbolToken(r, token.TT_SQUARE_CLOSE, 1)
	case ru1 == ',':
		tk = symbolToken(r, token.TT_VALUE_DELIM, 1)
	case ru1 == '_':
		tk = symbolToken(r, token.TT_VOID, 1)
	default:
		return nil, nil, unknownSymbol(r, ru1)
	}

	return scanNext(r, tk)
}

// symbolToken creates the new token when a symbol match is found.
func symbolToken(r *Runer, t token.TokenType, runes int) *token.Token {

	ru, _ := r.ReadRune()
	s := string(ru)

	if runes == 2 {
		ru, _ = r.ReadRune()
		s += string(ru)
	}

	return &token.Token{
		Val:   s,
		Start: r.NextCol() - runes,
		End:   r.NextCol(),
		Type:  t,
	}
}

// unknownSymbol creates a fault for when a symbol is not known.
func unknownSymbol(r *Runer, ru rune) *token.Token {
	return &token.Token{
		Line: r.Line(),
		End:  r.NextCol(),
		Type: token.TT_ERROR_UPSTREAM,
		Errors: []string{
			"I don't know what this symbol means '" + string(ru) + "'",
		},
	}
}
