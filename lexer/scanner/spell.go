package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// scanSpell scans symbols that start with a the '@' rune
// returning a spell identifier. Spells are inbuilt functions.
func scanSpell(itr *runer.RuneItr) (s *symbol.Lexeme, err fault.Fault) {

	if !itr.IsRelLetter(1) {
		m := "Expected first rune after `@` to be a letter"
		err = fault.Func(m).From(itr.Index())
		return
	}

	start := itr.Index()
	h := string(itr.NextRune())
	t := scanWordStr(itr)

	s = &symbol.Lexeme{
		Val:   h + t,
		Start: start,
		End:   itr.Index(),
		Type:  symbol.SOURCERY,
	}

	return
}
