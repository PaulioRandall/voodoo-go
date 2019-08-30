package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// scanShebang scans a shebang line.
func scanShebang(r *runer.Runer) (token.Token, ScanError) {
	start := r.NextCol()
	sb := strings.Builder{}

	for ru, end, err := shebangRead(r); !end; ru, end, err = shebangRead(r) {
		if err != nil {
			return nil, err
		}

		sb.WriteRune(ru)
	}

	return newShebangToken(r, sb.String(), start), nil
}

// shebangRead peeks at the next rune to see if it forms part of the shebang
// line returning it if so. If not -1 is returned.
func shebangRead(r *runer.Runer) (rune, bool, ScanError) {
	switch ru, eof, err := r.Peek(); {
	case err != nil:
		return 0, false, runerError(r, err)
	case eof, ru == '\n':
		return 0, true, nil
	default:
		if _, err = r.Skip(); err != nil {
			return 0, false, runerError(r, err)
		}
		return ru, false, nil
	}
}

// newShebangToken returns a new SHEBANG token.
func newShebangToken(r *runer.Runer, text string, start int) token.Token {
	return scanTok{
		text:  text,
		line:  r.Line(),
		start: start,
		end:   r.NextCol(),
		kind:  token.TT_SHEBANG,
	}
}
