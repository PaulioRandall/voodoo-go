package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanNumber scans symbols that start with a unicode category Nd
// rune returning a literal number; all numbers are floats.
func scanNumber(in []rune, col int) (tk *token.Token, out []rune, err fault.Fault) {

	var num []rune
	var frac []rune

	num, in = scanInt(in)
	frac, out = scanFrac(in)

	num, err = appendFractional(num, frac, col)
	if err != nil {
		out = nil
		return
	}

	tk = &token.Token{
		Val:   string(num),
		Start: col,
		Type:  token.LITERAL_NUMBER,
	}

	return
}

// appendFractional appends the fractional part of the number.
func appendFractional(num, frac []rune, col int) ([]rune, fault.Fault) {
	size := len(frac)

	if size > 0 {
		if size < 2 {
			return nil, badNumberFormat(col, len(num))
		}

		num = append(num, frac...)
	}

	return num, nil
}

// badNumberFormat returns a fault detailing a bad number
// format.
func badNumberFormat(col, offset int) fault.Fault {
	return fault.SyntaxFault{
		Index: col + offset + 1,
		Msgs: []string{
			"Invalid number format, either...",
			" - fractional digits are missing",
			" - or the decimal separator is a typo",
		},
	}
}
