package lexer

import (
	"errors"
	"strings"
	"unicode"

	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// ScanLine scans a line and creates an array of symbols
// based on the grammer rules of the language.
func ScanLine(line string, lineNum int) (r []Symbol, err error) {

	if line == `` {
		r = emptyLineSyms(lineNum)
		return
	}

	itr := NewStrItr(line)

	for itr.HasNext() {
		var s Symbol
		ru := itr.Peek()

		switch {
		case unicode.IsLetter(ru):
			s, err = wordSym(itr, lineNum)
		case unicode.IsDigit(ru):
			fallthrough
		case unicode.IsDigit(ru):
			s, err = numSym(itr, lineNum)
		case unicode.IsSpace(ru):
			s, err = spaceSym(itr, lineNum)
		case ru == '@':
			s, err = sourcerySym(itr, lineNum)
		case ru == '"':
			s, err = strSym(itr, lineNum)
		case isComment(itr):
			s, err = commentSym(itr, lineNum)
		default:
			s, err = otherSym(itr, lineNum)
		}

		if err != nil {
			break
		}

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
func initSym(start, lineNum int) Symbol {
	return Symbol{
		Start: start,
		Line:  lineNum,
	}
}

func (itr *StrItr) toRuneItr() *sh.RuneItr { // TEMP REFACTORING
	r := sh.NewRuneItr(itr.str) // TEMP REFACTORING
	r.SetIndex(itr.NextIndex()) // TEMP REFACTORING
	return r                    // TEMP REFACTORING
}

func (itr *StrItr) setIndex(i int) { // TEMP REFACTORING
	itr.index = i // TEMP REFACTORING
} // TEMP REFACTORING

// wordSym handles symbols that start with a unicode category L rune.
// I.e. a letter from any alphabet, a word may resolve into a:
// - variable name
// - keyword
// - boolean value (`true` or `false`)
func wordSym(sItr *StrItr, lineNum int) (Symbol, error) {

	itr := sItr.toRuneItr() // TEMP REFACTORING

	if !itr.IsNextLetter() {
		m := "Expected first rune to be a letter"
		return Symbol{}, errors.New(m)
	}

	r := extractWord(itr, lineNum)

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

	sItr.setIndex(itr.Index()) // TEMP REFACTORING
	return r, nil
}

// numSym handles symbols that start with a unicode category Nd rune.
// I.e. any number from 0 to 9, a number may resolve into a:
// - literal number
func numSym(itr *StrItr, lineNum int) (Symbol, error) {

	if !itr.HasNext() || !unicode.IsDigit(itr.Peek()) {
		m := "Can't call this function when the first rune is not digit"
		return Symbol{}, errors.New(m)
	}

	r := initSym(itr.NextIndex(), lineNum)
	sb := strings.Builder{}
	exit := false
	hasPoint := false

	onFinish := func() {
		r.Val = sb.String()
		r.End = itr.NextIndex()
		r.Type = NUMBER
	}

	for itr.HasNext() && !exit {
		ru := itr.Peek()

		switch {
		case ru == '.':
			if hasPoint {
				m := "Numbers can't have two fractional parts"
				return Symbol{}, errors.New(m)
			} else if itr.PeekAsatte() == '.' {
				onFinish()
				return r, nil
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

	onFinish()
	return r, nil
}

// spaceSym handles symbols that start with a rune with the
// unicode whitespace property.
// I.e. any whitespace rune, whitespace may resolve into a:
// - meaningless symbol that can be ignored when parsing
func spaceSym(itr *StrItr, lineNum int) (Symbol, error) {

	ru := itr.Peek()
	if !itr.HasNext() || !unicode.IsSpace(ru) {
		m := "Can't call this function when the first rune is not whitespace"
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
	r.Type = WHITESPACE
	return r, nil
}

// sourcerySym handles symbols that start with a at sign rune `@`.
// Sourcery symbols may resolve into a:
// - go function call
func sourcerySym(sItr *StrItr, lineNum int) (Symbol, error) {

	itr := sItr.toRuneItr() // TEMP REFACTORING

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
	r := extractWord(itr, lineNum)

	r.Start = start
	r.Val = val + r.Val
	r.Type = SOURCERY

	sItr.setIndex(itr.Index()) // TEMP REFACTORING
	return r, nil
}

// strSym handles symbols that start with the double quote `"` rune.
// Quoted strings may resolve into a:
// - string literal
func strSym(sItr *StrItr, lineNum int) (Symbol, error) {

	itr := sItr.toRuneItr() // TEMP REFACTORING

	if !itr.IsNext('"') {
		m := "Expected first rune to be `\"`"
		return Symbol{}, errors.New(m)
	}

	r := initSym(itr.Index(), lineNum)
	isEscaped, s := extractStr(itr)

	if isEscaped || len(s) < 2 || s[len(s)-1] != '"' {
		m := "Did someone forget to close a string literal?!"
		return Symbol{}, errors.New(m)
	}

	r.Val = s
	r.End = itr.Index()
	r.Type = STRING

	sItr.setIndex(itr.Index()) // TEMP REFACTORING
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

// isComment return true if the rest of the string is a comment.
func isComment(sItr *StrItr) bool {
	itr := sItr.toRuneItr() // TEMP REFACTORING
	if itr.IsNextStr(`//`) {
		return true
	}
	return false
}

// commentSym handles symbols that start with two forward slashes
// `//`. Double forward slashes may resolve into a:
// - comment
func commentSym(sItr *StrItr, lineNum int) (Symbol, error) {

	itr := sItr.toRuneItr() // TEMP REFACTORING

	if !itr.IsNextStr(`//`) {
		m := "Expected first two runes to be `//`"
		return Symbol{}, errors.New(m)
	}

	r := initSym(itr.Index(), lineNum)
	r.Val = itr.RemainingStr()
	r.End = itr.Index()
	r.Type = COMMENT

	sItr.setIndex(itr.Index()) // TEMP REFACTORING
	return r, nil
}

// otherSym handles any symbols that don't have a specific handling
// function. These symbols may resolve into a:
// - operator, 1 or 2 runes including truthy and not
// - code block start or end, i.e. bracket
// - value separator, i.e. comma
// - key-value separator, i.e. colon
// - void value, i.e. underscore
func otherSym(itr *StrItr, lineNum int) (Symbol, error) {

	if !itr.HasNext() {
		m := "Can't call this function with a finished iterator"
		return Symbol{}, errors.New(m)
	}

	r := initSym(itr.NextIndex(), lineNum)
	ru := itr.Next()
	hasTwoRunes := false

	ifNextIsInElse := func(s string, onTrue SymbolType, onFalse SymbolType) {
		if itr.NextIsIn(s) {
			r.Type = onTrue
			hasTwoRunes = true
		} else {
			r.Type = onFalse
		}
	}

	switch {
	case ru == '<':
		ifNextIsInElse(`-`, ASSIGNMENT, LESS_THAN)
		ifNextIsInElse(`=`, LESS_THAN_OR_EQUAL, r.Type)
	case ru == '>':
		ifNextIsInElse(`=`, GREATER_THAN_OR_EQUAL, GREATER_THAN)
	case ru == '=' && itr.NextIsIn(`=>`):
		ifNextIsInElse(`=`, EQUAL, IF_TRUE_THEN)
		ifNextIsInElse(`>`, IF_TRUE_THEN, r.Type)
	case ru == '!':
		ifNextIsInElse(`=`, NOT_EQUAL, NEGATION)
	case ru == '|' && itr.NextIsIn(`|`):
		r.Type = OR
		hasTwoRunes = true
	case ru == '&' && itr.NextIsIn(`&`):
		r.Type = AND
		hasTwoRunes = true
	case ru == '+':
		r.Type = ADD
	case ru == '-':
		r.Type = SUBTRACT
	case ru == '*':
		r.Type = MULTIPLY
	case ru == '/':
		r.Type = DIVIDE
	case ru == '%':
		r.Type = MODULO
	case ru == '(':
		r.Type = CIRCLE_BRACE_OPEN
	case ru == ')':
		r.Type = CIRCLE_BRACE_CLOSE
	case ru == '[':
		r.Type = SQUARE_BRACE_OPEN
	case ru == ']':
		r.Type = SQUARE_BRACE_CLOSE
	case ru == ',':
		r.Type = VALUE_SEPARATOR
	case ru == ':':
		r.Type = KEY_VALUE_SEPARATOR
	case ru == '.' && itr.NextIsIn(`.`):
		r.Type = RANGE
		hasTwoRunes = true
	case ru == '_':
		r.Type = VOID
	default:
		m := "I don't know what this symbol means '" + string(ru) + "'"
		return Symbol{}, errors.New(m)
	}

	if hasTwoRunes {
		r.Val = string(ru) + string(itr.Next())
	} else {
		r.Val = string(ru)
	}

	r.End = itr.NextIndex()
	return r, nil
}

// extractWord iterates a rune iterator until a single word has been
// extracted returning it as a symbol.
func extractWord(itr *sh.RuneItr, lineNum int) Symbol {
	r := initSym(itr.Index(), lineNum)
	r.Val = extractWordStr(itr)
	r.End = itr.Index()
	return r
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
