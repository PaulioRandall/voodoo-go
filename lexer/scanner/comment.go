package scanner

import (
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanComment scans symbols that start with a two consecutive
// forward slashes `//` returning a comment.
func scanComment(itr *runer.RuneItr) *token.Token {

	start := itr.Index()
	str := itr.RemainingStr()

	return &token.Token{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  token.COMMENT,
	}
}
