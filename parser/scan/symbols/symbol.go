package symbols

import (
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSymbol scans one or two runes returning a non-alphanumeric symbol set.
func ScanSymbol(r *runer.Runer) (token.Token, perror.Perror) {

	ru1, _, e := r.Peek()
	if e != nil {
		return nil, e
	}

	ru2, _, e := r.PeekMore()
	if e != nil {
		return nil, e
	}

	switch {
	case ru1 == '<' && ru2 == '-':
		return symbolToken(r, 2, token.TT_ASSIGN)
	case ru1 == ':' && ru2 == '=':
		return symbolToken(r, 2, token.TT_ASSIGN)
	case ru1 == '+':
		return symbolToken(r, 1, token.TT_ADD)
	case ru1 == '-':
		return symbolToken(r, 1, token.TT_SUBTRACT)
	case ru1 == '*':
		return symbolToken(r, 1, token.TT_MULTIPLY)
	case ru1 == '/':
		return symbolToken(r, 1, token.TT_DIVIDE)
	case ru1 == '%':
		return symbolToken(r, 1, token.TT_MODULO)
	case ru1 == '_':
		return symbolToken(r, 1, token.TT_VOID)
	case ru1 == ',':
		return symbolToken(r, 1, token.TT_DELIM)
	default:
		return nil, unknownSymbol(r, ru1)
	}
}

// symbolToken creates the new token when a symbol match is found.
func symbolToken(r *runer.Runer, n int, k token.Kind) (token.Token, perror.Perror) {

	text, size, e := runesToText(r, n)
	if e != nil {
		return nil, e
	}

	tk := token.New(
		text,
		r.Line(),
		r.NextCol()-size,
		r.NextCol(),
		k,
	)

	return tk, nil
}

// runesToText converts the specified number of runes to a string.
func runesToText(r *runer.Runer, n int) (string, int, perror.Perror) {

	ru1, _, e := r.Read()
	if e != nil {
		return ``, 0, e
	}

	if n == 2 {
		ru2, _, e := r.Read()
		if e != nil {
			return ``, 0, e
		}

		return string(ru1) + string(ru2), 2, nil
	}

	return string(ru1), 1, nil
}

// unknownSymbol creates a ScanError for when a symbol is not recognised.
func unknownSymbol(r *runer.Runer, ru rune) perror.Perror {
	return perror.New(
		r.Line(),
		r.NextCol(),
		[]string{
			"I don't know what this symbol means '" + string(ru) + "'",
		},
	)
}
