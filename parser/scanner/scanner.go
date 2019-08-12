package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Scan scans tokens from a stream of code using longest match and pushes them
// onto a channel for processing.
func Scan(r *Runer, shebang bool, out chan token.Token) fault.Fault {
	defer close(out)

	var tk token.Token
	var err fault.Fault

	if shebang {
		tk, err = scanShebang(r)
		if err != nil {
			return err
		}
		out <- tk
	}

	for {
		ru1, ru2, err := r.LookAhead()
		if err != nil {
			return err
		}

		if ru1 == EOF {
			return nil
		}

		switch {
		case isNewline(ru1):
			tk = newlineToken(r)
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

// scanShebang scans the shebang line.
func scanShebang(r *Runer) (token.Token, fault.Fault) {
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return token.EMPTY, err
		}

		if isNewline(ru) || ru == EOF {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	s := sb.String()
	size := len(s)

	tk := token.Token{
		Val:   sb.String(),
		Start: r.Col() - size + 1,
		End:   r.Col() + 1,
		Type:  token.SHEBANG,
	}

	return tk, nil
}

// newlineToken creates a newline token.
func newlineToken(r *Runer) token.Token {
	r.SkipRune()
	return token.Token{
		Val:   "\n",
		Start: r.Col(),
		End:   r.Col() + 1,
		Type:  token.NEWLINE,
	}
}
