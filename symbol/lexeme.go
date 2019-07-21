package symbol

// Symbol represents a rune or string within the code
// that equates to a meaningful item within the grammer
// rules.
type Lexeme Symbol

// String creates a string representation of the lexeme.
func (lex Lexeme) String() string {
	return Symbol(lex).String()
}

// PrintlnLexemeValues prints the value of an array of lexemes.
func PrintlnLexemeValues(ls []Lexeme) {
	ss := lexemesToSymbols(ls)
	f := func(s Symbol) string {
		return s.Val
	}
	printlnSymbols(ss, f)
}

// PrintlnLexemeTypes prints the types of an array of lexemes.
func PrintlnLexemeTypes(ls []Lexeme) {
	ss := lexemesToSymbols(ls)
	f := func(s Symbol) string {
		return nameOfType(s.Type)
	}
	printlnSymbols(ss, f)
}

// lexemesToSymbols converts a lexeme array to a symbol array.
func lexemesToSymbols(ls []Lexeme) []Symbol {
	ss := []Symbol{}
	for _, v := range ls {
		ss = append(ss, Symbol(v))
	}
	return ss
}
