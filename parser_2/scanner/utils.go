package scanner

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// ParseToken represents a function produces a single specific type of Token
// from an input stream. Only the first part of the stream that provides longest
// match against the production rules of the specific Token will be read. If an
// error occurs then an error token will be returned instead.
//
// The first returned token alwyas represents a valid parsed token while the
// last always represents an error. On return, one should be nil and the other
// non-nil.
type ParseToken func(*runer.Runer) (token.Token, ParseToken, ScanError)

// runerError creates a new ScanError from an error returned by a Runer.
func runerError(r *runer.Runer, err error) ScanError {
	return &scanErr{
		l: r.Line(),
		i: r.NextCol(),
		e: []string{err.Error()},
	}
}
