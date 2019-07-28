package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// scanNumber scans symbols that start with a unicode category Nd
// rune returning a literal number; all numbers are floats.
func scanNumber(itr *runer.RuneItr) (tk *symbol.Token, err fault.Fault) {

	start := itr.Index()
	s, err := extractNum(itr)
	if err != nil {
		return
	}

	tk = &symbol.Token{
		Val:   s,
		Start: start,
		End:   itr.Index(),
		Type:  symbol.LITERAL_NUMBER,
	}

	return
}

// extractNum extracts a number, as a string, from the supplied
// iterator.
func extractNum(itr *runer.RuneItr) (string, fault.Fault) {
	sb := strings.Builder{}
	var f string
	var err fault.Fault

	for itr.HasNext() {
		if itr.IsNextStr(`..`) {
			break
		}

		if itr.IsNext('.') {
			sb.WriteRune(itr.NextRune())
			f, err = extractFrac(itr)
			sb.WriteString(f)
			break
		}

		if !itr.IsNextDigit() && !itr.IsNext('_') {
			break
		}

		sb.WriteRune(itr.NextRune())
	}

	return sb.String(), err
}

// extractFrac extracts the fractional part of a number,
// as a string, from the supplied iterator and returns it.
func extractFrac(itr *runer.RuneItr) (string, fault.Fault) {
	sb := strings.Builder{}

	for itr.HasNext() {
		if itr.IsNext('.') {
			m := "Numbers can't have two fractional parts"
			return "", fault.Num(m).From(itr.Index())
		}

		if !itr.IsNextDigit() && !itr.IsNext('_') {
			break
		}

		sb.WriteRune(itr.NextRune())
	}

	return sb.String(), nil
}
