package lexeme

// Token represents a lexeme that is useable by a
// syntax analyser.
type Token Lexeme

// String creates a string representation of the token.
func (tok Token) String() string {
	return Lexeme(tok).String()
}

// PrintlnTokens prints an array of tokens.
func PrintlnTokens(ts []Token) {
	ls := tokensToLexemes(ts)
	PrintlnLexemes(ls)
}

// PrintlnTokenTypes prints the types of an array of tokens.
func PrintlnTokenTypes(ts []Token) {
	ls := tokensToLexemes(ts)
	PrintlnLexemeTypes(ls)
}

// tokensToLexemes converts a token array to a lexeme array.
func tokensToLexemes(ts []Token) []Lexeme {
	ls := []Lexeme{}
	for _, v := range ts {
		ls = append(ls, Lexeme(v))
	}
	return ls
}
