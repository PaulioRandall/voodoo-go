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
	// Composite
	KEYWORD_SCROLL // scroll
	KEYWORD_SPELL  // spell
	KEYWORD_LOOP   // loop
	KEYWORD_WHEN   // when
	KEYWORD_END    // end
	KEYWORD_KEY    // key
	KEYWORD_VAL    // value
	IDENTIFIER
	BOOLEAN  // true/false
	NUMBER   // ##.###
	STRING   // "blahblah"
	COMMENT  // // blahblah
	SOURCERY // @Blahblah
	// Misc
	WHITESPACE
	ASSIGNMENT // <-
	VOID       // _
	RANGE      // ..
	// Conditional
	IF_TRUE_THEN // =>
	NEGATION     // !
	// Boolean operators
	EQUAL                 // ==
	NOT_EQUAL             // !=
	LESS_THAN             // <
	LESS_THAN_OR_EQUAL    // <=
	GREATER_THAN          // >
	GREATER_THAN_OR_EQUAL // >=
	OR                    // ||
	AND                   // &&
	// Arithmetic operators
	ADD      // +
	SUBTRACT // -
	MULTIPLY // *
	DIVIDE   // /
	MODULO   // %
	// Brackets
	CURVED_BRACE_OPEN  // (
	CURVED_BRACE_CLOSE // )
	SQUARE_BRACE_OPEN  // [
	SQUARE_BRACE_CLOSE // ]
	// Separators
	VALUE_SEPARATOR     // ,
	KEY_VALUE_SEPARATOR // :
)

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
	l := len(ls)
	if l == 0 {
		fmt.Println(`[]`)
		return
	}

	fmt.Print(`[`)
	for i, v := range ls {
		fmt.Print(v.Val)
		if i < l-1 {
			fmt.Print(`, `)
		}
	}
	fmt.Println(`]`)
}
