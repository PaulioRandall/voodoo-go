package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanNumber scans symbols that start with a unicode category Nd
// rune returning a literal number; all numbers are floats.
func scanNumber(in []rune) (tk *token.Token, out []rune, err fault.Fault) {

	var num []rune
	var frac []rune

	num, in = scanInt(in)
	frac, out = scanFrac(in)

	size := len(frac)

	if size > 0 {
		if size < 2 {
			m := "Number has a decimal separator but the fractional digits are missing"
			err = fault.Num(m).SetTo(len(num) + 1)
			out = nil
			return
		}

		num = append(num, frac...)
	}

	tk = &token.Token{
		Val:  string(num),
		Type: token.LITERAL_NUMBER,
	}

	return
}
