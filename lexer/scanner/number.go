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

	size := len(frac)

	if size > 0 {
		if size < 2 {
			out = nil
			err = fault.SyntaxFault{
				Index: col + len(num) + 1,
				Msgs: []string{
					"Invalid number format, either...",
					" - fractional digits are missing",
					" - or the decimal separator is a typo",
				},
			}

			return
		}

		num = append(num, frac...)
	}

	tk = &token.Token{
		Val:   string(num),
		Start: col,
		Type:  token.LITERAL_NUMBER,
	}

	return
}
