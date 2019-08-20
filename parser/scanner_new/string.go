package scanner_new

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanString scans symbols that start and end with an non-escaped `"` returning
// a string literal token.
func scanString(r *Runer) (*token.Token, ParseToken, *token.Token) {
	start := r.NextCol()

	s, errTk := scanStr(r)
	if errTk != nil {
		return nil, nil, errTk
	}

	tk := &token.Token{
		Val:   s,
		Line:  r.Line(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_STRING,
	}

	return scanNext(r, tk)
}

// scanStr extracts a string literal, including the quotes, from a Runer.
func scanStr(r *Runer) (string, *token.Token) {

	open, err := r.ReadRune()
	if err != nil {
		return ``, runerErrorToken(r, err)
	}

	body, errTk := scanStrBody(r)
	if errTk != nil {
		return ``, errTk
	}

	close, err := r.ReadRune()
	if err != nil {
		return ``, runerErrorToken(r, err)
	}

	s := string(open) + body + string(close)
	return s, nil
}

// scanStrBody scans the body of a string literal, everything between the
// quotes.
func scanStrBody(r *Runer) (string, *token.Token) {
	sb := strings.Builder{}
	isEscaped := false

	for {
		ru, errTk := lookAhead(r)
		if errTk != nil {
			return ``, errTk
		}

		if !isEscaped && ru == '"' {
			break
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

// lookAhead returns the next rune in the Runer without incrementing its
// internal rune index counter and checks the EOF has not been reached.
func lookAhead(r *Runer) (rune, *token.Token) {
	ru, _, err := r.LookAhead()
	if err != nil {
		return ru, runerErrorToken(r, err)
	}

	if ru == EOF || isNewline(ru) {
		return ru, errorToken(r, []string{
			`Did someone forget to close a string literal?!`,
		})
	}

	return ru, nil
}
