package scanner

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ScanToken represents a function that scans a Token from a Runer. The scanning
// employs longest match against the productions relevant to the specific Kind
// of Token.
type ScanToken func(*runer.Runer) (token.Token, ScanError)

// Scanner represents the scanning part of a lexical analyser. If false the
// scanner will return a SheBang line token canning function on the first
// invocation of Next().
type Scanner bool

// Next returns a suitable function that will scan the next Token.
func (s *Scanner) Next(r *runer.Runer) (ScanToken, ScanError) {
	if !*s {
		*s = true
		return scanShebang, nil
	}

	switch ru, eof, err := r.Peek(); {
	case err != nil:
		return nil, runerError(r, err)
	case eof:
		return nil, nil
	case unicode.IsDigit(ru):
		return scanNumber, nil
	case unicode.IsLetter(ru):
		return scanWord, nil
	default:
		return scanSymbol, nil
	}
}
