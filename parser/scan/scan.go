package scan

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser/scan/err"
	"github.com/PaulioRandall/voodoo-go/parser/scan/number"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/scan/shebang"
	"github.com/PaulioRandall/voodoo-go/parser/scan/space"
	"github.com/PaulioRandall/voodoo-go/parser/scan/symbols"
	"github.com/PaulioRandall/voodoo-go/parser/scan/word"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// TokenScanner represents a function that scans a Token from a Runer. The
// scanning employs longest match against the productions relevant to the
// specific Kind of Token.
type TokenScanner func(*runer.Runer) (token.Token, err.ScanError)

// Scanner represents the scanning part of a lexical analyser.
type Scanner struct {
	shebang bool // True if the next line is a shebang line
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
