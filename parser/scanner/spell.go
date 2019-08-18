package scanner

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSpell scans symbols that start with a the '@' rune returning a spell
// identifier. Spells are inbuilt functions.
func scanSpell(r *Runer) token.Token {
	start := r.Col() + 1

	first, err := r.ReadRune()
	ru, _, err := r.LookAhead()
	if err != nil {
		return faultToToken(r, start, err)
	}

	if !isLetter(ru) {
		r.SkipRune()
		return errorToken(r, start, []string{
			"Expected first rune after '@' to be a letter",
		})
	}

	s, err := scanWordStr(r)
	if err != nil {
		return faultToToken(r, start, err)
	}

	s = string(first) + s
	return spellToken(r, start, s)
}

// spellToken creates a new spell token.
func spellToken(r *Runer, start int, val string) token.Token {
	return token.Token{
		Val:   val,
		Start: start,
		End:   r.Col() + 1,
		Type:  token.TT_SPELL,
	}
}
