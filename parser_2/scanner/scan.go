package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// Scanner represents the scanning part of a lexical analyser.
type Scanner struct {
	count int
}

// ScanToken represents a function that scans a Token from a Runer. The scanning
// employs longest match against the productions relevant to the specific Kind
// of Token.
type ScanToken func(*runer.Runer) (token.Token, ScanError)

// Next returns a suitable function that will scan the next Token.
func (s *Scanner) Next(r *runer.Runer) (ScanToken, ScanError) {
	if s.count == 0 {
		return scanShebang, nil
	}

	return nil, nil
}

// scanShebang scans a shebang line.
func scanShebang(r *runer.Runer) (token.Token, ScanError) {
	start := r.NextCol()
	sb := strings.Builder{}

	for {
		ru, eof, err := r.Peek()
		if err != nil {
			return nil, runerError(r, err)
		}

		if eof || ru == '\n' {
			break
		}

		if _, err = r.Skip(); err != nil {
			return nil, runerError(r, err)
		}
		sb.WriteRune(ru)
	}

	tk := scanTok{
		text:  sb.String(),
		line:  r.Line(),
		start: start,
		end:   r.NextCol(),
		kind:  token.TT_SHEBANG,
	}

	return tk, nil
}
