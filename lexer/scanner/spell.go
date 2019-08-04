package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSpell scans symbols that start with a the '@' rune
// returning a spell identifier. Spells are inbuilt functions.
//
// This function asumes the first rune of the input array is a '@'.
func scanSpell(in []rune, col int) (tk *token.Token, out []rune, err fault.Fault) {

	if len(in) < 2 || !isLetter(in[1]) {
		err = badSpellName(col)
		return
	}

	var s string
	at := string(in[:1])
	s, out = scanWordStr(in[1:])

	tk = &token.Token{
		Val:   at + s,
		Start: col,
		Type:  token.SPELL,
	}

	return
}

// badSpellName creates a syntax fault for badly defined spell
// names.
func badSpellName(col int) fault.Fault {
	return fault.SyntaxFault{
		Index: col + 1,
		Msgs: []string{
			"Expected first rune after '@' to be a letter",
		},
	}
}
