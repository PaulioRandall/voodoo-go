package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanNumber scans symbols that start with a unicode category Nd rune returning
// a literal number token.
func scanNumber(r *Runer) token.Token {
	start := r.Col() + 1

	sig, err := scanSignificant(r)
	if err != nil {
		return errorToken(r, start, err)
	}

	frac, err := scanFractional(r)
	if err != nil {
		return errorToken(r, start, err)
	}

	s := sig + frac
	return numberToken(r, start, s)
}

// numberToken creates a new number token.
func numberToken(r *Runer, start int, val string) token.Token {
	return token.Token{
		Val:   val,
		Line:  r.Line(),
		Start: start,
		End:   r.Col() + 1,
		Type:  token.TT_NUMBER,
	}
}

// scanSignificant scans the significant part of a number.
func scanSignificant(r *Runer) (string, []string) {
	first, err := r.ReadRune()
	if err != nil {
		return ``, []string{err.Error()}
	}

	sig, errs := scanInt(r)
	if errs != nil {
		return ``, errs
	}

	return string(first) + sig, nil
}

// scanFractional scans the fractional part of a number including the decimal
// separator.
func scanFractional(r *Runer) (string, []string) {
	ru, _, err := r.LookAhead()
	if err != nil {
		return ``, []string{err.Error()}
	}

	if !isDecimalDelim(ru) {
		return ``, nil
	}

	r.SkipRune()
	frac, errs := scanInt(r)
	if errs != nil {
		return ``, errs
	}

	if len(frac) == 0 || !strings.ContainsAny(frac, "0123456789") {
		return ``, []string{
			"Invalid number format, either...",
			" - fractional digits are missing",
			" - or the decimal separator is a typo",
		}
	}

	return string(ru) + frac, nil
}

// scanInt scans an integer.
func scanInt(r *Runer) (string, []string) {
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return ``, []string{err.Error()}
		}

		if !isDigit(ru) && !isUnderscore(ru) {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return sb.String(), nil
}
