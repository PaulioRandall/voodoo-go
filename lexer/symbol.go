package lexer

import (
	"fmt"
	"strconv"
	"strings"

	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// Symbol represents a rune or string within the code
// that equates to a meaningful item within the
// grammer rules.
type Symbol struct {
	Val   string // Symbol value
	Start int    // Index of first rune
	End   int    // Index after last rune
	Line  int    // Line number from scroll
}

// String creates a string representation of the symbol.
func (sym Symbol) String() string {
	start := strconv.Itoa(sym.Start)
	start = strings.Repeat(` `, 3-len(start)) + start
	return fmt.Sprintf("Line %-3d [%s->%-3d] `%s`", sym.Line, start, sym.End, sym.Val)
}

// Println prints the symbol on its own line.
func (sym Symbol) Println() {
	fmt.Println(sym)
}

// PrintlnSymbols prints an array of symbols.
func PrintlnSymbols(syms []Symbol) {
	for _, v := range syms {
		fmt.Println(v)
	}
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
	itr.bugIfNoNext()
	itr.increment()
}

// Next returns the next symbol in the array and increases the
// iterator index.
func (itr *SymItr) Next() Symbol {
	defer itr.increment()
	return itr.Peek()
}

// Peek returns the next symbol in the array without
// incrementing the iterator index.
func (itr *SymItr) Peek() Symbol {
	itr.bugIfNoNext()
	i := itr.index + 1
	return itr.syms[i]
}

// bugIfNoNext will print error message and exit compilation if there
// are no items lft in the array.
func (itr *SymItr) bugIfNoNext() {
	if !itr.HasNext() {
		sh.CompilerBug(-1, "Iterator call to next() or peek() but no items remain")
	}
}
