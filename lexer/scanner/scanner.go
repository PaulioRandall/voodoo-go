package scanner

import (
	"bufio"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Scan scans tokens from a stream of code using longest match and pushes them
// onto a channel for processing.
func Scan(br *bufio.Reader, out chan token.Token) fault.Fault {
	defer close(out)

	r := NewRuner(br)

	for {
		var tk token.Token
		var err fault.Fault

		ru1, ru2, err := r.LookAhead()
		if err != nil {
			return err
		}

		if ru1 == EOF {
			return nil
		}

		switch {
		case isLetter(ru1):
			tk, err = scanWord(r)
		case isNaturalDigit(ru1):
			tk, err = scanNumber(r)
		case isSpace(ru1):
			tk, err = scanSpace(r)
		case isSpellPrefix(ru1):
			tk, err = scanSpell(r)
		case isStringPrefix(ru1):
			tk, err = scanString(r)
		case isCommentPrefix(ru1, ru2):
			tk, err = scanComment(r)
		default:
			tk, err = scanSymbol(r)
		}

		if err != nil {
			return err
		}

		out <- tk
	}
}
