
package cleaver

import (
	"fmt"
	"strings"
	"strconv"
)

// The purpose of the welder is to take the Fragments of Elements
// produced by the cleaver and use, but not enforce, some of the
// language grammer rules to produce Elements. Some Fragments map
// directly to Elements and some are 'welded' (joined) together to
// form them.
//
// TODO: Create SQL flow diagrams of these!
//
// Elements that require welding include:
// - String literals: anything between double quotes `"`.
// - Number literals: number folowed by numbers and underscores
//										(optional) followed by a single dot then at least one number
//										(optional) followed by numbers and underscores.
// - Words: A letter followed by letters, numbers and underscores.
// - Comments: Two forward slashes `//` followed by the rest of the fragments in the line.
//
// E.g. @Print("It's true!") will weld together [ `"`, `It`, `'`,
// `s`, ` `, `true`, `!`, `"` ] because they are all part of a string
// literal.
//
// Just like Fragments, Elements are a 'lossless' way to recreate the
// original scroll with the ommission of leading and trailing
// whitespace on each line of course, leading indentation being
// inferrable.

// Element represents the string that makes up a lexeme.
type Element struct {
	Val string
	Start int
	End int				// Exclusive
}

// String creates a string representation of the Element.
// TODO: Is this a duplicate of a Fragment?
func (elem Element) String() string {
	start := strconv.Itoa(elem.Start)
	start = strings.Repeat(` `, 3 - len(start)) + start
	return fmt.Sprintf("[%s:%-3d] `%s`", start, elem.End, elem.Val)
}

// PrintElements prints an array of Elements.
func PrintElements(elems []Element) {
	for _, v := range elems {
		fmt.Println(v)
	}
}

// fragItr allows a Fragment array to be iterated.
type fragItr struct {
	index int
	length int
	array []Fragment
}

// newFragItr creates a new Fragment iterator.
func newFragItr(frags []Fragment) *fragItr {
	return &fragItr{
		length: len(frags),
		array: frags,
	}
}

// hasNext returns true if the iterator has yet to be fully used.
func (itr *fragItr) hasNext() bool {
	i = itr.index + 1
	if itr.index < itr.length {
		return true
	}
	return false
}

// skip the next Fragment in the array and by incrementing the
// iterator index without returning anything.
func (itr *fragItr) skip() {
	bugIfNoNext()
	itr.index++
} 

// next returns the next Fragment in the array and increases the
// iterator index.
func (itr *fragItr) next() Fragment {
	defer itr.index++
	return itr.peek()
}

// peek returns the next Fragment in the array without incrementing
// the iterator index.
func (itr *fragItr) peek() Fragment {
	bugIfNoNext()
	i := itr.index + 1
	return itr.array[i]
}

// bugIfNoNext will print error message and exit compilation if there
// are no items lft in the array.
func (itr *fragItr) bugIfNoNext() {
	if !itr.hasNext() {
		sh.CompilerBug(-1, "Iterator call to next() or peek() but no items remain")
	}
}

// TODO: Create SQL flow diagram of what the welder does!
//
// Weld joins together Fragments into Elements so the next phase
// deals with units of statments and doesn't have to worry about
// splitting and joining strings.
func Weld(frags []Fragment) []Element {

	itr := newFragIter(frags)
	v := itr.peek()
	
	if len(v.Val) == 0 {
		return []Element{
			Element{}
		}
	}
	
	elems := []Element{}
	for itr.HasNext() {
		
		v = itr.Next()
		if len(v.Val) == 0 {
			result = append(result, Element{})
			break
		}
		
		r := rune(v.Val)
		var el Element
		
		switch {
		case isDoubleQuote(r):
			el = strElem(v, itr)			
		case isNumber(r):
			el = numElem(v, itr)
		case isLetter(r):
			el = wordElem(v, itr)
		default:
			if itr.hasNext() {
				n := itr.Peek()
				if isComment(v.Val + n.Val) {
					el = commentElem(v, itr)
				} else {
					el = fragToElem(v)
				}
			}
		}
		
		elems = append(elems, e)
	}
	
	return elems
}

// strElem continues iterating the fragment array until a string
// literal element is produced.
func strElem(first Fragment, itr *fragItr) Element {
	elem := Element{
		Val: first.Val,
		Start: first.Start,
	}
	
	isEscaped = false
	for itr.HasNext() {
		v := itr.Next()
		elem.End = v.End
		
		s := v.Val
		elem.Val = elem.Val + s
		
		switch {
		case s == `\`:
			isEscaped = !isEscaped
		case !isEscaped && s == `"`:
			return elem
		default:
			isEscaped = false
		}
	}
	
	return elem
}

// numElem continues iterating the fragment array until a number
// literal element is produced.
func numElem(first Fragment, itr *fragItr) Element {
	elem := Element{
		Val: first.Val,
		Start: first.Start,
	}
	
	update = func(f Fragment) {
		elem.Val = elem.Val + f.Val
		elem.End = f.End
		itr.skip()
	}
	
	for itr.HasNext() {
		v := itr.Peek()
		s := v.Val
		
		switch {
		case isNumStr(s):
			update(v)
		case isPoint(s[0]):
			update(v)
		case isUnderscore(s[0]):
			update(v)
		default:
			break
		}
	}
	
	return elem
}

// wordElem continues iterating the fragment array until a word
// element is produced.
func wordElem(first Fragment, itr *fragItr) Element {
	elem := Element{
		Val: first.Val,
		Start: first.Start,
	}
	
	update = func(f Fragment) {
		elem.Val = elem.Val + f.Val
		elem.End = f.End
		itr.skip()
	}
	
	for itr.HasNext() {
		v := itr.Peek()
		s := v.Val
		
		switch {
		case isNumStr(s):
			update(v)
		case isLetterStr(s):
			update(v)
		case isUnderscore(s[0]):
			update(v)
		default:
			break
		}
	}
	
	return elem
}

// commentElem continues iterating the fragment array until a
// comment element is produced.
func commentElem(first Fragment, itr *fragItr) Element {
	elem := Element{
		Val: first.Val,
		Start: first.Start,
	}
	
	var v Fragment
	for itr.HasNext() {
		v = itr.next()
		elem.Val = elem.Val + v.Val
	}
	elem.End = v.End
	
	return elem
}

// fragToElem converts a Fragment into an Element.
func fragToElem(f Fragment) Element {
	return Element{
		Val: f.Val,
		Start: f.Start,
		End: f.End,
	}
}

// JoinFragments joins together all Fragments in the passed
// array to form an Element.
func joinFragments(frags []Fragment) Element {
	elem := Element{}
	lastIndex := len(frags) - 1 
	
	for i, frag := range frags {
		if i == 0 {
			elem.Start = frag.Start
		} else if i == lastIndex {
			elem.End = frag.End
		}
		
		elem.Val = elem.Val + frag.Val
	}
	
	return elem
}

// isNewElem returns true if the Element, represented by
// an array of Fragments is new, array is empty.
func isNewElem(elem []Fragment) {
	return len(elem) == 0
}

// whenNewElem handles the first fragment in the new element. 
func whenNewElem(frag Fragment) (isComplete bool, isStrDelim bool) {
	s := frag.Val()
	l := len(s)
	
	switch {
	case len(s) == 1:
		isComplete, isStrDelim = whenSingleRune(rune(s[0]))
	default:
		elemComplete = false 
	}
	
	return
}

// whenSingleRune handles when the first fragment has a single rune.
func whenSingleRune(r rune) (isComplete bool, isStrDelim bool) {
	switch {
	case isWhitespace(r):
		elemComplete = true
	case isDoubleQuote(r):
		isStrDelim = true
	case isLetter(r):
	case isNumber(r):
	case isUnderscore(r):
	default:
		elemComplete = true
	}
	return
}
