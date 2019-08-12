package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSpell scans symbols that start with a the '@' rune returning a spell
// identifier. Spells are inbuilt functions.
func scanSpell(r *Runer) (token.Token, fault.Fault) {

	first, err := r.ReadRune()
	ru, _, err := r.LookAhead()
	if err != nil {
		return token.EMPTY, err
	}

	if !isLetter(ru) {
		return token.EMPTY, badSpellName(r.Line(), r.Col()+2)
	}

	s, size, err := scanWordStr(r)
	if err != nil {
		return token.EMPTY, err
	}

	tk := token.Token{
		Val:   string(first) + s,
		Start: r.Col() - size,
		End:   r.Col() + 1,
		Type:  token.SPELL,
	}

	return tk, nil
}

// badSpellName creates a syntax fault for badly defined spell names.
func badSpellName(line, col int) fault.Fault {
	return token.SyntaxFault{
		Line: line,
		Col:  col,
		Msgs: []string{
			"Expected first rune after '@' to be a letter",
		},
	}
}
