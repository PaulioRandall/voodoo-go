package scanner_new

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// ParseToken represents a function produces a single specific type of Token
// from an input stream. Only the first part of the stream that provides longest
// match against the production rules of the specific Token will be read. If an
// error occurs then an error token will be returned instead.
type ParseToken func(*Runer) token.Token

// Scan finds an appropriate function to parse the next token producable from
// the Runer.
func NEW_Scan(r *Runer) (f ParseToken, errTk *token.Token) {

	ru1, ru2, err := r.LookAhead()
	if err != nil {
		errTk = runerErrorToken(r, err)
		return
	}

	switch {
	case ru1 == EOF:
	case r.Line() == 0:
		f = NEW_scanShebang
	case isNewline(ru1):
		f = scanNewline
	case isLetter(ru1):
		//f = scanWord
	case isNaturalDigit(ru1):
		//f = scanNumber
	case isSpace(ru1):
		//f = scanSpace
	case isSpellPrefix(ru1):
		//f = scanSpell
	case isStringPrefix(ru1):
		//f = scanString
	case isCommentPrefix(ru1, ru2):
		//f = scanComment
	default:
		//f = scanSymbol
	}

	return
}

// scanShebang scans a shebang line.
func NEW_scanShebang(r *Runer) token.Token {
	start := r.NextCol()
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return *runerErrorToken(r, err)
		}

		if isNewline(ru) || ru == EOF {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	return token.Token{
		Val:   sb.String(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_SHEBANG,
	}
}

// scanNewline scans a newline token.
func scanNewline(r *Runer) token.Token {
	r.SkipRune()
	return token.Token{
		Val:   "\n",
		Start: r.Col(),
		End:   r.NextCol(),
		Type:  token.TT_NEWLINE,
	}
}
