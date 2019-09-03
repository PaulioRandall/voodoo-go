package token

// Token represents a token produced by lexical analysis.
// I.e. identifier, operator, punctionation, etc.
type Token interface {

	// Text returns the textual representation of the token. For scanned tokens
	// this will always be the actual text that represents the token while others
	// may choose to use as they please.
	Text() string

	Line() int

	Start() int

	End() int

	// Kind returns the type of token.
	Kind() Kind

	// String returns a string representation of the token.
	String() string
}

// Copy performs a deep copy of a token array.
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
