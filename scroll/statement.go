
package scroll

import (
	"regexp"
	
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// Statement represents a line or part of a line of code from a scroll.
// Expressions, operations, and code blocks are all types of statement.
type Statement struct {
	Val string
	Row int
	Col int
}

// IsAssignment returns true if the statement is an assignment.
func (stat Statement) IsAssignment() bool {
	return stat.regex(`=`)
}

// findIndex searches a statement to find whatever the 'onEachRune'
// function is searching for. Runes within string literals are not
// passed to the function so no special handling is required within
// the supplied function. 'onEachRune' function accepts the index
// as the first value and the rune as the second while a non-negative
// index if the required index has been found.
func (stat Statement) findIndex(onEachRune func(int, rune) int) int {
	inLiteral := false
	prev := ""
	
	for i, r := range stat.Val {
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

// TODO: Replace with new code in executors pkg

// regex performs a regular expression issuing a compiler bug if
// the regex is bad.
func (stat Statement) regex(regex string) bool {
	r, err := regexp.MatchString(regex, stat.Val)
	stat.regexErr(err)
	return r
}

// regexErr handles rexexp errors when identifying expressions.
func (stat Statement) regexErr(err error) {
	if err != nil {
		sh.CompilerBug(stat.Row, "bad regex used for identifying expression")
	}
}