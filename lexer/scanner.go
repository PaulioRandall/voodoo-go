package lexer

import (
	"strings"
	"unicode"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
)

// ScanLine scans a line and creates an array of symbols
// based on the grammer rules of the language.
func ScanLine(line string, lineNum int) (r []sym.Symbol, lxErr LexError) {

	if line == `` {
		r = emptyLineSyms(lineNum)
		return
	}

	itr := sh.NewRuneItr(line)

	for itr.HasNext() {
		var s *sym.Symbol

		switch {
		case itr.IsNextLetter():
			s, lxErr = wordSym(itr)
		case itr.IsNextDigit():
			s, lxErr = numSym(itr)
		case itr.IsNextSpace():
			s, lxErr = spaceSym(itr)
		case itr.IsNext('@'):
			s, lxErr = sourcerySym(itr)
		case itr.IsNext('"'):
			s, lxErr = strSym(itr)
		case itr.IsNextStr(`//`):
			s, lxErr = commentSym(itr)
		default:
			s, lxErr = otherSym(itr)
		}

		if lxErr != nil {
			lxErr.Line(lineNum)
			break
		}

		s.Line = lineNum
		r = append(r, *s)
	}

	return
}

// emptyLineSyms returns an array containing a single empty
// symbol that represents an empty line.
func emptyLineSyms(lineNum int) []sym.Symbol {
	return []sym.Symbol{
		sym.Symbol{
			Val:   ``,
			Start: 0,
			End:   0,
			Line:  lineNum,
		},
	}
}

// wordSym handles symbols that start with a unicode category L rune.
// I.e. a letter from any alphabet, a word may resolve into a:
// - variable name
// - keyword
// - boolean value (`true` or `false`)
func wordSym(itr *sh.RuneItr) (s *sym.Symbol, err LexError) {

	if !itr.IsNextLetter() {
		m := "Expected first rune to be a letter"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	str := extractWordStr(itr)
	symType := sym.UNDEFINED

	switch strings.ToLower(str) {
	case `scroll`:
		symType = sym.KEYWORD_SCROLL
	case `spell`:
		symType = sym.KEYWORD_SPELL
	case `loop`:
		symType = sym.KEYWORD_LOOP
	case `when`:
		symType = sym.KEYWORD_WHEN
	case `end`:
		symType = sym.KEYWORD_END
	case `key`:
		symType = sym.KEYWORD_KEY
	case `val`:
		symType = sym.KEYWORD_VAL
	case `true`:
		symType = sym.BOOLEAN
	case `false`:
		symType = sym.BOOLEAN
	default:
		symType = sym.VARIABLE
	}

	s = &sym.Symbol{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  symType,
	}

	return
}

// numSym handles symbols that start with a unicode category Nd rune.
// I.e. any number from 0 to 9, a number may resolve into a:
// - literal number
func numSym(itr *sh.RuneItr) (s *sym.Symbol, err LexError) {

	if !itr.IsNextDigit() {
		m := "Expected first rune to be a digit"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	n, err := extractNum(itr)
	if err != nil {
		return
	}

	s = &sym.Symbol{
		Val:   n,
		Start: start,
		End:   itr.Index(),
		Type:  sym.NUMBER,
	}

	return
}

// extractNum extracts a number, as a string, from the supplied
// iterator.
func extractNum(itr *sh.RuneItr) (string, LexError) {
	sb := strings.Builder{}

	for itr.HasNext() {
		if itr.IsNextStr(`..`) {
			break
		}

		if itr.IsNext('.') {
			sb.WriteRune(itr.NextRune())
			return extractFractional(itr, &sb)
		}

		if itr.IsNextDigit() || itr.IsNext('_') {
			sb.WriteRune(itr.NextRune())
		} else {
			break
		}
	}

	return sb.String(), nil
}

// extractFractional extracts the fractional part of a number,
// as a string, from the supplied iterator and returns the
// number including the integer part.
func extractFractional(itr *sh.RuneItr, sb *strings.Builder) (string, LexError) {
	for itr.HasNext() {
		if itr.IsNext('.') {
			m := "Numbers can't have two fractional parts"
			return "", NewLexError(m, itr.Index())
		}

		if itr.IsNextDigit() || itr.IsNext('_') {
			sb.WriteRune(itr.NextRune())
		} else {
			break
		}
	}

	return sb.String(), nil
}

// spaceSym handles symbols that start with a rune with the
// unicode whitespace property.
// I.e. any whitespace rune, whitespace may resolve into a:
// - meaningless symbol that can be ignored when parsing
func spaceSym(itr *sh.RuneItr) (s *sym.Symbol, err LexError) {

	if !itr.IsNextSpace() {
		m := "Expected first rune to be whitespace"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	sb := strings.Builder{}

	for itr.HasNext() {
		if !itr.IsNextSpace() {
			break
		}
		sb.WriteRune(itr.NextRune())
	}

	s = &sym.Symbol{
		Val:   sb.String(),
		Start: start,
		End:   itr.Index(),
		Type:  sym.WHITESPACE,
	}

	return
}

// sourcerySym handles symbols that start with a at sign rune `@`.
// Sourcery symbols may resolve into a:
// - go function call
func sourcerySym(itr *sh.RuneItr) (s *sym.Symbol, err LexError) {

	if !itr.IsNext('@') {
		m := "Expected first rune to be `@`"
		err = NewLexError(m, itr.Index())
		return
	}

	if !unicode.IsLetter(itr.PeekRelRune(1)) {
		m := "Expected first rune after `@` to be a letter"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	firstLetter := string(itr.NextRune())
	val := firstLetter + extractWordStr(itr)

	s = &sym.Symbol{
		Val:   val,
		Start: start,
		End:   itr.Index(),
		Type:  sym.SOURCERY,
	}

	return
}

// strSym handles symbols that start with the double quote `"` rune.
// Quoted strings may resolve into a:
// - string literal
func strSym(itr *sh.RuneItr) (s *sym.Symbol, err LexError) {

	if !itr.IsNext('"') {
		m := "Expected first rune to be `\"`"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	closed, str := extractStr(itr)

	if !closed {
		m := "Did someone forget to close a string literal?!"
		err = NewLexError(m, itr.Index())
		return
	}

	s = &sym.Symbol{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  sym.STRING,
	}

	return
}

// extractStr extracts a string literal from a string iterator
// returning true if the last rune was escaped.
func extractStr(itr *sh.RuneItr) (closed bool, s string) {

	sb := strings.Builder{}
	sb.WriteRune(itr.NextRune())
	isEscaped := false

	for itr.HasNext() {
		ru := itr.NextRune()
		sb.WriteRune(ru)

		if !isEscaped && ru == '"' {
			closed = true
			break
		}

		if ru == '\\' {
			isEscaped = !isEscaped
		} else {
			isEscaped = false
		}
	}

	s = sb.String()
	return
}

// commentSym handles symbols that start with two forward slashes
// `//`. Double forward slashes may resolve into a:
// - comment
func commentSym(itr *sh.RuneItr) (s *sym.Symbol, err LexError) {

	if !itr.IsNextStr(`//`) {
		m := "Expected first two runes to be `//`"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	str := itr.RemainingStr()

	s = &sym.Symbol{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  sym.COMMENT,
	}

	return
}

// otherSym handles any symbols that don't have a specific handling
// function. These symbols may resolve into a:
// - operator, 1 or 2 runes including truthy and not
// - code block start or end, i.e. bracket
// - value separator, i.e. comma
// - key-value separator, i.e. colon
// - void value, i.e. underscore
func otherSym(itr *sh.RuneItr) (s *sym.Symbol, err LexError) {

	if !itr.HasNext() {
		m := "Expected an unfinished iterator"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	symType := sym.UNDEFINED
	runeCount := 0

	set := func(t sym.SymbolType, totalRunes int) {
		symType = t
		runeCount = totalRunes
	}

	switch {
	case itr.IsNextStr(`<-`):
		set(sym.ASSIGNMENT, 2)
	case itr.IsNextStr(`<=`):
		set(sym.LESS_THAN_OR_EQUAL, 2)
	case itr.IsNext('<'):
		set(sym.LESS_THAN, 1)
	case itr.IsNextStr(`>=`):
		set(sym.GREATER_THAN_OR_EQUAL, 2)
	case itr.IsNext('>'):
		set(sym.GREATER_THAN, 1)
	case itr.IsNextStr(`==`):
		set(sym.EQUAL, 2)
	case itr.IsNextStr(`=>`):
		set(sym.IF_TRUE_THEN, 2)
	case itr.IsNextStr(`!=`):
		set(sym.NOT_EQUAL, 2)
	case itr.IsNext('!'):
		set(sym.NEGATION, 1)
	case itr.IsNextStr(`||`):
		set(sym.OR, 2)
	case itr.IsNextStr(`&&`):
		set(sym.AND, 2)
	case itr.IsNext('+'):
		set(sym.ADD, 1)
	case itr.IsNext('-'):
		set(sym.SUBTRACT, 1)
	case itr.IsNext('*'):
		set(sym.MULTIPLY, 1)
	case itr.IsNext('/'):
		set(sym.DIVIDE, 1)
	case itr.IsNext('%'):
		set(sym.MODULO, 1)
	case itr.IsNext('('):
		set(sym.CIRCLE_BRACE_OPEN, 1)
	case itr.IsNext(')'):
		set(sym.CIRCLE_BRACE_CLOSE, 1)
	case itr.IsNext('['):
		set(sym.SQUARE_BRACE_OPEN, 1)
	case itr.IsNext(']'):
		set(sym.SQUARE_BRACE_CLOSE, 1)
	case itr.IsNext(','):
		set(sym.VALUE_SEPARATOR, 1)
	case itr.IsNext(':'):
		set(sym.KEY_VALUE_SEPARATOR, 1)
	case itr.IsNextStr(`..`):
		set(sym.RANGE, 2)
	case itr.IsNext('_'):
		set(sym.VOID, 1)
	default:
		ru := itr.NextRune()
		m := "I don't know what this symbol means '" + string(ru) + "'"
		err = NewLexError(m, itr.Index())
		return
	}

	str, e := itr.NextStr(runeCount)
	if e != nil {
		err = NewLexError(err.Error(), itr.Index())
		return
	}

	s = &sym.Symbol{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  symType,
	}

	return
}

// extractWordStr iterates a rune iterator until a single word has
// been extracted retruning the string.
func extractWordStr(itr *sh.RuneItr) string {
	sb := strings.Builder{}

	for itr.HasNext() {
		switch {
		case itr.IsNextLetter():
			fallthrough
		case itr.IsNextDigit():
			fallthrough
		case itr.IsNext('_'):
			sb.WriteRune(itr.NextRune())
		default:
			return sb.String()
		}
	}

	return sb.String()
}
