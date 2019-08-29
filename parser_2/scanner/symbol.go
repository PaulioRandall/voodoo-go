package scanner

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// scanSymbol scans one or two runes returning one of:
// - comparison operator token
// - arithmetic operator token
// - logical operator token
// - opening or closing brace token
// - value separator token
// - void token
func scanSymbol(r *runer.Runer) (token.Token, ScanError) {

	ru1, _, err := r.Read()
	if err != nil {
		return nil, runerError(r, err)
	}

	ru2, _, err := r.Peek()
	if err != nil {
		return nil, runerError(r, err)
	}

	switch {
	case ru1 == '<' && ru2 == '-':
		return symbolToken(r, ru1, ru2, token.TT_ASSIGN)
	case ru1 == '=':
		return symbolToken(r, ru1, -1, token.TT_ASSIGN)
	case ru1 == ':' && ru2 == '=':
		return symbolToken(r, ru1, ru2, token.TT_ASSIGN)
		/*
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
				tk = symbolToken(r, token.TT_IF_THEN, 2)
			case ru1 == '!':
				tk = symbolToken(r, token.TT_NOT, 1)
			case ru1 == '|' && ru2 == '|':
				tk = symbolToken(r, token.TT_OR, 2)
			case ru1 == '&' && ru2 == '&':
				tk = symbolToken(r, token.TT_AND, 2)
		*/
	case ru1 == '+':
		return symbolToken(r, ru1, -1, token.TT_ADD)
	case ru1 == '-':
		return symbolToken(r, ru1, -1, token.TT_SUBTRACT)
	case ru1 == '*':
		return symbolToken(r, ru1, -1, token.TT_MULTIPLY)
	case ru1 == '/':
		return symbolToken(r, ru1, -1, token.TT_DIVIDE)
	case ru1 == '%':
		return symbolToken(r, ru1, -1, token.TT_MODULO)
		/*
			case ru1 == '{':
				tk = symbolToken(r, token.TT_CURLY_OPEN, 1)
			case ru1 == '}':
				tk = symbolToken(r, token.TT_CURLY_CLOSE, 1)
			case ru1 == '(':
				tk = symbolToken(r, token.TT_CURVED_OPEN, 1)
			case ru1 == ')':
				tk = symbolToken(r, token.TT_CURVED_CLOSE, 1)
			case ru1 == '[':
				tk = symbolToken(r, token.TT_SQUARE_OPEN, 1)
			case ru1 == ']':
				tk = symbolToken(r, token.TT_SQUARE_CLOSE, 1)
			case ru1 == ',':
				tk = symbolToken(r, token.TT_VALUE_DELIM, 1)
			case ru1 == '_':
				tk = symbolToken(r, token.TT_VOID, 1)
		*/
	default:
		return nil, unknownSymbol(r, ru1)
	}
}

// symbolToken creates the new token when a symbol match is found.
func symbolToken(r *runer.Runer, ru1, ru2 rune, k token.Kind) (token.Token, ScanError) {

	text, size, err := runesToText(r, ru1, ru2)
	if err != nil {
		return nil, err
	}

	tk := scanTok{
		text:  text,
		line:  r.Line(),
		start: r.NextCol() - size,
		end:   r.NextCol(),
		kind:  k,
	}

	return tk, nil
}

// runesToText converts the runes to a string and skips the next rune in the
// reader if the second rune formed part of the symbol.
func runesToText(r *runer.Runer, ru1, ru2 rune) (string, int, ScanError) {
	if ru2 < 0 {
		return string(ru1), 1, nil
	}

	if _, err := r.Skip(); err != nil {
		return ``, 0, runerError(r, err)
	}

	text := string(ru1) + string(ru2)
	return text, 2, nil
}

// unknownSymbol creates a ScanError for when a symbol is not recognised.
func unknownSymbol(r *runer.Runer, ru rune) ScanError {
	return scanErr{
		l: r.Line(),
		i: r.NextCol(),
		e: []string{
			"I don't know what this symbol means '" + string(ru) + "'",
		},
	}
}
