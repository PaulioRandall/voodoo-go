package scanner_new

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// TODO: Maybe Token should be an interface?

// ParseToken represents a function produces a single specific type of Token
// from an input stream. Only the first part of the stream that provides longest
// match against the production rules of the specific Token will be read. If an
// error occurs then an error token will be returned instead.
//
// The first returned token alwyas represents a valid parsed token while the
// last always represents an error. On return, one should be nil and the other
// non-nil.
type ParseToken func(*Runer) (tk *token.Token, f ParseToken, errTk *token.Token)

// Scan finds an appropriate function to parse the next token producable from
// the Runer.
func Scan(r *Runer) (f ParseToken, errTk *token.Token) {

	ru1, ru2, err := r.LookAhead()
	if err != nil {
		errTk = runerErrorToken(r, err)
		return
	}

	switch {
	case ru1 == EOF:
	case r.Line() == 0:
		f = scanShebang
	case isNewline(ru1):
		f = scanNewline
	case isLetter(ru1):
		f = scanWord
	case isNaturalDigit(ru1):
		f = scanNumber
	case isSpace(ru1):
		f = scanSpace
	case isSpellPrefix(ru1):
		f = scanSpell
	case isStringPrefix(ru1):
		f = scanString
	case isCommentPrefix(ru1, ru2):
		f = scanComment
	default:
		f = scanSymbol
	}

	return
}

// scanNext invokes Scan() returning the input token and the next ParseToken
// function to execute. If Scan() fails then an error Token is returned instead.
func scanNext(r *Runer, tk *token.Token) (*token.Token, ParseToken, *token.Token) {
	f, errTk := Scan(r)
	if errTk != nil {
		return nil, nil, errTk
	}
	return tk, f, nil
}

// scanShebang scans a shebang line.
func scanShebang(r *Runer) (*token.Token, ParseToken, *token.Token) {
	start := r.NextCol()
	sb := strings.Builder{}

	for {
		ru, _, err := r.LookAhead()
		if err != nil {
			return nil, nil, runerErrorToken(r, err)
		}

		if isNewline(ru) || ru == EOF {
			break
		}

		r.SkipRune()
		sb.WriteRune(ru)
	}

	tk := &token.Token{
		Val:   sb.String(),
		Start: start,
		End:   r.NextCol(),
		Type:  token.TT_SHEBANG,
	}

	return scanNext(r, tk)
}

// scanNewline scans a newline token.
func scanNewline(r *Runer) (*token.Token, ParseToken, *token.Token) {
	r.SkipRune()
	tk := &token.Token{
		Val:   "\n",
		Start: r.Col(),
		End:   r.NextCol(),
		Type:  token.TT_NEWLINE,
	}

	return scanNext(r, tk)
}
