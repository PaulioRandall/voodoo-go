package symbol

// Token represents a lexeme that is useable by a
// syntax analyser.
type Token Lexeme

// String creates a string representation of the token.
func (tok Token) String() string {
	return Symbol(tok).String()
}

// PrintlnTokenValues prints the value of an array of tokens.
func PrintlnTokenValues(ts []Token) {
	ss := tokensToSymbols(ts)
	f := func(s Symbol) string {
		return s.Val
	}
	printlnSymbols(ss, f)
}

// PrintlnTokenTypes prints the types of an array of tokens.
func PrintlnTokenTypes(ts []Token) {
	ss := tokensToSymbols(ts)
	f := func(s Symbol) string {
		return nameOfType(s.Type)
	}
	printlnSymbols(ss, f)
}

// tokensToSymbols converts a token array to a symbol array.
func tokensToSymbols(ts []Token) []Symbol {
	ss := []Symbol{}
	for _, v := range ts {
		ss = append(ss, Symbol(v))
	}
	return ss
}
