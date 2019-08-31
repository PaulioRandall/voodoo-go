package scantok

import (
	"strconv"
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// scanTok is an implementation of Token produced by the scanner.
type scanTok struct {
	text  string     // Token value
	line  int        // Line number in scroll
	start int        // Index of first rune
	end   int        // Index after last rune
	kind  token.Kind // Type of token
}

// New returns a new initilised Token.
func New(text string, line, start, end int, kind token.Kind) token.Token {
	return scanTok{
		text:  text,
		line:  line,
		start: start,
		end:   end,
		kind:  kind,
	}
}

// UpdateText updates the text within a scantok.
func UpdateText(tk token.Token, newText string) token.Token {
	stk, ok := tk.(scanTok)
	if !ok {
		panic("This isn't a scanTok")
	}
	stk.text = newText
	return stk
}

// Text satisfies the Token interface.
func (tk scanTok) Text() string {
	return tk.text
}

// Kind satisfies the Token interface.
func (tk scanTok) Kind() token.Kind {
	return tk.kind
}

// String returns the string representation of the scanned token.
func (tk scanTok) String() string {
	return tk.Stringify(0)
}

// Stringify creates a string representation of the token.
func (tk scanTok) Stringify(indent int) string {
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
	sb.WriteString(token.KindName(tk.kind))

	sb.WriteString("\n")
	sb.WriteString(in)
	sb.WriteString("}")

	return sb.String()
}
