package scanner

import (
	"strings"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// scanNumber scans symbols that start with a unicode category Nd rune returning
// a literal number token.
func scanNumber(r *runer.Runer) (token.Token, ScanError) {
	start := r.NextCol()

	sig, err := scanInt(r)
	if err != nil {
		return nil, err
	}

	frac, err := scanFractional(r)
	if err != nil {
		return nil, err
	}

	s := sig + frac
	tk := numberToken(r, start, s)

	return tk, nil
}

// scanInt scans all digits that make up an integer.
func scanInt(r *runer.Runer) (string, ScanError) {
	sb := strings.Builder{}

	for {
		ru, eof, err := r.Peek()
		if err != nil {
			return ``, runerError(r, err)
		}

		if eof || !unicode.IsDigit(ru) {
			break
		}

		if _, err = r.Skip(); err != nil {
			return ``, runerError(r, err)
		}

		sb.WriteRune(ru)
	}

	return sb.String(), nil
}

// scanFractional scans the fractional part of a number including the decimal
// separator.
func scanFractional(r *runer.Runer) (string, ScanError) {
	ru, eof, err := r.Peek()
	if err != nil {
		return ``, runerError(r, err)
	}

	if eof || ru != '.' {
		return ``, nil
	}

	if _, err = r.Skip(); err != nil {
		return ``, runerError(r, err)
	}

	n, scErr := scanInt(r)
	if scErr != nil {
		return ``, scErr
	}

	if len(n) == 0 {
		return ``, badFractionalToken(r)
	}

	return string(ru) + n, nil
}

// badFractionalToken creates a new scan error for invalid fractional syntax.
func badFractionalToken(r *runer.Runer) ScanError {
	return scanErr{
		l: r.Line(),
		i: r.NextCol(),
		e: []string{
			"Invalid number format, either...",
			"- fractional digits are missing",
			"- or the decimal separator is a typo",
		},
	}
}

// numberToken creates a new number token.
func numberToken(r *runer.Runer, start int, text string) token.Token {
	return scanTok{
		text:  text,
		line:  r.Line(),
		start: start,
		end:   r.NextCol(),
		kind:  token.TT_NUMBER,
	}
}
