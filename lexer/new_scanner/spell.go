package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSpell scans symbols that start with a the '@' rune
// returning a spell identifier. Spells are inbuilt functions.
func scanSpell(in []rune) (tk *token.Token, out []rune, err fault.Fault) {

	if len(in) < 2 || !isLetter(in[1]) {
		m := "Expected first rune after `@` to be a letter"
		err = fault.Func(m)
		return
	}

	var s string
	at := string(in[:1])
	s, out = scanWordStr(in[1:])

	tk = &token.Token{
		Val:  at + s,
		Type: token.SOURCERY,
	}

	return
}
