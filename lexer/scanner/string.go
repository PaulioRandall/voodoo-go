package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// scanString scans symbols that start and end with an
// non-escaped `"` returning a string literal token.
func scanString(itr *runer.RuneItr) (l *symbol.Lexeme, err fault.Fault) {

	start := itr.Index()
	closed, s := extractStr(itr)

	if !closed {
		m := "Did someone forget to close a string literal?!"
		err = fault.Str(m).From(start).To(itr.Index())
		return
	}

	l = &symbol.Lexeme{
		Val:   s,
		Start: start,
		End:   itr.Index(),
		Type:  symbol.LITERAL_STRING,
	}

	return
}

// extractStr extracts a string literal from a string iterator
// returning true if the last rune was escaped.
func extractStr(itr *runer.RuneItr) (closed bool, s string) {

	sb := strings.Builder{}
	sb.WriteRune(itr.NextRune())
	isEscaped := false

	for itr.HasNext() {
		ru := itr.NextRune()
		sb.WriteRune(ru)

		if !isEscaped && ru == '"' {
			closed = true
			break
		}

		if ru == '\\' {
			isEscaped = !isEscaped
		} else {
			isEscaped = false
		}
	}

	s = sb.String()
	return
}
