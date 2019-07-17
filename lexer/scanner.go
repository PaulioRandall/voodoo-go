package lexer

import (
	"strings"
	"unicode"

	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// ScanLine scans a line and creates an array of symbols
// based on the grammer rules of the language.
func ScanLine(line string, lineNum int) (r []Symbol, lxErr LexError) {

	if line == `` {
		r = emptyLineSyms(lineNum)
		return
	}

	itr := sh.NewRuneItr(line)

	for itr.HasNext() {
		var s *Symbol

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
func initSym(start int) *Symbol {
	return &Symbol{
		Start: start,
	}
}

// wordSym handles symbols that start with a unicode category L rune.
// I.e. a letter from any alphabet, a word may resolve into a:
// - variable name
// - keyword
// - boolean value (`true` or `false`)
func wordSym(itr *sh.RuneItr) (*Symbol, LexError) {

	if !itr.IsNextLetter() {
		m := "Expected first rune to be a letter"
		return nil, NewLexError(m, itr.Index())
	}

	s := extractWord(itr)

	switch strings.ToLower(s.Val) {
	case `scroll`:
		s.Type = KEYWORD_SCROLL
	case `spell`:
		s.Type = KEYWORD_SPELL
	case `loop`:
		s.Type = KEYWORD_LOOP
	case `when`:
		s.Type = KEYWORD_WHEN
	case `end`:
		s.Type = KEYWORD_END
	case `key`:
		s.Type = KEYWORD_KEY
	case `val`:
		s.Type = KEYWORD_VAL
	case `true`:
		s.Type = BOOLEAN
	case `false`:
		s.Type = BOOLEAN
	default:
		s.Type = VARIABLE
	}

	return s, nil
}

// numSym handles symbols that start with a unicode category Nd rune.
// I.e. any number from 0 to 9, a number may resolve into a:
// - literal number
func numSym(itr *sh.RuneItr) (*Symbol, LexError) {

	if !itr.IsNextDigit() {
		m := "Expected first rune to be a digit"
		return nil, NewLexError(m, itr.Index())
	}

	s := initSym(itr.Index())

	n, err := extractNum(itr)
	if err != nil {
		return nil, err
	}

	s.Val = n
	s.End = itr.Index()
	s.Type = NUMBER

	return s, nil
}

// extractNum extracts a number, as a string, from the supplied
// iterator.
func extractNum(itr *sh.RuneItr) (string, LexError) {
	sb := strings.Builder{}

	for itr.HasNext() {
		if itr.IsNextStr(`..`) {
			return sb.String(), nil
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
func spaceSym(itr *sh.RuneItr) (*Symbol, LexError) {

	if !itr.IsNextSpace() {
		m := "Expected first rune to be whitespace"
		return nil, NewLexError(m, itr.Index())
	}

	s := initSym(itr.Index())
	sb := strings.Builder{}

	for itr.HasNext() {
		if !itr.IsNextSpace() {
			break
		}

		sb.WriteRune(itr.NextRune())
	}

	s.Val = sb.String()
	s.End = itr.Index()
	s.Type = WHITESPACE

	return s, nil
}

// sourcerySym handles symbols that start with a at sign rune `@`.
// Sourcery symbols may resolve into a:
// - go function call
func sourcerySym(itr *sh.RuneItr) (*Symbol, LexError) {

	if !itr.IsNext('@') {
		m := "Expected first rune to be `@`"
		return nil, NewLexError(m, itr.Index())
	}

	if !unicode.IsLetter(itr.PeekRelRune(1)) {
		m := "Expected first rune after `@` to be a letter"
		return nil, NewLexError(m, itr.Index())
	}

	start := itr.Index()
	val := string(itr.NextRune())

	s := extractWord(itr)
	s.Start = start
	s.Val = val + s.Val
	s.Type = SOURCERY

	return s, nil
}

// strSym handles symbols that start with the double quote `"` rune.
// Quoted strings may resolve into a:
// - string literal
func strSym(itr *sh.RuneItr) (*Symbol, LexError) {

	if !itr.IsNext('"') {
		m := "Expected first rune to be `\"`"
		return nil, NewLexError(m, itr.Index())
	}

	s := initSym(itr.Index())
	isEscaped, str := extractStr(itr)

	if isEscaped || len(str) < 2 || str[len(str)-1] != '"' {
		m := "Did someone forget to close a string literal?!"
		return nil, NewLexError(m, itr.Index())
	}

	s.Val = str
	s.End = itr.Index()
	s.Type = STRING

	return s, nil
}

// extractStr extracts a string literal from a string iterator
// returning true if the last rune was escaped.
func extractStr(itr *sh.RuneItr) (bool, string) {

	sb := strings.Builder{}
	sb.WriteRune(itr.NextRune())
	isEscaped := false

	for itr.HasNext() {
		ru := itr.NextRune()
		sb.WriteRune(ru)

		switch {
		case !isEscaped && ru == '"':
			return false, sb.String()
		case ru == '\\':
			isEscaped = !isEscaped
		case itr.HasNext():
			isEscaped = false
		}
	}

	return isEscaped, sb.String()
}

// commentSym handles symbols that start with two forward slashes
// `//`. Double forward slashes may resolve into a:
// - comment
func commentSym(itr *sh.RuneItr) (*Symbol, LexError) {

	if !itr.IsNextStr(`//`) {
		m := "Expected first two runes to be `//`"
		return nil, NewLexError(m, itr.Index())
	}

	s := initSym(itr.Index())
	s.Val = itr.RemainingStr()
	s.End = itr.Index()
	s.Type = COMMENT

	return s, nil
}

// otherSym handles any symbols that don't have a specific handling
// function. These symbols may resolve into a:
// - operator, 1 or 2 runes including truthy and not
// - code block start or end, i.e. bracket
// - value separator, i.e. comma
// - key-value separator, i.e. colon
// - void value, i.e. underscore
func otherSym(itr *sh.RuneItr) (*Symbol, LexError) {

	if !itr.HasNext() {
		m := "Expected an unfinished iterator"
		return nil, NewLexError(m, itr.Index())
	}

	s := initSym(itr.Index())

	runeCount := 0
	set := func(t SymbolType, runesInOperator int) {
		s.Type = t
		runeCount = runesInOperator
	}

	switch {
	case itr.IsNextStr(`<-`):
		set(ASSIGNMENT, 2)
	case itr.IsNextStr(`<=`):
		set(LESS_THAN_OR_EQUAL, 2)
	case itr.IsNext('<'):
		set(LESS_THAN, 1)
	case itr.IsNextStr(`>=`):
		set(GREATER_THAN_OR_EQUAL, 2)
	case itr.IsNext('>'):
		set(GREATER_THAN, 1)
	case itr.IsNextStr(`==`):
		set(EQUAL, 2)
	case itr.IsNextStr(`=>`):
		set(IF_TRUE_THEN, 2)
	case itr.IsNextStr(`!=`):
		set(NOT_EQUAL, 2)
	case itr.IsNext('!'):
		set(NEGATION, 1)
	case itr.IsNextStr(`||`):
		set(OR, 2)
	case itr.IsNextStr(`&&`):
		set(AND, 2)
	case itr.IsNext('+'):
		set(ADD, 1)
	case itr.IsNext('-'):
		set(SUBTRACT, 1)
	case itr.IsNext('*'):
		set(MULTIPLY, 1)
	case itr.IsNext('/'):
		set(DIVIDE, 1)
	case itr.IsNext('%'):
		set(MODULO, 1)
	case itr.IsNext('('):
		set(CIRCLE_BRACE_OPEN, 1)
	case itr.IsNext(')'):
		set(CIRCLE_BRACE_CLOSE, 1)
	case itr.IsNext('['):
		set(SQUARE_BRACE_OPEN, 1)
	case itr.IsNext(']'):
		set(SQUARE_BRACE_CLOSE, 1)
	case itr.IsNext(','):
		set(VALUE_SEPARATOR, 1)
	case itr.IsNext(':'):
		set(KEY_VALUE_SEPARATOR, 1)
	case itr.IsNextStr(`..`):
		set(RANGE, 2)
	case itr.IsNext('_'):
		set(VOID, 1)
	default:
		ru := itr.NextRune()
		m := "I don't know what this symbol means '" + string(ru) + "'"
		return nil, NewLexError(m, itr.Index())
	}

	str, err := itr.NextStr(runeCount)
	if err != nil {
		lxErr := NewLexError(err.Error(), itr.Index())
		return nil, lxErr
	}

	s.Val = str
	s.End = itr.Index()

	return s, nil
}

// extractWord iterates a rune iterator until a single word has been
// extracted returning it as a symbol.
func extractWord(itr *sh.RuneItr) *Symbol {
	s := initSym(itr.Index())
	s.Val = extractWordStr(itr)
	s.End = itr.Index()
	return s
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
