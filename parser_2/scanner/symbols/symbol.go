package symbols

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// scanSymbol scans one or two runes returning a non-alphanumeric symbol set.
func ScanSymbol(r *runer.Runer) (token.Token, err.ScanError) {

	ru1, _, e := r.Peek()
	if e != nil {
		return nil, err.NewByRuner(r, e)
	}

	/*
		ru2, _, e := r.PeekMore()
		if e != nil {
			return nil, err.NewByRuner(r, e)
		}
	*/
	switch {
	case ru1 == ':':
		return symbolToken(r, 1, token.TT_ASSIGN)
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
	default:
		return nil, unknownSymbol(r, ru1)
	}
}

// symbolToken creates the new token when a symbol match is found.
func symbolToken(r *runer.Runer, n int, k token.Kind) (token.Token, err.ScanError) {

	text, size, e := runesToText(r, n)
	if e != nil {
		return nil, e
	}

	tk := scantok.New(
		text,
		r.Line(),
		r.NextCol()-size,
		r.NextCol(),
		k,
	)

	return tk, nil
}

// runesToText converts the specified number of runes to a string.
func runesToText(r *runer.Runer, n int) (string, int, err.ScanError) {

	ru1, _, e := r.Read()
	if e != nil {
		return ``, 0, err.NewByRuner(r, e)
	}

	if n == 2 {
		ru2, _, e := r.Read()
		if e != nil {
			return ``, 0, err.NewByRuner(r, e)
		}

		return string(ru1) + string(ru2), 2, nil
	}

	return string(ru1), 1, nil
}

// unknownSymbol creates a ScanError for when a symbol is not recognised.
func unknownSymbol(r *runer.Runer, ru rune) err.ScanError {
	return err.New(
		r.Line(),
		r.NextCol(),
		[]string{
			"I don't know what this symbol means '" + string(ru) + "'",
		},
	)
}
