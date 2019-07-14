package lexer

import (
	"fmt"
	"strconv"
	"strings"
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
