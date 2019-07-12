package lexer

import (
	"errors"
	"strings"
	"unicode"

	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// ScanLine scans a line and creates an array of symbols
// based on the grammer rules of the language.
func ScanLine(line string, lineNum int) []Symbol {

	if line == `` {
		return emptyLineSyms(lineNum)
	}

	itr := NewStrItr(line)
	r := []Symbol{}

	for itr.HasNext() {
		ru := itr.Peek()
		start := itr.NextIndex()
		var s Symbol
		var err error

		switch {
		case unicode.IsLetter(ru):
			s, err = wordSym(itr, lineNum)
		case unicode.IsDigit(ru):
			s, err = numSym(itr, lineNum)
		case unicode.IsSpace(ru):
			s, err = spaceSym(itr, lineNum)
		case ru == '@':
			s, err = curseSym(itr, lineNum)
		case ru == '"':
			s, err = strSym(itr, lineNum)
		case isComment(itr):
			s, err = commentSym(itr, lineNum)
		default:
			s, err = otherSym(itr, lineNum)
		}

		if err != nil {
			end := itr.NextIndex()
			sh.SyntaxError(lineNum, start, end, err)
		}

		r = append(r, s)
		itr.Next() // TEMP
	}

	return r
}

// emptyLineSyms returns an array containing a single empty
// symbol that represents an empty line.
func emptyLineSyms(lineNum int) []Symbol {
	return []Symbol{
		Symbol{
			Val:   ``,
			Start: 0,
			End:   0,
			Line:  lineNum,
		},
	}
}

// initSym creates a new symbol with start index and line number
// initialised.
func initSym(start, lineNum int) Symbol {
	return Symbol{
		Start: start,
		Line:  lineNum,
	}
}

// wordSym handles symbols that start with a unicode category L rune.
// I.e. a letter from any alphabet, a word may resolve into a:
// - variable name
// - keyword
// - boolean value (`true` or `false`)
func wordSym(itr *StrItr, lineNum int) (Symbol, error) {

	ru := itr.Peek()
	if !itr.HasNext() || !unicode.IsLetter(ru) {
		m := "You can't call this function unless the iterators first rune is a letter"
		sh.CompilerBug(lineNum, m)
	}

	r := extractWord(itr, lineNum)
	return r, nil
}

// numSym handles symbols that start with a unicode category Nd rune.
// I.e. any number from 0 to 9, a number may resolve into a:
// - literal number
func numSym(itr *StrItr, lineNum int) (Symbol, error) {

	ru := itr.Peek()
	if !itr.HasNext() || !unicode.IsDigit(ru) {
		m := "You can't call this function unless the iterators first rune is a digit"
		sh.CompilerBug(lineNum, m)
	}

	r := initSym(itr.NextIndex(), lineNum)
	sb := strings.Builder{}
	exit := false
	hasPoint := false

	for itr.HasNext() && !exit {
		ru = itr.Peek()

		switch {
		case ru == '.':
			if hasPoint {
				m := `Theres two decimal points in this number`
				return Symbol{}, errors.New(m)
			}
			hasPoint = true
			fallthrough
		case unicode.IsDigit(ru):
			fallthrough
		case ru == '_':
			sb.WriteRune(itr.Next())
		default:
			exit = true
		}
	}

	r.Val = sb.String()
	r.End = itr.NextIndex()
	return r, nil
}

// spaceSym handles symbols that start with a rune with the
// unicode whitespace property.
// I.e. any whitespace rune, whitespace may resolve into a:
// - meaningless symbol that can be ignored when parsing
func spaceSym(itr *StrItr, lineNum int) (Symbol, error) {

	ru := itr.Peek()
	if !itr.HasNext() || !unicode.IsSpace(ru) {
		m := "You can't call this function unless the iterators first rune is whitespace"
		sh.CompilerBug(lineNum, m)
	}

	r := initSym(itr.NextIndex(), lineNum)
	sb := strings.Builder{}

	for itr.HasNext() {
		ru = itr.Peek()

		if unicode.IsSpace(ru) {
			sb.WriteRune(itr.Next())
		} else {
			break
		}
	}

	r.Val = sb.String()
	r.End = itr.NextIndex()
	return r, nil
}

// curseSym handles symbols that start with a at sign rune `@`.
// Curse symbols may resolve into a:
// - go function call
func curseSym(itr *StrItr, lineNum int) (Symbol, error) {

	if !itr.HasNext() || itr.Peek() != '@' {
		m := "You can't call this function unless the iterators first rune is an `@`"
		sh.CompilerBug(lineNum, m)
	}

	switch {
	case !itr.HasAsatte():
		fallthrough
	case !unicode.IsLetter(itr.PeekAsatte()):
		m := "You can't call this function unless the iterators first rune is an `@`"
		sh.SyntaxErr(lineNum, itr.NextIndex()+1, itr.Length(), m)
	}

	index := itr.NextIndex()
	first := string(itr.Next())
	r := extractWord(itr, lineNum)
	r.Start = index
	r.Val = first + r.Val

	return r, nil
}

// strSym handles symbols that start with the double quote `"` rune.
// Quoted strings may resolve into a:
// - string literal
func strSym(itr *StrItr, lineNum int) (Symbol, error) {

	if !itr.HasNext() || itr.Peek() != '"' {
		m := "You can't call this function unless the iterators first rune is `\"`"
		sh.CompilerBug(lineNum, m)
	}

	r := initSym(itr.NextIndex(), lineNum)
	sb := strings.Builder{}

	isEscaped := false
	isFirst := true

	for itr.HasNext() {
		ru := itr.Next()
		sb.WriteRune(ru)

		switch {
		case !itr.HasNext():
			break
		case isFirst:
			isFirst = false
		case ru == '\\':
			isEscaped = !isEscaped
		case !isEscaped && ru == '"':
			break
		default:
			isEscaped = false
		}
	}

	if isFirst || isEscaped || itr.PeekPrev() != '"' {
		return Symbol{}, errors.New("Did someone forget to close a string literal?!")
	}

	r.Val = sb.String()
	r.End = itr.NextIndex()
	return r, nil
}

// isComment return true if the rest of the string is a comment.
func isComment(itr *StrItr) bool {
	// TODO
	return false
}

// commentSym handles symbols that start with two forward slashes
// `//`. Double forward slashes may resolve into a:
// - comment
func commentSym(itr *StrItr, lineNum int) (Symbol, error) {
	// TODO
	return Symbol{}, errors.New("TODO: Implement this function!")
}

// otherSym handles any symbols that don't have a specific handling
// function. These symbols may resolve into a:
// - operator, 1 or 2 runes including truthy and not
// - code block start or end, i.e. bracket
// - value separator, i.e. comma
// - key-value separator, i.e. colon
// - void value, i.e. underscore
func otherSym(itr *StrItr, lineNum int) (Symbol, error) {
	// TODO
	return Symbol{}, errors.New("TODO: Implement this function!")
}

// extractWord iterates a string iterator until a single word has been
// extracted.
func extractWord(itr *StrItr, lineNum int) Symbol {
	r := initSym(itr.NextIndex(), lineNum)
	sb := strings.Builder{}
	exit := false

	for itr.HasNext() && !exit {
		ru := itr.Peek()

		switch {
		case unicode.IsLetter(ru):
			fallthrough
		case unicode.IsDigit(ru):
			fallthrough
		case ru == '_':
			sb.WriteRune(itr.Next())
		default:
			exit = true
		}
	}

	r.Val = sb.String()
	r.End = itr.NextIndex()
	return r
}
