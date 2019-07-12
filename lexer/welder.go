package lexer

import (
	"unicode"
)

// The purpose of the welder is to take the fragments of elements
// produced by the cleaver and produce elements. Some fragments
// map directly to elements and some are 'welded' (joined) together to
// form them.
//
// Elements that require welding include:
// - String literals: anything between double quotes `"`.
// - Number literals: number folowed by numbers and underscores
//										(optional) followed by a single dot then at least one number
//										(optional) followed by numbers and underscores.
// - Words: A letter followed by letters, numbers and underscores.
// - Comments: Two forward slashes `//` followed by the rest of the
//							 fragments in the line.
// - Whitespace: all whitespace runes that appear concurrently to
//							  each other should be combined into a single element.
//
// Just like fragments, elements are a 'lossless' way to recreate the
// original scroll with the ommission of leading and trailing
// whitespace on each line, leading indentation is inferrable.
//
// TODO: Create SQL flow diagrams of these!

// Weld joins together fragments of elements so the next phase
// deals with units of statments and doesn't have to worry about
// splitting and joining strings.
func Weld(frags []Symbol) []Symbol {

	itr := NewSymItr(frags)
	r := handleIfEmptyLine(itr)

	if r != nil {
		return r
	}

	r = []Symbol{}
	for itr.HasNext() {

		f := Symbol(itr.Next())
		ru := rune(f.Val[0])
		var el Symbol

		switch {
		case ru == '"':
			el = strElem(f, itr)
		case unicode.IsDigit(ru):
			el = numElem(f, itr)
		case unicode.IsLetter(ru):
			el = wordElem(f, itr)
		default:
			el = fragToElem(f)
		}

		r = append(r, el)
	}

	return r
}

// handleIfEmptyLine returns an element array if the fragment
// array represents an empty line, nil is returned if not.
func handleIfEmptyLine(itr *SymItr) []Symbol {
	f := itr.Peek()
	if len(f.Val) != 0 {
		return nil
	}

	return []Symbol{
		Symbol{
			Line: f.Line,
		},
	}
}

// initElem creates a new element with the value and
// start index initialised to that of the input fragment.
func initElem(f Symbol) Symbol {
	return Symbol{
		Val:   f.Val,
		Start: f.Start,
	}
}

// strElem continues iterating the fragment array until a string
// literal element is produced. Grammer rules not enforced here.
func strElem(f Symbol, itr *SymItr) Symbol {
	el := initElem(f)
	isEscaped := false

	for itr.HasNext() {
		f = itr.Next()
		s := f.Val
		el.Val = el.Val + s

		switch {
		case s == `\`:
			isEscaped = !isEscaped
		case !isEscaped && s == `"`:
			break
		default:
			isEscaped = false
		}
	}

	el.End = f.End
	return el
}

// numElem continues iterating the fragment array until a number
// literal element is produced. Grammer rules not enforced here.
func numElem(f Symbol, itr *SymItr) Symbol {
	el := initElem(f)

	update := func(f Symbol) {
		el.Val = el.Val + f.Val
		el.End = f.End
		itr.Skip()
	}

	for itr.HasNext() {
		f = itr.Peek()
		s := f.Val

		switch {
		case isDigitStr(s):
			update(f)
		case rune(s[0]) == '.':
			update(f)
		case rune(s[0]) == '_':
			update(f)
		default:
			break
		}
	}

	return el
}

// wordElem continues iterating the fragment array until a word
// element is produced.
func wordElem(f Symbol, itr *SymItr) Symbol {
	el := initElem(f)

	update := func(f Symbol) {
		el.Val = el.Val + f.Val
		el.End = f.End
		itr.Skip()
	}

	for itr.HasNext() {
		f = itr.Peek()
		s := f.Val

		switch {
		case isDigitStr(s):
			update(f)
		case isLetterStr(s):
			update(f)
		case rune(s[0]) == '_':
			update(f)
		default:
			break
		}
	}

	return el
}

// commentElem continues iterating the fragment array until a
// comment element is produced.
func commentElem(f Symbol, itr *SymItr) Symbol {
	el := initElem(f)

	for itr.HasNext() {
		f = itr.Next()
		el.Val = el.Val + f.Val
	}

	el.End = f.End
	return el
}

// fragToElem converts a fragment into an element.
func fragToElem(f Symbol) Symbol {
	return Symbol{
		Val:   f.Val,
		Start: f.Start,
		End:   f.End,
		Line:  f.Line,
	}
}
