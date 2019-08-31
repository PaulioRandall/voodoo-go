package scanner

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/number"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/shebang"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/space"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/symbols"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/word"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// TokenScanner represents a function that scans a Token from a Runer. The
// scanning employs longest match against the productions relevant to the
// specific Kind of Token.
type TokenScanner func(*runer.Runer) (token.Token, err.ScanError)

// Scanner represents the scanning part of a lexical analyser. If true the
// scanner will return a SheBang line token scanning function on the first
// invocation of Next().
type Scanner struct {
	shebang bool
}

// New returns a new Scanner.
func New(shebang bool) *Scanner {
	return &Scanner{
		shebang: shebang,
	}
}

// Next returns a suitable TokenScanner function that will scan the next Token.
func (s *Scanner) Next(r *runer.Runer) (TokenScanner, err.ScanError) {
	if s.shebang {
		s.shebang = false
		return shebang.ScanShebang, nil
	}

	switch ru, eof, e := r.Peek(); {
	case e != nil:
		return nil, err.NewByRuner(r, e)
	case eof:
		return nil, nil
	case unicode.IsSpace(ru):
		return space.ScanSpace, nil
	case unicode.IsLetter(ru):
		return word.ScanWord, nil
	case unicode.IsDigit(ru):
		return number.ScanNumber, nil
	default:
		return symbols.ScanSymbol, nil
	}
}
