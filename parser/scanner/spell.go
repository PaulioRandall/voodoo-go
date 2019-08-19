package scanner

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSpell scans symbols that start with a the '@' rune returning a spell
// identifier. Spells are inbuilt functions.
func scanSpell(r *Runer) token.Token {
	start := r.NextCol()

	first, err := r.ReadRune()
	if err != nil {
		return *runerErrorToken(r, err)
	}

	ru, _, err := r.LookAhead()
	if err != nil {
		return *runerErrorToken(r, err)
	}

	if !isLetter(ru) {
		r.SkipRune()
		return *errorToken(r, []string{"Expected letter after '@'"})
	}

	s, err := scanWordStr(r)
	if err != nil {
		return *runerErrorToken(r, err)
	}

	s = string(first) + s
	return spellToken(r, start, s)
}

// spellToken creates a new spell token.
func spellToken(r *Runer, start int, val string) token.Token {
	return token.Token{
		Val:   val,
		Line:  r.Line(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_SPELL,
	}
}
