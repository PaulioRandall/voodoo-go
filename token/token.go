package token

import (
	"fmt"
	"strconv"
	"strings"
)

// Token represents a token produced by lexical analysis.
// I.e. identifier, operator, punctionation, etc.
type Token struct {
	Val   string    // Token value
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

// PrintlnTokenValues prints the value of an array of tokens.
func PrintlnTokenValues(tks []Token) {
	f := func(tk Token) string {
		return tk.Val
	}
	printlnTokens(tks, f)
}

// PrintlnTokenTypes prints the types of an array of tokens.
func PrintlnTokenTypes(tks []Token) {
	f := func(tk Token) string {
		return TokenName(tk.Type)
	}
	printlnTokens(tks, f)
}

// printlnTokens prints an array of tokens where the
// value to print for each token is obtained via the
// supplied function.
func printlnTokens(tks []Token, f func(Token) string) {
	l := len(tks)
	if l == 0 {
		fmt.Println(`[]`)
		return
	}

	fmt.Print(`[`)
	for i, v := range tks {
		s := f(v)
		fmt.Print(s)

		if i < l-1 {
			fmt.Print(`, `)
		}
	}

	fmt.Println(`]`)
}

// PrintlnTokenChan prints each token arriving on the channel until the channel
// is closed. The value to print is obtained via calling the supplied function
// with each token.
func PrintlnTokenChan(done chan bool, in chan Token, f func(Token) string) {
	defer close(done)

	fmt.Print(`[`)
	first := true

	for tk := range in {
		if !first {
			fmt.Print(`, `)
		}
		first = false

		s := f(tk)
		fmt.Print(s)
	}

	fmt.Println(`]`)
}
