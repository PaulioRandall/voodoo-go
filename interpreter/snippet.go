
package interpreter

import (
	"regexp"
)

// Snippet represents a snippet of code from a single line.
type Snippet string

// isAssignment returns true if the snippet is assigning a value or the
// result of an expression to a variable. The assignment operator '='
// must have whitespace on both sides of it. It is only used for
// assignment, therefore, if the operator exists in a snippet it must be
// assigning something to a variable.
func (snip Snippet) isAssignment() (bool, error) {
  s := string(snip)
	return regexp.MatchString(`\s=\s`, s)
}

// findIndex searches a snippet to find whatever the 'onEachRune'
// function is searching for. Runes within string literals are not
// passed to the function so no special handling is required within
// the supplied function. 'onEachRune' function accepts the index
// as the first value and the rune as the second while a non-negative
// index if the required index has been found.
func (snip Snippet) findIndex(onEachRune func(int, rune) int) int {
	inLiteral := false
	prev := ""
	
	for i, r := range snip {
		s := string(r)
		wasInLiteral := false
		
		if inLiteral && prev != "\\" && s == "\"" {
			inLiteral = false
			wasInLiteral = true
		}
		
		if !inLiteral {
			index := onEachRune(i, r)
			if index > -1 {
				return index
			}
		}
		
		if !inLiteral && !wasInLiteral && s == "\"" {
			inLiteral = true
		}
		
		prev = s
	}
	
	return -1
}