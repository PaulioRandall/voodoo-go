package token

import (
	"fmt"
	"strconv"
	"strings"
)

var EMPTY Token = Token{}

// Token represents a token produced by lexical analysis.
// I.e. identifier, operator, punctionation, etc.
type Token struct {
	Val   string    // Token value
	Line  int       // Line number in scroll
	Start int       // Index of first rune
	End   int       // Index after last rune
	Type  TokenType // Type of token
}

// String creates a string representation of the token.
func (tk Token) String() string {
	start := strconv.Itoa(tk.Start)
	start = strings.Repeat(` `, 3-len(start)) + start
	return fmt.Sprintf("[%s->%-3d] `%s`", start, tk.End, tk.Val)
}

// PrintlnTokenChan prints each token arriving on the channel until the channel
// is closed. The value to print is obtained via calling the supplied function
// with each token.
func PrintlnTokenChan(done chan bool, in chan Token, f func(Token) string) {
	defer close(done)

	fmt.Println(`[`)
	newline := true

	for tk := range in {

		s := f(tk)

		if newline {
			newline = false
			fmt.Print(`  `)
			fmt.Print(s)

		} else if strings.ContainsRune(s, '\n') {
			fmt.Print(`, `)
			fmt.Print(s)
			newline = true

		} else {
			fmt.Print(`, `)
			fmt.Print(s)
		}
	}

	fmt.Println(`]`)
}
