package lexer

import (
	"unicode"
)

// ScanLine scans a line and creates an array of symbols
// based on the grammer rules of the language.
func ScanLine(line string, lineNum int) []Symbol {

	if line == `` {
		// TODO: Move this to its own function
		return []Symbol{
			Symbol{
				Val:   ``,
				Start: 0,
				End:   0,
				Line:  lineNum,
			},
		}
	}

	itr := NewStrItr(line)
	r := []Symbol{}

	for itr.HasNext() {
		ru := itr.Peek()

		switch {
		case unicode.IsLetter(ru):
			r = append(r, wordSym(itr))
		case unicode.IsDigit(ru):
			r = append(r, numSym(itr))
		case unicode.IsSpace(ru):
			r = append(r, spaceSym(itr))
		case ru == '@':
			r = append(r, curseSym(itr))
		case ru == '"':
			r = append(r, strSym(itr))
		case isComment(itr):
			r = append(r, commentSym(itr))
		default:
			r = append(r, otherSym(itr))
		}
	}

	return r
}

// wordSym handles symbols that start with a unicode category L rune.
// I.e. a letter from any alphabet, a word may resolve into a:
// - variable name
// - keyword
// - boolean value (`true` or `false`)
func wordSym(itr *StrItr) Symbol {
	// TODO
	return Symbol{}
}

// numSym handles symbols that start with a unicode category Nd rune.
// I.e. any number from 0 to 9, a number may resolve into a:
// - literal number
func numSym(itr *StrItr) Symbol {
	// TODO
	return Symbol{}
}

// spaceSym handles symbols that start with a rune with the
// unicode whitespace property.
// I.e. any whitespace rune, whitespace may resolve into a:
// - meaningless symbol that can be ignored when parsing
func spaceSym(itr *StrItr) Symbol {
	// TODO
	return Symbol{}
}

// curseSym handles symbols that start with a at sign rune `@`.
// Curse symbols may resolve into a:
// - go function call
func curseSym(itr *StrItr) Symbol {
	// TODO
	return Symbol{}
}

// strSym handles symbols that start with the double quote `"` rune.
// Quoted strings may resolve into a:
// - string literal
func strSym(itr *StrItr) Symbol {
	// TODO
	return Symbol{}
}

// isComment return true if the rest of the string is a comment.
func isComment(itr *StrItr) bool {
	// TODO
	return false
}

// commentSym handles symbols that start with two forward slashes
// `//`. Double forward slashes may resolve into a:
// - comment
func commentSym(itr *StrItr) Symbol {
	// TODO
	return Symbol{}
}

// otherSym handles any symbols that don't have a specific handling
// function. These symbols may resolve into a:
// - operator, 1 or 2 runes including truthy and not
// - code block start or end, i.e. bracket
// - value separator, i.e. comma
// - key-value separator, i.e. colon
// - void value, i.e. underscore
func otherSym(itr *StrItr) Symbol {
	// TODO
	return Symbol{}
}
