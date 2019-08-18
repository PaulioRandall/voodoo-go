package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanString scans symbols that start and end with an non-escaped `"` returning
// a string literal token.
func scanString(r *Runer) token.Token {
	start := r.Col() + 1

	s, err := scanStr(r)
	if err != nil {
		return errorToken(r, start, err)
	}

	tk := token.Token{
		Val:   s,
		Line:  r.Line(),
		Start: start,
		End:   r.Col() + 1,
		Type:  token.TT_STRING,
	}

	return tk
}

// scanStr extracts a string literal from a string iterator returning true if
// the last rune was escaped.
func scanStr(r *Runer) (string, []string) {

	open, err := r.ReadRune()
	if err != nil {
		return ``, readerFaultToStringArray(err)
	}

	body, errs := scanStrBody(r)
	if errs != nil {
		return ``, errs
	}

	close, err := r.ReadRune()
	if err != nil {
		return ``, readerFaultToStringArray(err)
	}

	s := string(open) + body + string(close)
	return s, nil
}

// scanStrBody scans the body of a string literal.
func scanStrBody(r *Runer) (string, []string) {
	sb := strings.Builder{}
	isEscaped := false

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return ``, readerFaultToStringArray(err)
		}

		if !isEscaped && ru == '"' {
			break
		}

		if ru == EOF || isNewline(ru) {
			return ``, []string{
				"Did someone forget to close a string literal?!",
			}
		}

		if ru == '\\' {
			isEscaped = !isEscaped
		} else {
			isEscaped = false
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return sb.String(), nil
}
