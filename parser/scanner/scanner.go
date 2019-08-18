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

	if shebang {
		tk = scanShebang(r)
		out <- tk
		if tk.Type == token.TT_ERROR_UPSTREAM {
			return nil
		}
	}

	for {
		ru1, ru2, err := r.LookAhead()
		if err != nil {
			return fault.ReaderFault(err.Error())
		}

		if ru1 == EOF {
			return nil
		}

		switch {
		case isNewline(ru1):
			tk = newlineToken(r)
		case isLetter(ru1):
			tk = scanWord(r)
		case isNaturalDigit(ru1):
			tk = scanNumber(r)
		case isSpace(ru1):
			tk = scanSpace(r)
		case isSpellPrefix(ru1):
			tk = scanSpell(r)
		case isStringPrefix(ru1):
			tk = scanString(r)
		case isCommentPrefix(ru1, ru2):
			tk = scanComment(r)
		default:
			tk = scanSymbol(r)
		}

		if tk.Type == token.TT_ERROR_UPSTREAM {
			out <- tk
			return nil
		}

		if err != nil {
			return fault.ReaderFault(err.Error())
		}

		out <- tk
	}
}

// scanShebang scans the shebang line.
func scanShebang(r *Runer) token.Token {
	start := r.Col() + 1
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return errorToToken(r, start, err)
		}

		if isNewline(ru) || ru == EOF {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return token.Token{
		Val:   sb.String(),
		Start: start,
		End:   r.Col() + 1,
		Type:  token.TT_SHEBANG,
	}
}

// newlineToken creates a newline token.
func newlineToken(r *Runer) token.Token {
	r.SkipRune()
	return token.Token{
		Val:   "\n",
		Start: r.Col(),
		End:   r.Col() + 1,
		Type:  token.TT_NEWLINE,
	}
}
