package lexer

import (
	"errors"
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
		var s Symbol
		var err error

		switch {
		case itr.IsNextLetter():
			s, lxErr = wordSym(itr)
		case itr.IsNextDigit():
			s, lxErr = numSym(itr)
		case itr.IsNextSpace():
			s, lxErr = spaceSym(itr)
		case itr.IsNext('@'):
			s, err = sourcerySym(itr)
		case itr.IsNext('"'):
			s, err = strSym(itr)
		case itr.IsNextStr(`//`):
			s, err = commentSym(itr)
		default:
			s, err = otherSym(itr)
		}

		if err != nil {
			lxErr = NewLexError(err.Error(), -1)
		}

		if lxErr != nil {
			lxErr.Line(lineNum)
			break
		}

		s.Line = lineNum
		r = append(r, s)
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
func initSym(start int) Symbol {
	return Symbol{
		Start: start,
	}
}

// wordSym handles symbols that start with a unicode category L rune.
// I.e. a letter from any alphabet, a word may resolve into a:
// - variable name
// - keyword
// - boolean value (`true` or `false`)
func wordSym(itr *sh.RuneItr) (Symbol, LexError) {

	if !itr.IsNextLetter() {
		m := "Expected first rune to be a letter"
		return Symbol{}, NewLexError(m, itr.Index())
	}

	r := extractWord(itr)

	switch strings.ToLower(r.Val) {
	case `scroll`:
		r.Type = KEYWORD_SCROLL
	case `spell`:
		r.Type = KEYWORD_SPELL
	case `loop`:
		r.Type = KEYWORD_LOOP
	case `when`:
		r.Type = KEYWORD_WHEN
	case `end`:
		r.Type = KEYWORD_END
	case `key`:
		r.Type = KEYWORD_KEY
	case `val`:
		r.Type = KEYWORD_VAL
	case `true`:
		r.Type = BOOLEAN
	case `false`:
		r.Type = BOOLEAN
	default:
		r.Type = VARIABLE
	}

	return r, nil
}

// numSym handles symbols that start with a unicode category Nd rune.
// I.e. any number from 0 to 9, a number may resolve into a:
// - literal number
func numSym(itr *sh.RuneItr) (Symbol, LexError) {

	if !itr.IsNextDigit() {
		m := "Expected first rune to be a digit"
		return Symbol{}, NewLexError(m, itr.Index())
	}

	r := initSym(itr.Index())

	s, err := extractNum(itr)
	if err != nil {
		return Symbol{}, err
	}

	r.Val = s
	r.End = itr.Index()
	r.Type = NUMBER

	return r, nil
}

// extractNum extracts a number, as a string, from the supplied
// iterator.
func extractNum(itr *sh.RuneItr) (string, LexError) {
	sb := strings.Builder{}
	hasPoint := false

	for itr.HasNext() {
		switch {
		case hasPoint && itr.IsNext('.'):
			m := "Numbers can't have two fractional parts"
			return "", NewLexError(m, itr.Index())
		case itr.IsNext('.'):
			if itr.PeekRelRune(1) == '.' {
				return sb.String(), nil
			}
			hasPoint = true
			fallthrough
		case itr.IsNextDigit(), itr.IsNext('_'):
			sb.WriteRune(itr.NextRune())
		default:
			return sb.String(), nil
		}
	}

	return sb.String(), nil
}

// spaceSym handles symbols that start with a rune with the
// unicode whitespace property.
// I.e. any whitespace rune, whitespace may resolve into a:
// - meaningless symbol that can be ignored when parsing
func spaceSym(itr *sh.RuneItr) (Symbol, LexError) {

	if !itr.IsNextSpace() {
		m := "Expected first rune to be whitespace"
		return Symbol{}, NewLexError(m, itr.Index())
	}

	r := initSym(itr.Index())
	sb := strings.Builder{}

	for itr.HasNext() {
		if itr.IsNextSpace() {
			sb.WriteRune(itr.NextRune())
		} else {
			break
		}
	}

	r.Val = sb.String()
	r.End = itr.Index()
	r.Type = WHITESPACE

	return r, nil
}

// sourcerySym handles symbols that start with a at sign rune `@`.
// Sourcery symbols may resolve into a:
// - go function call
func sourcerySym(itr *sh.RuneItr) (Symbol, error) {

	if !itr.IsNext('@') {
		m := "Expected first rune to be `@`"
		return Symbol{}, errors.New(m)
	}

	if !unicode.IsLetter(itr.PeekRelRune(1)) {
		m := "Expected first rune after `@` to be a letter"
		return Symbol{}, errors.New(m)
	}

	start := itr.Index()
	val := string(itr.NextRune())
	r := extractWord(itr)

	r.Start = start
	r.Val = val + r.Val
	r.Type = SOURCERY

	return r, nil
}

// strSym handles symbols that start with the double quote `"` rune.
// Quoted strings may resolve into a:
// - string literal
func strSym(itr *sh.RuneItr) (Symbol, error) {

	if !itr.IsNext('"') {
		m := "Expected first rune to be `\"`"
		return Symbol{}, errors.New(m)
	}

	r := initSym(itr.Index())
	isEscaped, s := extractStr(itr)

	if isEscaped || len(s) < 2 || s[len(s)-1] != '"' {
		m := "Did someone forget to close a string literal?!"
		return Symbol{}, errors.New(m)
	}

	r.Val = s
	r.End = itr.Index()
	r.Type = STRING

	return r, nil
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
func commentSym(itr *sh.RuneItr) (Symbol, error) {

	if !itr.IsNextStr(`//`) {
		m := "Expected first two runes to be `//`"
		return Symbol{}, errors.New(m)
	}

	r := initSym(itr.Index())
	r.Val = itr.RemainingStr()
	r.End = itr.Index()
	r.Type = COMMENT

	return r, nil
}

// otherSym handles any symbols that don't have a specific handling
// function. These symbols may resolve into a:
// - operator, 1 or 2 runes including truthy and not
// - code block start or end, i.e. bracket
// - value separator, i.e. comma
// - key-value separator, i.e. colon
// - void value, i.e. underscore
func otherSym(itr *sh.RuneItr) (Symbol, error) {

	if !itr.HasNext() {
		m := "Expected an unfinished iterator"
		return Symbol{}, errors.New(m)
	}

	r := initSym(itr.Index())

	runeCount := 0
	set := func(t SymbolType, runesInOperator int) {
		r.Type = t
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
		return Symbol{}, errors.New(m)
	}

	s, err := itr.NextStr(runeCount)
	if err != nil {
		return Symbol{}, err
	}

	r.Val = s
	r.End = itr.Index()

	return r, nil
}

// extractWord iterates a rune iterator until a single word has been
// extracted returning it as a symbol.
func extractWord(itr *sh.RuneItr) Symbol {
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
