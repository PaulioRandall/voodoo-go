package string

import (
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// ScanString scans string tokens returning a string literal.
func ScanString(r *runer.Runer) (token.Token, perror.Perror) {
	start := r.NextCol()

	first, closed := true, false
	readString := func(ru, _ rune) bool {
		switch {
		case first:
			first = false
			return true
		case closed:
			return false
		case ru == '`':
			closed = true
			return true
		default:
			return ru != '\n'
		}
	}

	s, e := r.ReadWhile(readString)
	if e != nil {
		return nil, e
	}

	if !closed {
		return nil, unclosedString(r)
	}

	return stringToken(r, start, s), nil
}

// unclosedString creates a new Perror for an unclosed string.
func unclosedString(r *runer.Runer) perror.Perror {
	return perror.New(
		r.Line(),
		r.NextCol(),
		[]string{
			"Did some one forget to close a string literal?!",
		},
	)
}

// stringToken returns a new string Token.
func stringToken(r *runer.Runer, start int, s string) token.Token {
	return token.New(
		s,
		r.Line(),
		start,
		r.NextCol(),
		token.TK_STRING,
	)
}
