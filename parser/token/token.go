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

	// Line returns the line index of the token within its scroll.
	Line() int

	// Start returns the column index of the first rune within the line.
	Start() int

	// End returns the column index after the last rune within the line.
	End() int

	// Kind returns the type of token.
	Kind() Kind

	// String returns a string representation of the token.
	String() string
}

// token implements Token.
type token struct {
	text  string // Token value
	line  int    // Line number in scroll
	start int    // Index of first rune
	end   int    // Index after last rune
	kind  Kind   // Type of token
}

// New returns a new initilised Token.
func New(text string, line, start, end int, kind Kind) Token {
	return token{
		text:  text,
		line:  line,
		start: start,
		end:   end,
		kind:  kind,
	}
}

// Copy performs a shallow copy of a token array.
func Copy(in []Token) []Token {
	if in == nil {
		return nil
	}

	out := make([]Token, len(in))
	for i, _ := range in {
		out[i] = in[i]
	}

	return out
}

// UpdateText updates the text within a token.
func UpdateText(tk Token, newText string) Token {
	stk, ok := tk.(token)
	if !ok {
		panic("This isn't a known token implementation")
	}
	stk.text = newText
	return stk
}

// Text satisfies the Token interface.
func (tk token) Text() string {
	return tk.text
}

// Line satisfies the Token interface.
func (tk token) Line() int {
	return tk.line
}

// Start satisfies the Token interface.
func (tk token) Start() int {
	return tk.start
}

// End satisfies the Token interface.
func (tk token) End() int {
	return tk.end
}

// Kind satisfies the Token interface.
func (tk token) Kind() Kind {
	return tk.kind
}

// String returns the string representation of the scanned token.
func (tk token) String() string {
	return tk.Stringify(0)
}

// Stringify creates a string representation of the token.
func (tk token) Stringify(indent int) string {
	in := strings.Repeat(" ", indent*2)

	sb := strings.Builder{}
	sb.WriteString("Token{")

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Val: ")
	sb.WriteString(strconv.QuoteToGraphic(tk.text))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Line: ")
	sb.WriteString(strconv.Itoa(tk.line))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Start: ")
	sb.WriteString(strconv.Itoa(tk.start))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  End: ")
	sb.WriteString(strconv.Itoa(tk.end))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("  Kind: ")
	sb.WriteString(KindName(tk.kind))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("}")

	return sb.String()
}
