package scanner

import (
	fault "github.com/PaulioRandall/voodoo-go/new_fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSpell scans symbols that start with a the '@' rune
// returning a spell identifier. Spells are inbuilt functions.
//
// This function asumes the first rune of the input array is a '@'.
func scanSpell(in []rune) (tk *token.Token, out []rune, err fault.Fault) {

	if len(in) < 2 || !isLetter(in[1]) {
		err = fault.SyntaxFault{
			Index: 1,
			Msgs: []string{
				"Expected first rune after '@' to be a letter",
			},
		}
		return
	}

	var s string
	at := string(in[:1])
	s, out = scanWordStr(in[1:])

	tk = &token.Token{
		Val:  at + s,
		Type: token.SPELL,
	}

	return
}
