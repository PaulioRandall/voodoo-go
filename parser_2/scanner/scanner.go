package scanner

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/number"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/shebang"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/symbols"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/word"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ScanToken represents a function that scans a Token from a Runer. The scanning
// employs longest match against the productions relevant to the specific Kind
// of Token.
type ScanToken func(*runer.Runer) (token.Token, err.ScanError)

// Scanner represents the scanning part of a lexical analyser. If false the
// scanner will return a SheBang line token canning function on the first
// invocation of Next().
type Scanner bool

// Next returns a suitable function that will scan the next Token.
func (s *Scanner) Next(r *runer.Runer) (ScanToken, err.ScanError) {
	if !*s {
		*s = true
		return shebang.ScanShebang, nil
	}

	switch ru, eof, e := r.Peek(); {
	case e != nil:
		return nil, err.NewByRuner(r, e)
	case eof:
		return nil, nil
	case unicode.IsDigit(ru):
		return number.ScanNumber, nil
	case unicode.IsLetter(ru):
		return word.ScanWord, nil
	default:
		return symbols.ScanSymbol, nil
	}
}
