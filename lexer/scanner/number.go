package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/scroll"
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
			err = fault.Num(m)
			err = fault.SetTo(err, len(num)+1)
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

// numLiteralFault represents a fault with a literal number.
type numLiteralFault struct {
	Index int      // Index where the error actually occurred
	Msgs  []string // Description of the error
}

// Print satisfies the Fault interface.
func (err numLiteralFault) Print(sc *scroll.Scroll, line int) {
	sc.PrettyPrintError(line, err.Index, err.Msgs...)
}
