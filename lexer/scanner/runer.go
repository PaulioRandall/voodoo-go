package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/runer"
)

// scanWordStr iterates a rune iterator until a single word has
// been extracted retruning the string.
func scanWordStr(itr *runer.RuneItr) string {
	sb := strings.Builder{}

	for itr.HasNext() {
		switch {
		case itr.IsNextLetter():
			fallthrough
		case itr.IsNextDigit():
			fallthrough
		case itr.IsNext('_'):
			sb.WriteRune(itr.NextRune())
		default:
			return sb.String()
		}
	}

	return sb.String()
}
