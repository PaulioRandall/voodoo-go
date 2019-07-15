package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

// SymbolType represents the type of the symbol.
type SymbolType int

const (
	UNDEFINED SymbolType = iota
	// Composite
	KEYWORD_SCROLL // scroll
	KEYWORD_SPELL  // spell
	KEYWORD_LOOP   // loop
	KEYWORD_WHEN   // when
	KEYWORD_END    // end
	KEYWORD_KEY    // key
	KEYWORD_VAL    // value
	VARIABLE
	BOOLEAN // true/false
	NUMBER  // -##.###
	STRING  // "blahblah"
	COMMENT // // blahblah
	SPELL   // @Blahblah
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
	MULTIPLY // *
	DIVIDE   // /
	MODULO   // %
	// Brackets
	CIRCLE_BRACE_LEFT  // (
	CIRCLE_BRACE_RIGHT // )
	SQUARE_BRACE_LEFT  // [
	SQUARE_BRACE_RIGHT // ]
	// Separators
	VALUE_SEPARATOR     // ,
	KEY_VALUE_SEPARATOR // :
)

// Symbol represents a rune or string within the code
// that equates to a meaningful item within the
// grammer rules.
type Symbol struct {
	Val   string     // Symbol value
	Start int        // Index of first rune
	End   int        // Index after last rune
	Line  int        // Line number from scroll
	Type  SymbolType // Type of symbol
}

// String creates a string representation of the symbol.
func (sym Symbol) String() string {
	start := strconv.Itoa(sym.Start)
	start = strings.Repeat(` `, 3-len(start)) + start
	return fmt.Sprintf("Line %-3d [%s->%-3d] `%s`", sym.Line, start, sym.End, sym.Val)
}

// PrintlnSymbols prints an array of symbols.
func PrintlnSymbols(syms []Symbol) {
	l := len(syms)
	if l == 0 {
		fmt.Println(`[]`)
		return
	}

	fmt.Print(`[`)
	for i, v := range syms {
		fmt.Print(v.Val)
		if i < l-1 {
			fmt.Print(`, `)
		}
	}
	fmt.Println(`]`)
}

// SymItr provides a way to iterate symbol arrays.
type SymItr struct {
	index  int
	length int
	syms   []Symbol
}

// NewSymItr creates a new symbol iterator.
func NewSymItr(syms []Symbol) *SymItr {
	return &SymItr{
		length: len(syms),
		syms:   syms,
	}
}

// HasNext returns true if there are items still to be iterated.
func (itr *SymItr) HasNext() bool {
	i := itr.index + 1
	if i < itr.length {
		return true
	}
	return false
}

// increment increments the index counter.
func (itr *SymItr) increment() {
	itr.index = itr.index + 1
}

// Skip the next symbol in the array by incrementing the
// iterator index without returning anything.
func (itr *SymItr) Skip() {
	itr.increment()
}

// Next returns the next symbol in the array and increases the
// iterator index. Nil is returned if no symbols remain.
func (itr *SymItr) Next() *Symbol {
	defer itr.increment()
	return itr.Peek()
}

// Peek returns the next symbol in the array without
// incrementing the iterator index.
func (itr *SymItr) Peek() *Symbol {
	if itr.HasNext() {
		i := itr.index + 1
		return &itr.syms[i]
	}
	return nil
}
