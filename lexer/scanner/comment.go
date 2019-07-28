package scanner

import (
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// scanComment scans symbols that start with a two consecutive
// forward slashes `//` returning a comment.
func scanComment(itr *runer.RuneItr) *symbol.Lexeme {

	start := itr.Index()
	str := itr.RemainingStr()

	return &symbol.Lexeme{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  symbol.COMMENT,
	}
}
