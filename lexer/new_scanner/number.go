package new_scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanNumber scans symbols that start with a unicode category Nd rune returning
// a literal number token.
func scanNumber(r *Runer) (token.Token, fault.Fault) {

	sig, err := scanSignificant(r)
	if err != nil {
		return token.EMPTY, err
	}

	frac, err := scanFractional(r)
	if err != nil {
		return token.EMPTY, err
	}

	s := sig + frac
	size := len(s)

	tk := token.Token{
		Val:   s,
		Start: r.Col() - size + 1,
		End:   r.Col() + 1,
		Type:  token.LITERAL_NUMBER,
	}

	return tk, nil
}

// scanSignificant scans the significant part of a number.
func scanSignificant(r *Runer) (string, fault.Fault) {
	first, err := r.ReadRune()
	if err != nil {
		return ``, err
	}

	sig, err := scanInt(r)
	if err != nil {
		return ``, err
	}

	return string(first) + sig, nil
}

// scanFractional scans the fractional part of a number including the decimal
// separator.
func scanFractional(r *Runer) (string, fault.Fault) {
	ru, _, err := r.LookAhead()
	if err != nil {
		return ``, err
	}

	if !isDecimalDelim(ru) {
		return ``, nil
	}

	r.SkipRune()
	frac, err := scanInt(r)
	if err != nil {
		return ``, err
	}

	if len(frac) == 0 {
		return ``, badNumberFormat(r.Col() + 1)
	}

	return string(ru) + frac, nil
}

// scanInt scans an integer.
func scanInt(r *Runer) (string, fault.Fault) {
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

// badNumberFormat returns a new fault detailing when a number is badly
// formatted.
func badNumberFormat(i int) fault.Fault {
	return fault.SyntaxFault{
		Index: i,
		Msgs: []string{
			"Invalid number format, either...",
			" - fractional digits are missing",
			" - or the decimal separator is a typo",
		},
	}
}
