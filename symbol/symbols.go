package symbol

import (
	"fmt"
	"strconv"
	"strings"
)

// SymbolType represents the type of a symbol.
type SymbolType int

const (
	UNDEFINED SymbolType = iota
	// Fully or partly alphabetic
	ALPHABETIC_START
	KEYWORD_SPELL // spell
	KEYWORD_LOOP  // loop
	KEYWORD_WHEN  // when
	KEYWORD_END   // end
	KEYWORD_KEY   // key
	KEYWORD_VALUE // value
	// Identifiers
	IDENTIFIER_EXPLICIT // Declared by the coder
	IDENTIFIER_IMPLICIT // Inserted during syntax analysis
	// Booleans
	BOOLEAN_TRUE  // true
	BOOLEAN_FALSE // false
	SOURCERY      // @Blahblah
	ALPHABETIC_END
	// Literals
	LITERAL_NUMBER // ##.###
	LITERAL_STRING // "blahblah"
	COMMENT        // // blahblah
	// Comparison operators
	CMP_EQUAL                 // ==
	CMP_NOT_EQUAL             // !=
	CMP_LESS_THAN             // <
	CMP_LESS_THAN_OR_EQUAL    // <=
	CMP_GREATER_THAN          // >
	CMP_GREATER_THAN_OR_EQUAL // >=
	// Logical operators
	LOGICAL_OR    // ||
	LOGICAL_AND   // &&
	LOGICAL_NOT   // !
	LOGICAL_MATCH // =>
	// Arithmetic operators
	CALC_ADD      // +
	CALC_SUBTRACT // -
	CALC_MULTIPLY // *
	CALC_DIVIDE   // /
	CALC_MODULO   // %
	// Context boundries
	PAREN_CURVY_OPEN   // (
	PAREN_CURVY_CLOSE  // )
	PAREN_SQUARE_OPEN  // [
	PAREN_SQUARE_CLOSE // ]
	// Punctuation
	ASSIGNMENT          // <-
	SEPARATOR_VALUE     // ,
	SEPARATOR_KEY_VALUE // :
	RANGE               // ..
	// Whitespace
	WHITESPACE
	// Ignoramuses
	VOID // _
)

// SymbolName returns the name of the symbol type.
func SymbolName(t SymbolType) string {
	switch t {
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
	case KEYWORD_VALUE:
		return `KEYWORD_VALUE`
	case IDENTIFIER_EXPLICIT:
		return `IDENTIFIER_EXPLICIT`
	case IDENTIFIER_IMPLICIT:
		return `IDENTIFIER_IMPLICIT`
	case BOOLEAN_TRUE:
		return `BOOLEAN_TRUE`
	case BOOLEAN_FALSE:
		return `BOOLEAN_FALSE`
	case SOURCERY:
		return `SOURCERY`
	case LITERAL_NUMBER:
		return `LITERAL_NUMBER`
	case LITERAL_STRING:
		return `LITERAL_STRING`
	case COMMENT:
		return `COMMENT`
	case CMP_EQUAL:
		return `CMP_EQUAL`
	case CMP_NOT_EQUAL:
		return `CMP_NOT_EQUAL`
	case CMP_LESS_THAN:
		return `CMP_LESS_THAN`
	case CMP_LESS_THAN_OR_EQUAL:
		return `CMP_LESS_THAN_OR_EQUAL`
	case CMP_GREATER_THAN:
		return `CMP_GREATER_THAN`
	case CMP_GREATER_THAN_OR_EQUAL:
		return `CMP_GREATER_THAN_OR_EQUAL`
	case LOGICAL_OR:
		return `OR`
	case LOGICAL_AND:
		return `AND`
	case LOGICAL_NOT:
		return `NOT`
	case LOGICAL_MATCH:
		return `LOGICAL_MATCH`
	case CALC_ADD:
		return `CALC_ADD`
	case CALC_SUBTRACT:
		return `CALC_SUBTRACT`
	case CALC_MULTIPLY:
		return `CALC_MULTIPLY`
	case CALC_DIVIDE:
		return `CALC_DIVIDE`
	case CALC_MODULO:
		return `CALC_MODULO`
	case PAREN_CURVY_OPEN:
		return `PAREN_CURVY_OPEN`
	case PAREN_CURVY_CLOSE:
		return `PAREN_CURVY_CLOSE`
	case PAREN_SQUARE_OPEN:
		return `PAREN_SQUARE_OPEN`
	case PAREN_SQUARE_CLOSE:
		return `PAREN_SQUARE_CLOSE`
	case ASSIGNMENT:
		return `ASSIGNMENT`
	case SEPARATOR_VALUE:
		return `SEPARATOR_VALUE`
	case SEPARATOR_KEY_VALUE:
		return `SEPARATOR_KEY_VALUE`
	case RANGE:
		return `RANGE`
	case WHITESPACE:
		return `WHITESPACE`
	case VOID:
		return `VOID`
	}

	return `UNDEFINED`
}

// Symbol represents  a terminal or non-terminal symbol.
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
func printlnSymbols(syms []Symbol, f func(Symbol) string) {
	l := len(syms)
	if l == 0 {
		fmt.Println(`[]`)
		return
	}

	fmt.Print(`[`)
	for i, v := range syms {
		s := f(v)
		fmt.Print(s)

		if i < l-1 {
			fmt.Print(`, `)
		}
	}

	fmt.Println(`]`)
}
