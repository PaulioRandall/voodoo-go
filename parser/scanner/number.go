package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanNumber scans symbols that start with a unicode category Nd rune returning
// a literal number token.
func scanNumber(r *Runer) (*token.Token, ParseToken, *token.Token) {
	start := r.NextCol()

	sig, err := scanInt(r)
	if err != nil {
		return nil, nil, err
	}

	frac, err := scanFractional(r)
	if err != nil {
		return nil, nil, err
	}

	s := sig + frac
	tk := numberToken(r, start, s)

	return scanNext(r, tk)
}

// scanInt scans all digits that make up an integer. Underscores are allowed as
// separators.
func scanInt(r *Runer) (string, *token.Token) {
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return ``, runerErrorToken(r, err)
		}

		if !isDigit(ru) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return sb.String(), nil
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
	n, errTk := scanInt(r)

	if errTk != nil {
		return ``, errTk
	}

	if len(n) == 0 {
		return ``, badFractionalToken(r)
	}

	return string(ru) + n, nil
}

// badFractionalToken creates a new error token for invalid fractional syntax.
func badFractionalToken(r *Runer) *token.Token {
	return &token.Token{
		Line: r.Line(),
		End:  r.NextCol(),
		Type: token.TT_ERROR_UPSTREAM,
		Errors: []string{
			"Invalid number format, either...",
			"- fractional digits are missing",
			"- or the decimal separator is a typo",
		},
	}
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
