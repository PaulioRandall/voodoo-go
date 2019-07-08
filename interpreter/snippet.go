
package interpreter

import (
	"regexp"
)

// Most operators must have whitespace on both sides and are only used for
// a single purpose, therefore, if the operator exists in a snippet it must
// be doing a specifically identifyable thing somewhere within it.

// Snippet represents a snippet of code that forms part of or all of
// a line from a scroll.
type Snippet struct {
	Code string
	Row int
	Col int
}

// Operator represents a arithmetic or boolean operator.
type Operator string

// BoolOperator represents a enum type of boolean operators.
type BoolOperator Operator

// MathOperator represents a enum type of maths operators.
type MathOperator Operator

const (
	BoolEqual BoolOperator = `\s==\s`
	BoolNotEqual BoolOperator = `\s!=\s`
	BoolLessThan BoolOperator = `\s<\s`
	BoolGreaterThan BoolOperator = `\s>\s`
	BoolLessThanOrEqual BoolOperator = `\s<=\s`
	BoolGreaterThanOrEqual BoolOperator = `\s>=\s`
	BoolAnd BoolOperator = `\s&&\s`
	BoolOr BoolOperator = `\s||\s`
	
	MathAdd MathOperator = `\s+\s`
	MathSub MathOperator = `\s-\s`
	MathMul MathOperator = `\s*\s`
	MathDiv MathOperator = `\s/\s`
)

// HasAssignOperator returns true if the snippet contains an assigning of a
// value or the result of an expression to a variable. 
func (snip Snippet) HasAssignOperator() bool {
  return snip.regex(`\s=\s`)
}

// HasBoolOperator returns true if the snippet contains a boolean operator. 
func (snip Snippet) HasBoolOperator(boolOp BoolOperator) bool {
	op := Operator(boolOp)
  return snip.regex(op)
}

// ContainsBoolOperator returns true if the snippet contains a boolean operator.
func (snip Snippet) ContainsBoolOperator() bool {
	switch {
	case snip.HasBoolOperator(BoolEqual),
		snip.HasBoolOperator(BoolNotEqual),
		snip.HasBoolOperator(BoolLessThan),
		snip.HasBoolOperator(BoolGreaterThan),
		snip.HasBoolOperator(BoolLessThanOrEqual),
		snip.HasBoolOperator(BoolGreaterThanOrEqual),
		snip.HasBoolOperator(BoolAnd),
		snip.HasBoolOperator(BoolOr):
		return true
	}
	
	return false
}

// HasMathOperator returns true if the snippet contains a arithmetic operator.
func (snip Snippet) HasMathOperator(mathOp MathOperator) bool {
	op := Operator(mathOp)
  return snip.regex(op)
}

// ContainsMathOperator returns true if the snippet contains an arithmetic operator.
func (snip Snippet) ContainsMathOperator() bool {
	switch {
	case snip.HasMathOperator(MathAdd),
		snip.HasMathOperator(MathSub),
		snip.HasMathOperator(MathMul),
		snip.HasMathOperator(MathDiv):
		return true
	}
	
	return false
}

// HasTruthyOperator returns true if the snippet contains a 'truthy' operator.
func (snip Snippet) HasTruthyOperator() bool {
	return snip.regex(`?[\)|\s]`)
}

// HasNotOperator returns true if the snippet contains a 'not' operator.
func (snip Snippet) HasNotOperator() bool {
	return snip.regex(`[\(|\s][nN][oO][tT][\(|\s]`)
}

// HasCalc returns true if the snippet contains some form of arithmetic
// or boolean calculation.
func (snip Snippet) HasCalc() bool {
	switch {
		case snip.ContainsBoolOperator(),
			snip.ContainsMathOperator(),
			snip.HasTruthyOperator(),
			snip.HasNotOperator():
			return true
	}
	
	return false
}

// regex performs a regular expression issuing a compiler bug if
// the regex is bad.
func (snip Snippet) regex(regex Operator) bool {
	reg := string(regex)
	r, err := regexp.MatchString(reg, snip.Code)
	snip.regexErr(err)
	return r
}

// regexErr handles rexexp errors when identifying expressions.
func (snip Snippet) regexErr(err error) {
	if err != nil {
		compilerBug(snip.Row, "bad regex used for identifying expression")
	}
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
	
	for i, r := range snip.Code {
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
