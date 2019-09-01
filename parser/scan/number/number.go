package number

import (
	"unicode"

	"github.com/PaulioRandall/voodoo-go/parser/scan/err"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/scantok"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// ScanNumber scans symbols that start with a unicode category Nd rune returning
// a literal number token.
func ScanNumber(r *runer.Runer) (token.Token, err.ScanError) {
	start := r.NextCol()

	whole, e := r.ReadWhile(isDigit)
	if e != nil {
		return nil, err.NewByRuner(r, e)
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
func isDigit(ru1, _ rune) (bool, error) {
	return unicode.IsDigit(ru1), nil
}

// scanFrac scans the delimiter and fractional part of a number.
func scanFrac(r *runer.Runer) (string, err.ScanError) {
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
func scanFracDelim(r *runer.Runer) (string, bool, err.ScanError) {

	isDelim := func(ru, _ rune) (bool, error) {
		return ru == '.', nil
	}

	if ru, is, e := r.ReadIf(isDelim); e != nil {
		return ``, false, err.NewByRuner(r, e)
	} else if is {
		return string(ru), true, nil
	}

	return ``, false, nil
}

// scanFracInt scans the fractional part of a number.
func scanFracInt(r *runer.Runer) (string, err.ScanError) {
	switch frac, e := r.ReadWhile(isDigit); {
	case e != nil:
		return ``, err.NewByRuner(r, e)
	case frac == ``:
		return ``, badFractionalToken(r)
	default:
		return frac, nil
	}
}

// badFractionalToken creates a new scan error for invalid fractional syntax.
func badFractionalToken(r *runer.Runer) err.ScanError {
	return err.New(
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
	return scantok.New(
		text,
		r.Line(),
		start,
		r.NextCol(),
		token.TT_NUMBER,
	)
}
