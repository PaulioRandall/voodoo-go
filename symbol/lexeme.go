package symbol

// Lexeme represents the output symbol of lexing.
type Lexeme Symbol

// String creates a string representation of the lexeme.
func (lx Lexeme) String() string {
	return Symbol(lx).String()
}

// PrintlnLexemeValues prints the value of an array of lexemes.
func PrintlnLexemeValues(lxs []Lexeme) {
	f := func(sym Symbol) string {
		return sym.Val
	}
	s := lexemesToSymbols(lxs)
	printlnSymbols(s, f)
}

// PrintlnLexemeTypes prints the types of an array of lexemes.
func PrintlnLexemeTypes(lxs []Lexeme) {
	f := func(sym Symbol) string {
		return SymbolName(sym.Type)
	}
	s := lexemesToSymbols(lxs)
	printlnSymbols(s, f)
}

// lexemesToSymbols converts a lexeme array to a symbol array.
func lexemesToSymbols(lxs []Lexeme) []Symbol {
	s := make([]Symbol, len(lxs))
	for i, v := range lxs {
		s[i] = Symbol(v)
	}
	return s
}
