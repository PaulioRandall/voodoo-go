package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanString scans symbols that start and end with an non-escaped `"`
// returning a string literal token.
//
// This function asumes the first rune of the input array is a '"'.
func scanString(in []rune) (tk *token.Token, out []rune, err fault.Fault) {

	var s string
	var closed bool
	s, out, closed = scanStr(in)

	if !closed {
		err = fault.SyntaxFault{
			Index: len(in),
			Msgs: []string{
				"Did someone forget to close a string literal?!",
			},
		}
		return
	}

	tk = &token.Token{
		Val:  s,
		Type: token.LITERAL_STRING,
	}

	return
}

// scanStr extracts a string literal from a string iterator
// returning true if the last rune was escaped.
func scanStr(in []rune) (s string, out []rune, closed bool) {

	isEscaped := false
	end := -1

	for i, r := range in[1:] {

		if !isEscaped && r == '"' {
			end = i + 1 // +1 because first rune wwaas ignored
			end += 1    // +1 converts last index to length
			break
		}

		if r == '\\' {
			isEscaped = !isEscaped
		} else {
			isEscaped = false
		}
	}

	if end == -1 {
		return
	}

	closed = true
	s = string(in[:end])
	out = in[end:]
	return
}
