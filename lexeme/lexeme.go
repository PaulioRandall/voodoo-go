package lexeme

import (
	"fmt"
	"strconv"
	"strings"
)

// LexemeType represents the type of the lexeme.
type LexemeType int

const (
	UNDEFINED LexemeType = iota
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
	BOOLEAN_TRUE  // true
	BOOLEAN_FALSE // false
	SOURCERY      // @Blahblah
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
	OR           // ||
	AND          // &&
	NEGATION     // !
	IF_TRUE_THEN // =>
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

// nameOfType returns the name of the lexeme type.
func nameOfType(t LexemeType) string {
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
	case IF_TRUE_THEN:
		return `IF_TRUE_THEN`
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

// Lexeme represents a rune or string within the code
// that equates to a meaningful item within the grammer
// rules.
type Lexeme struct {
	Val   string     // Lexeme value
	Start int        // Index of first rune
	End   int        // Index after last rune
	Line  int        // Line number from scroll
	Type  LexemeType // Type of lexeme
}

// String creates a string representation of the lexeme.
func (lex Lexeme) String() string {
	start := strconv.Itoa(lex.Start)
	start = strings.Repeat(` `, 3-len(start)) + start
	return fmt.Sprintf("Line %-3d [%s->%-3d] `%s`", lex.Line, start, lex.End, lex.Val)
}

// PrintlnLexemes prints an array of lexemes.
func PrintlnLexemes(ls []Lexeme) {
	f := func(l Lexeme) string {
		return l.Val
	}
	printlnLexemesArray(ls, f)
}

// PrintlnLexemeTypes prints the types of an array of lexemes.
func PrintlnLexemeTypes(ls []Lexeme) {
	f := func(l Lexeme) string {
		return nameOfType(l.Type)
	}
	printlnLexemesArray(ls, f)
}

// printlnLexemesArray prints an array of lexemes where the
// value to print for each lexeme is obtained via a the
// supplied function.
func printlnLexemesArray(ls []Lexeme, f func(l Lexeme) string) {
	l := len(ls)
	if l == 0 {
		fmt.Println(`[]`)
		return
	}

	fmt.Print(`[`)
	for i, v := range ls {
		s := f(v)
		fmt.Print(s)
		if i < l-1 {
			fmt.Print(`, `)
		}
	}

	fmt.Println(`]`)
}
