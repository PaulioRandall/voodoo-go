package token

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
