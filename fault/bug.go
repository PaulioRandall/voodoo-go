package fault

import (
	"fmt"
	"unicode"

	"github.com/PaulioRandall/voodoo-go/scroll"
)

// Bug represents a fault with this program.
type Bug struct {
	Stage  string // Stage or phase were the bug originated
	Intent string // What was the intended action or result
	Actual string // What was the actual action or result
}

// Print satisfies the Fault interface.
func (err Bug) Print(sc *scroll.Scroll, line int) {
	fmt.Println("\n[BUG]")

	fmt.Printf("%3d: %s\n", line, sc.Lines[line])

	fmt.Println("During:")
	printPara([]rune(err.Stage))

	fmt.Println("Intent:")
	printPara([]rune(err.Intent))

	fmt.Println("Actual:")
	printPara([]rune(err.Intent))
}

// printPara prints a paragraph where by each line has a limited
// rune count and indented with a single tab.
func printPara(in []rune) {
	maxSize := 72
	size := maxSize

	printIndent := func(s string) {
		if size >= maxSize {
			size = len(s)
			fmt.Print("\n\t")
		} else {
			fmt.Print(" ")
		}
	}

	for _, s := range splitSpace(in) {
		size += len(s) + 1
		printIndent(s)
		fmt.Print(s)
	}

	fmt.Println()
}

// splitSpace splits a string on the whitespace removing
// empty and whitespace tokens.
func splitSpace(in []rune) []string {
	out := []string{}
	f := 0

	doAppend := func(i int) {
		if f+1 < i {
			out = append(out, string(in[f:i]))
		}
	}

	for i, r := range in {
		if unicode.IsSpace(r) {
			doAppend(i)
			f = i
		}
	}

	return out
}
