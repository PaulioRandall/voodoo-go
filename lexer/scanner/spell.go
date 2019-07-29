package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSpell scans symbols that start with a the '@' rune
// returning a spell identifier. Spells are inbuilt functions.
func scanSpell(itr *runer.RuneItr) (tk *token.Token, err fault.Fault) {

	if !itr.IsRelLetter(1) {
		m := "Expected first rune after `@` to be a letter"
		err = fault.Func(m).SetFrom(itr.Index())
		return
	}

	start := itr.Index()
	h := string(itr.NextRune())
	t := scanWordStr(itr)

	tk = &token.Token{
		Val:   h + t,
		Start: start,
		End:   itr.Index(),
		Type:  token.SOURCERY,
	}

	return
}
