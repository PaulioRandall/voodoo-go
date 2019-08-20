package scanner_new

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanNumber scans symbols that start with a unicode category Nd rune returning
// a literal number token.
func scanNumber(r *Runer) (*token.Token, ParseToken, *token.Token) {
	start := r.NextCol()

	sig, err := scanSignificant(r)
	if err != nil {
		return nil, nil, err
	}

	frac, err := scanFractional(r)
	if err != nil {
		return nil, nil, err
	}

	s := sig + frac
	tk := numberToken(r, start, s)

	return ScanNext(r, tk)
}

// numberToken creates a new number token.
func numberToken(r *Runer, start int, v string) *token.Token {
	return &token.Token{
		Val:   v,
		Line:  r.Line(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_NUMBER,
	}
}

// scanFractional scans the fractional part of a number including the decimal
// separator.
func scanFractional(r *Runer) (string, *token.Token) {
	ru, _, err := r.LookAhead()
	if err != nil {
		return ``, runerErrorToken(r, err)
	}

	if !isDecimalDelim(ru) {
		return ``, nil
	}

	r.SkipRune()
	return scanFractionalDigits(r, string(ru))
}

// scanFractionalDigits scans the digits of the fractional part of a number
// returning the full fractional part.
func scanFractionalDigits(r *Runer, delim string) (string, *token.Token) {
	n, err := scanInt(r)
	if err != nil {
		return ``, runerErrorToken(r, err)
	}

	if len(n) == 0 || onlyContainsUnderscores(n) {
		return ``, badFractionalToken(r)
	}

	return delim + n, nil
}

// badFractionalToken creates a new error token for invalid fractional syntax.
func badFractionalToken(r *Runer) *token.Token {
	return &token.Token{
		Line: r.Line(),
		End:  r.NextCol(),
		Type: token.TT_ERROR_UPSTREAM,
		Errors: []string{
			"Invalid number format, either...",
			" - fractional digits are missing",
			" - or the decimal separator is a typo",
		},
	}
}

// onlyContainsUnderscores returns true if the string only contains underscores.
func onlyContainsUnderscores(s string) bool {
	for _, ru := range s {
		if !isUnderscore(ru) {
			return false
		}
	}
	return true
}

// scanSignificant scans the significant part of a number. E.g. `123` from the
// number `123.456`.
func scanSignificant(r *Runer) (string, *token.Token) {
	head, err := r.ReadRune()
	if err != nil {
		return ``, runerErrorToken(r, err)
	}

	tail, err := scanInt(r)
	if err != nil {
		return ``, runerErrorToken(r, err)
	}

	return string(head) + tail, nil
}

// scanInt scans all digits that make up an integer. Underscores are allowed as
// separators.
func scanInt(r *Runer) (string, error) {
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return ``, err
		}

		if !isDigit(ru) && !isUnderscore(ru) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return sb.String(), nil
}
