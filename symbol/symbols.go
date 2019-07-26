package symbol

import (
	"fmt"
	"strconv"
	"strings"
)

// SymbolType represents the type of the symbol.
type SymbolType int

const (
	UNDEFINED SymbolType = iota
	// Fully or partly alphabetic
	ALPHABETIC_START
	KEYWORD_SCROLL // scroll
	KEYWORD_SPELL  // spell
	KEYWORD_LOOP   // loop
	KEYWORD_WHEN   // when
	KEYWORD_END    // end
	KEYWORD_KEY    // key
	KEYWORD_VAL    // value
	IDENTIFIER
	BOOLEAN_TRUE    // true
	BOOLEAN_FALSE   // false
	SOURCERY        // @Blahblah
	TEMP_IDENTIFIER // #blahblah
	ALPHABETIC_END
	// Fully or partly alphabetic but representation must reamin
	// as the user defined
	STRING  // "blahblah"
	COMMENT // // blahblah
	// Numeric
	NUMBER // ##.###
	// Comparison operators
	EQUAL                 // ==
	NOT_EQUAL             // !=
	LESS_THAN             // <
	LESS_THAN_OR_EQUAL    // <=
	GREATER_THAN          // >
	GREATER_THAN_OR_EQUAL // >=
	// Logical operators
	OR            // ||
	AND           // &&
	NEGATION      // !
	IF_MATCH_THEN // =>
	// Arithmetic operators
	ADD      // +
	SUBTRACT // -
	MULTIPLY // *
	DIVIDE   // /
	MODULO   // %
	// Context boundries
	CURVED_BRACE_OPEN  // (
	CURVED_BRACE_CLOSE // )
	SQUARE_BRACE_OPEN  // [
	SQUARE_BRACE_CLOSE // ]
	// Punctuation
	ASSIGNMENT          // <-
	VALUE_SEPARATOR     // ,
	KEY_VALUE_SEPARATOR // :
	RANGE               // ..
	// Whitespace
	WHITESPACE
	// Ignoramuses
	VOID // _
)

// nameOfType returns the name of the symbol type.
func nameOfType(t SymbolType) string {
	switch t {
	case KEYWORD_SCROLL:
		return `KEYWORD_SCROLL`
	case KEYWORD_SPELL:
		return `KEYWORD_SPELL`
	case KEYWORD_LOOP:
		return `KEYWORD_LOOP`
	case KEYWORD_WHEN:
		return `KEYWORD_WHEN`
	case KEYWORD_END:
		return `KEYWORD_END`
	case KEYWORD_KEY:
		return `KEYWORD_KEY`
	case KEYWORD_VAL:
		return `KEYWORD_VAL`
	case IDENTIFIER:
		return `IDENTIFIER`
	case BOOLEAN_TRUE:
		return `BOOLEAN_TRUE`
	case BOOLEAN_FALSE:
		return `BOOLEAN_FALSE`
	case SOURCERY:
		return `SOURCERY`
	case TEMP_IDENTIFIER:
		return `TEMP_IDENTIFIER`
	case STRING:
		return `STRING`
	case COMMENT:
		return `COMMENT`
	case NUMBER:
		return `NUMBER`
	case EQUAL:
		return `EQUAL`
	case NOT_EQUAL:
		return `NOT_EQUAL`
	case LESS_THAN:
		return `LESS_THAN`
	case LESS_THAN_OR_EQUAL:
		return `LESS_THAN_OR_EQUAL`
	case GREATER_THAN:
		return `GREATER_THAN`
	case GREATER_THAN_OR_EQUAL:
		return `GREATER_THAN_OR_EQUAL`
	case OR:
		return `OR`
	case AND:
		return `AND`
	case NEGATION:
		return `NEGATION`
	case IF_MATCH_THEN:
		return `IF_MATCH_THEN`
	case ADD:
		return `ADD`
	case SUBTRACT:
		return `SUBTRACT`
	case MULTIPLY:
		return `MULTIPLY`
	case DIVIDE:
		return `DIVIDE`
	case MODULO:
		return `MODULO`
	case CURVED_BRACE_OPEN:
		return `CURVED_BRACE_OPEN`
	case CURVED_BRACE_CLOSE:
		return `CURVED_BRACE_CLOSE`
	case SQUARE_BRACE_OPEN:
		return `SQUARE_BRACE_OPEN`
	case SQUARE_BRACE_CLOSE:
		return `SQUARE_BRACE_CLOSE`
	case ASSIGNMENT:
		return `ASSIGNMENT`
	case VALUE_SEPARATOR:
		return `VALUE_SEPARATOR`
	case KEY_VALUE_SEPARATOR:
		return `KEY_VALUE_SEPARATOR`
	case RANGE:
		return `RANGE`
	case WHITESPACE:
		return `WHITESPACE`
	case VOID:
		return `VOID`
	}

	return `UNDEFINED`
}

// Symbol represents a meaningful unit within the code.
// I.e. identifier, operator, punctionation, etc.
type Symbol struct {
	Val   string     // Symbol value
	Start int        // Index of first rune
	End   int        // Index after last rune
	Line  int        // Line number from scroll
	Type  SymbolType // Type of symbol
}

// String creates a string representation of the symbol.
func (s Symbol) String() string {
	start := strconv.Itoa(s.Start)
	start = strings.Repeat(` `, 3-len(start)) + start
	return fmt.Sprintf("Line %-3d [%s->%-3d] `%s`", s.Line, start, s.End, s.Val)
}

// printlnSymbols prints an array of symbols where the
// value to print for each symbol is obtained via the
// supplied function.
func printlnSymbols(ss []Symbol, f func(Symbol) string) {
	l := len(ss)
	if l == 0 {
		fmt.Println(`[]`)
		return
	}

	fmt.Print(`[`)
	for i, v := range ss {
		s := f(v)
		fmt.Print(s)
		if i < l-1 {
			fmt.Print(`, `)
		}
	}

	fmt.Println(`]`)
}
