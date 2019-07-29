package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Scan scans a line and creates an array of tokens based on
// the grammer rules of the language. Longest match is used to
// identify variable names and keywords etc.
//
// No panic is generated by the scanner so if a panic occurs it's
// either a system issue or a bug.
func Scan(s string) (out []token.Token, err fault.Fault) {

	out = []token.Token{}

	if s == `` {
		return
	}

	in := []rune(s)
	size := len(in)
	i := 0

	for i < size {
		var tk *token.Token
		r := in[0]

		switch {
		case isLetter(r):
			tk, in = scanWord(in)
		case isDigit(r):
			tk, in, err = scanNumber(in)
		case isSpace(r):
			tk, in = scanSpace(in)
		case isSpellStart(r):
			tk, in, err = scanSpell(in)
		case isStrStart(r):
			tk, in, err = scanString(in)
		case startsWith(in, `//`):
			tk = &token.Token{
				Val:  string(in),
				Type: token.COMMENT,
			}
			in = []rune{}
		default:
			tk, in, err = scanSymbol(in)
		}

		if err != nil {
			out = nil
			err = err.SetFrom(i)
			err = err.SetTo(i + err.To())
			return
		}

		tk.Start = i
		i = size - len(in)
		tk.End = i
		out = append(out, *tk)
	}

	return
}
