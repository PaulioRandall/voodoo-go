package scan

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
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
type TokenScanner func(*runer.Runer) (token.Token, perror.Perror)

// ShebangScanner returns the scanner that will scan all remaning runes in the
// current line
func ShebangScanner() TokenScanner {
	return shebang.ScanShebang
}

// Next returns a suitable TokenScanner function that will scan the next Token.
func Next(r *runer.Runer) (TokenScanner, perror.Perror) {
	switch ru, eof, e := r.Peek(); {
	case e != nil:
		return nil, e
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
