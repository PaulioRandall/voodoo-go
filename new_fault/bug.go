package fault_new

import (
	"fmt"
	"unicode"
)

// Bug represents a fault with this program.
type Bug struct {
	Stage  string // Stage or phase were the bug originated
	Intent string // What was the intended action or result
	Actual string // What was the actual action or result
}

// Error satisfies the error interface.
func (err Bug) Error() string {
	return err.Actual
}

// Print satisfies the Fault interface.
func (err Bug) Print() {
	fmt.Println("\n[BUG]")

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
	var s string

	for {
		s, in = nextStr(in)
		size += len(s) + 1

		if size >= maxSize {
			size = len(s)
			fmt.Print("\n\t")
		} else {
			fmt.Print(" ")
		}

		fmt.Print(s)
	}

	fmt.Println()
}

// nextStr returns the next trimmed string in the rune slice.
func nextStr(in []rune) (string, []rune) {
	in = skipSpaces(in)
	s, out := nextWord(in)
	return string(s), out
}

// nextWord returns the next word in the rune slice.
func nextWord(in []rune) (word, out []rune) {
	for i, r := range in {
		if unicode.IsSpace(r) {
			word = in[:i]
			out = in[i:]
			return
		}
	}

	word = in
	out = []rune{}
	return
}

// skipSpaces returns a slice of the array without any
// preceeding unicode whitespace property runes.
func skipSpaces(in []rune) []rune {
	for i, r := range in {
		if !unicode.IsSpace(r) {
			return in[i:]
		}
	}
	return []rune{}
}
