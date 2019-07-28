package scanner

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// scanSpace scans symbols that start with a unicode whitespace
// property rune returning a token representing all whitespace
// between two non-whitespace tokens.
//
// Note that there is an intention to switch to stream based
// scanning. When this change happens newline runes will
// become the one exception to the rule as they will become
// a token all by themselves used to delimit statements
// and the bodies of different context.
func scanSpace(itr *runer.RuneItr) *symbol.Token {

	start := itr.Index()
	sb := strings.Builder{}

	for itr.HasNext() {
		if !itr.IsNextSpace() {
			break
		}
		sb.WriteRune(itr.NextRune())
	}

	return &symbol.Token{
		Val:   sb.String(),
		Start: start,
		End:   itr.Index(),
		Type:  symbol.WHITESPACE,
	}
}
