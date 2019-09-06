package number

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// ScanNumber scans symbols that start with a unicode category Nd rune returning
// a literal number token.
func ScanNumber(r *runer.Runer) (token.Token, perror.Perror) {
	start := r.NextCol()

	whole, e := r.ReadWhile(isDigit)
	if e != nil {
		return nil, e
	}

	frac, sce := scanFrac(r)
	if sce != nil {
		return nil, sce
	}

	s := whole + frac
	return numberToken(r, start, s), nil
}

// isDigit returns true if the input rune is a digit. This function implements
// the runer.predicate interface.
func isDigit(ru1, _ rune) bool {
	return unicode.IsDigit(ru1)
}

// scanFrac scans the delimiter and fractional part of a number.
func scanFrac(r *runer.Runer) (string, perror.Perror) {
	delim, is, sce := scanFracDelim(r)
	if sce != nil {
		return ``, sce
	}

	if !is {
		return ``, nil
	}

	frac, sce := scanFracInt(r)
	if sce != nil {
		return ``, sce
	}
	return string(delim) + frac, nil
}

// scanFracDelim scans the delimiter between the whole part and fractional part
// of a floating point number. False is returned if no delimiter is present.
func scanFracDelim(r *runer.Runer) (string, bool, perror.Perror) {

	isDelim := func(ru, _ rune) bool {
		return ru == '.'
	}

	if ru, is, e := r.ReadIf(isDelim); e != nil {
		return ``, false, e
	} else if is {
		return string(ru), true, nil
	}

	return ``, false, nil
}

// scanFracInt scans the fractional part of a number.
func scanFracInt(r *runer.Runer) (string, perror.Perror) {
	switch frac, e := r.ReadWhile(isDigit); {
	case e != nil:
		return ``, e
	case frac == ``:
		return ``, badFractional(r)
	default:
		return frac, nil
	}
}

// badFractional creates a new scan error for invalid fractional syntax.
func badFractional(r *runer.Runer) perror.Perror {
	return perror.New(
		r.Line(),
		r.NextCol(),
		[]string{
			"Invalid number format, either...",
			"- fractional digits are missing",
			"- or the decimal separator is a typo",
		},
	)
}

// numberToken creates a new number token.
func numberToken(r *runer.Runer, start int, text string) token.Token {
	return token.New(
		text,
		r.Line(),
		start,
		r.NextCol(),
		token.TT_NUMBER,
	)
}
