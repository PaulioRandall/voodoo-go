package scanner_new

import (
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// scanSpell scans symbols that start with a the '@' rune returning a spell
// identifier. Spells are inbuilt functions.
func scanSpell(r *Runer) (tk *token.Token, _ ParseToken, errTk *token.Token) {
	start := r.NextCol()

	first, err := r.ReadRune()
	if err != nil {
		errTk = runerErrorToken(r, err)
		return
	}

	ru, _, err := r.LookAhead()
	if err != nil {
		errTk = runerErrorToken(r, err)
		return
	}

	if !isLetter(ru) {
		r.SkipRune()
		errTk = errorToken(r, []string{"Expected letter after '@'"})
		return
	}

	s, err := scanWordStr(r)
	if err != nil {
		errTk = runerErrorToken(r, err)
		return
	}

	s = string(first) + s
	tk = spellToken(r, start, s)

	return ScanNext(r, tk)
}

// spellToken creates a new spell token.
func spellToken(r *Runer, start int, val string) *token.Token {
	return &token.Token{
		Val:   val,
		Line:  r.Line(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_SPELL,
	}
}
