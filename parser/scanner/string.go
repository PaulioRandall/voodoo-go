package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanString scans symbols that start and end with an non-escaped `"` returning
// a string literal token.
func scanString(r *Runer) (token.Token, fault.Fault) {
	s, size, err := scanStr(r)
	if err != nil {
		return token.ERROR, err
	}

	tk := token.Token{
		Val:   s,
		Start: r.Col() - size + 1,
		End:   r.Col() + 1,
		Type:  token.TT_STRING,
	}

	return tk, nil
}

// scanStr extracts a string literal from a string iterator returning true if
// the last rune was escaped.
func scanStr(r *Runer) (string, int, fault.Fault) {

	open, err := r.ReadRune()
	if err != nil {
		return ``, -1, err
	}

	body, err := scanStrBody(r)
	if err != nil {
		return ``, -1, err
	}

	close, err := r.ReadRune()
	if err != nil {
		return ``, -1, err
	}

	s := string(open) + body + string(close)
	return s, len(s), nil
}

// scanStrBody scans the body of a string literal.
func scanStrBody(r *Runer) (string, fault.Fault) {
	sb := strings.Builder{}
	isEscaped := false

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return ``, err
		}

		if !isEscaped && ru == '"' {
			break
		}

		if ru == EOF || isNewline(ru) {
			return ``, unclosedString(r.Line(), r.Col()+1)
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

// unclosedString creates a fault for when a string literal is has not been
// closed before the end of a line or file.
func unclosedString(line, col int) fault.Fault {
	return token.SyntaxFault{
		Line: line,
		Col:  col,
		Msgs: []string{
			"Did someone forget to close a string literal?!",
		},
	}
}
