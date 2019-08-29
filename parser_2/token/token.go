package token

import (
	"strconv"
	"strings"
)

// Token represents a token produced by lexical analysis.
// I.e. identifier, operator, punctionation, etc.
type Token interface {

	// Text returns the textual representation of the token. For scanned tokens
	// this will always be the actual text that represents the token while others
	// may choose to use as they please.
	Text() string

	// Kind returns the type of token.
	Kind() Kind

	// String returns a string representation of the token.
	String() string
}

// scanTok is an implementation of Token produced by the scanner.
type scanTok struct {
	Val   string // Token value
	Line  int    // Line number in scroll
	Start int    // Index of first rune
	End   int    // Index after last rune
	Kind  Kind   // Type of token
}

// Stringify creates a string representation of the token.
func (tk scanTok) Stringify(indent int) string {
	in := strings.Repeat(" ", indent*2)

	sb := strings.Builder{}
	sb.WriteString("Token{")

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Val: ")
	sb.WriteString(strconv.QuoteToGraphic(tk.Val))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Line: ")
	sb.WriteString(strconv.Itoa(tk.Line))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Start: ")
	sb.WriteString(strconv.Itoa(tk.Start))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  End: ")
	sb.WriteString(strconv.Itoa(tk.End))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Kind: ")
	sb.WriteString(KindName(tk.Kind))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("}")

	return sb.String()
}
