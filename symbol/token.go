package symbol

// Token represents a symbol that is ready for syntax parsing.
type Token Symbol

// String creates a string representation of the token.
func (tk Token) String() string {
	return Symbol(tk).String()
}

// PrintlnTokenValues prints the value of an array of tokens.
func PrintlnTokenValues(tks []Token) {
	f := func(sym Symbol) string {
		return sym.Val
	}
	s := tokensToSymbols(tks)
	printlnSymbols(s, f)
}

// PrintlnTokenTypes prints the types of an array of tokens.
func PrintlnTokenTypes(tks []Token) {
	f := func(sym Symbol) string {
		return SymbolName(sym.Type)
	}
	s := tokensToSymbols(tks)
	printlnSymbols(s, f)
}

// tokensToSymbols converts a token array to a symbol array.
func tokensToSymbols(tks []Token) []Symbol {
	s := make([]Symbol, len(tks))
	for i, v := range tks {
		s[i] = Symbol(v)
	}
	return s
}
