package lexer

import (
	"strings"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/lexeme"
	"github.com/PaulioRandall/voodoo-go/runer"
)

// ScanLine scans a line and creates an array of lexemes based on
// the grammer rules of the language. If the line is empty then an
// empty array wwill be returned. If the line contains whitespace
// only a single symbol with the whitespace type is returned.
// No panic is generated by the scanner so if a panic occurs it's
// either a system issue or a bug.
func ScanLine(line string, lineNum int) (ls []lexeme.Lexeme, lxErr LexError) {

	if line == `` {
		ls = []lexeme.Lexeme{}
		return
	}

	itr := runer.NewRuneItr(line)

	for itr.HasNext() {
		var l *lexeme.Lexeme

		switch {
		case itr.IsNextLetter():
			l = wordLex(itr)
		case itr.IsNextDigit():
			l, lxErr = numSym(itr)
		case itr.IsNextSpace():
			l = spaceSym(itr)
		case itr.IsNext('@'):
			l, lxErr = sourcerySym(itr)
		case itr.IsNext('"'):
			l, lxErr = strSym(itr)
		case itr.IsNextStr(`//`):
			l = commentLex(itr)
		default:
			l, lxErr = otherSym(itr)
		}

		if lxErr != nil {
			lxErr = ChangeLine(lxErr, lineNum)
			ls = nil
			break
		}

		l.Line = lineNum
		ls = append(ls, *l)
	}

	return
}

// wordLex handles lexemes that start with a unicode category L rune.
// I.e. a letter from any alphabet, a word may resolve into a:
// - variable name
// - keyword
// - boolean value (`true` or `false`)
func wordLex(itr *runer.RuneItr) *lexeme.Lexeme {

	start := itr.Index()
	s := extractWordStr(itr)
	t := lexeme.UNDEFINED

	switch strings.ToLower(s) {
	case `scroll`:
		t = lexeme.KEYWORD_SCROLL
	case `spell`:
		t = lexeme.KEYWORD_SPELL
	case `loop`:
		t = lexeme.KEYWORD_LOOP
	case `when`:
		t = lexeme.KEYWORD_WHEN
	case `end`:
		t = lexeme.KEYWORD_END
	case `key`:
		t = lexeme.KEYWORD_KEY
	case `val`:
		t = lexeme.KEYWORD_VAL
	case `true`:
		t = lexeme.BOOLEAN
	case `false`:
		t = lexeme.BOOLEAN
	default:
		t = lexeme.VARIABLE
	}

	return &lexeme.Lexeme{
		Val:   s,
		Start: start,
		End:   itr.Index(),
		Type:  t,
	}
}

// numSym handles symbols that start with a unicode category Nd rune.
// I.e. any number from 0 to 9, a number may resolve into a:
// - literal number
func numSym(itr *runer.RuneItr) (s *lexeme.Lexeme, err LexError) {

	start := itr.Index()
	n, err := extractNum(itr)
	if err != nil {
		return
	}

	s = &lexeme.Lexeme{
		Val:   n,
		Start: start,
		End:   itr.Index(),
		Type:  lexeme.NUMBER,
	}

	return
}

// extractNum extracts a number, as a string, from the supplied
// iterator.
func extractNum(itr *runer.RuneItr) (string, LexError) {
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
func extractFractional(itr *runer.RuneItr, sb *strings.Builder) (string, LexError) {
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
func spaceSym(itr *runer.RuneItr) *lexeme.Lexeme {

	start := itr.Index()
	sb := strings.Builder{}

	for itr.HasNext() {
		if !itr.IsNextSpace() {
			break
		}
		sb.WriteRune(itr.NextRune())
	}

	return &lexeme.Lexeme{
		Val:   sb.String(),
		Start: start,
		End:   itr.Index(),
		Type:  lexeme.WHITESPACE,
	}
}

// sourcerySym handles symbols that start with a at sign rune `@`.
// Sourcery symbols may resolve into a:
// - go function call
func sourcerySym(itr *runer.RuneItr) (s *lexeme.Lexeme, err LexError) {

	if !unicode.IsLetter(itr.PeekRelRune(1)) {
		m := "Expected first rune after `@` to be a letter"
		err = NewLexError(m, itr.Index())
		return
	}

	start := itr.Index()
	firstLetter := string(itr.NextRune())
	val := firstLetter + extractWordStr(itr)

	s = &lexeme.Lexeme{
		Val:   val,
		Start: start,
		End:   itr.Index(),
		Type:  lexeme.SOURCERY,
	}

	return
}

// strSym handles symbols that start with the double quote `"` rune.
// Quoted strings may resolve into a:
// - string literal
func strSym(itr *runer.RuneItr) (s *lexeme.Lexeme, err LexError) {

	start := itr.Index()
	closed, str := extractStr(itr)

	if !closed {
		m := "Did someone forget to close a string literal?!"
		err = NewLexError(m, itr.Index())
		return
	}

	s = &lexeme.Lexeme{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  lexeme.STRING,
	}

	return
}

// extractStr extracts a string literal from a string iterator
// returning true if the last rune was escaped.
func extractStr(itr *runer.RuneItr) (closed bool, s string) {

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

// commentLex handles lexemes that start with two forward slashes
// `//`. Double forward slashes may resolve into a:
// - comment
func commentLex(itr *runer.RuneItr) *lexeme.Lexeme {

	start := itr.Index()
	str := itr.RemainingStr()

	return &lexeme.Lexeme{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  lexeme.COMMENT,
	}
}

// otherSym handles any symbols that don't have a specific handling
// function. These symbols may resolve into a:
// - operator, 1 or 2 runes including truthy and not
// - code block start or end, i.e. bracket
// - value separator, i.e. comma
// - key-value separator, i.e. colon
// - void value, i.e. underscore
func otherSym(itr *runer.RuneItr) (s *lexeme.Lexeme, err LexError) {

	start := itr.Index()
	symType := lexeme.UNDEFINED
	runeCount := 0

	set := func(t lexeme.LexemeType, totalRunes int) {
		symType = t
		runeCount = totalRunes
	}

	switch {
	case itr.IsNextStr(`<-`):
		set(lexeme.ASSIGNMENT, 2)
	case itr.IsNextStr(`<=`):
		set(lexeme.LESS_THAN_OR_EQUAL, 2)
	case itr.IsNext('<'):
		set(lexeme.LESS_THAN, 1)
	case itr.IsNextStr(`>=`):
		set(lexeme.GREATER_THAN_OR_EQUAL, 2)
	case itr.IsNext('>'):
		set(lexeme.GREATER_THAN, 1)
	case itr.IsNextStr(`==`):
		set(lexeme.EQUAL, 2)
	case itr.IsNextStr(`=>`):
		set(lexeme.IF_TRUE_THEN, 2)
	case itr.IsNextStr(`!=`):
		set(lexeme.NOT_EQUAL, 2)
	case itr.IsNext('!'):
		set(lexeme.NEGATION, 1)
	case itr.IsNextStr(`||`):
		set(lexeme.OR, 2)
	case itr.IsNextStr(`&&`):
		set(lexeme.AND, 2)
	case itr.IsNext('+'):
		set(lexeme.ADD, 1)
	case itr.IsNext('-'):
		set(lexeme.SUBTRACT, 1)
	case itr.IsNext('*'):
		set(lexeme.MULTIPLY, 1)
	case itr.IsNext('/'):
		set(lexeme.DIVIDE, 1)
	case itr.IsNext('%'):
		set(lexeme.MODULO, 1)
	case itr.IsNext('('):
		set(lexeme.CURVED_BRACE_OPEN, 1)
	case itr.IsNext(')'):
		set(lexeme.CURVED_BRACE_CLOSE, 1)
	case itr.IsNext('['):
		set(lexeme.SQUARE_BRACE_OPEN, 1)
	case itr.IsNext(']'):
		set(lexeme.SQUARE_BRACE_CLOSE, 1)
	case itr.IsNext(','):
		set(lexeme.VALUE_SEPARATOR, 1)
	case itr.IsNext(':'):
		set(lexeme.KEY_VALUE_SEPARATOR, 1)
	case itr.IsNextStr(`..`):
		set(lexeme.RANGE, 2)
	case itr.IsNext('_'):
		set(lexeme.VOID, 1)
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

	s = &lexeme.Lexeme{
		Val:   str,
		Start: start,
		End:   itr.Index(),
		Type:  symType,
	}

	return
}

// extractWordStr iterates a rune iterator until a single word has
// been extracted retruning the string.
func extractWordStr(itr *runer.RuneItr) string {
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
