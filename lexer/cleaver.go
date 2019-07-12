package lexer

import (
	"strings"
)

// The purpose of the cleaver is to split up a string, usually a line,
// into symbols that represent either a potential lexeme/token or
// part of a one. Assuming valid syntax, each symbol can be one of:
// - terminal UTF-8 rune such as '=', '+', a whitespace rune, etc
// - part of a terminal UTF-8 string such as '=', '>', etc (resolving to '=>')
// - terminal UTF-8 string such as 'spell', 'true', etc
// - non-terminal UTF-8 string such as 'i', 'isAlive', etc
// - part of a non-terminal UTf-8 string such as 'is', '_', 'alive' that
//   may share runes with terminal symbols such that the prior
//   example comes together to form the identifier 'is_alive'.
//
// During the welding phase, fragmented symbols are joined together
// then scanned to form lexemes which are evaluated to form tokens.
// Tokens represent a meaningful symbol, and if applicable, with a value.
//
// A nice property of cleaving is that a symbol array containing a single
// empty symbol, one that contains no runes, always represents an
// empty line in the scroll. This will be useful for formatting or
// reccreating the scroll later if we wish; allthough all lines would be
// void of leading and trailing whiitespace, however, correct indentation
// can be inferred.
//
// TODO: Create SQL flow diagram of what the cleaver does!

// Cleave splits a string into symbols to make the rest of the
// scanning process an additive one rather than a splitting one.
func Cleave(s string, line int) []Symbol {

	r := []Symbol{}
	var f Symbol
	c := unicodeCat(none)
	sb := strings.Builder{}

	for i, ru := range s {
		cat := unicodeCatOf(ru)

		switch {
		case (c == none):
			f, c = reset(&sb, i, cat)
		case (c == letter) && (cat == letter):
		case (c == digit) && (cat == digit):
		default:
			r = push(r, f, &sb, i, line)
			f, c = reset(&sb, i, cat)
		}

		sb.WriteRune(ru)
	}

	r = push(r, f, &sb, len(s), line)
	return r
}

// push pushes a fragment onto the result list.
func push(r []Symbol, f Symbol, sb *strings.Builder, end int, line int) []Symbol {
	f.Val = sb.String()
	f.End = end
	f.Line = line
	r = append(r, f)
	return r
}

// reset resets the string builder and creates a new
// fragment.
func reset(sb *strings.Builder, start int, cat unicodeCat) (Symbol, unicodeCat) {
	sb.Reset()
	f := Symbol{
		Start: start,
	}
	return f, cat
}
