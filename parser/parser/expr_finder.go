package parser

import (
	"errors"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// exprDelimFinder finds where the value delimiters are between the top level
// expressions of the input tokens.
type exprDelimFinder struct {
	braceStack []token.Token
}

// Find identifies where the delimiters between the top level expressions are
// for the input slice.
func (edf *exprDelimFinder) Find(in []token.Token) ([]int, error) {
	edf.braceStack = []token.Token{}
	out := []int{}

	for i, tk := range in {
		if edf.processNext(tk) {
			out = append(out, i)
		}

		if err := edf.checkNests(); err != nil {
			return nil, err
		}
	}

	if err := edf.checkNests(); err != nil {
		return nil, err
	}

	return out, nil
}

// nests returns the current level of nesting.
func (edf *exprDelimFinder) nests() int {
	return len(edf.braceStack)
}

// checkNests checks the level of nesting is never zero as this would indicate
// a closing brace preceeds a corresponding opening one.
func (edf *exprDelimFinder) checkNests() error {
	if edf.nests() < 0 {
		m := "Closing brace detected without preceeding opening brace"
		return errors.New(m)
	}
	return nil
}

// processNext processes the next token returning true if it is a value
// delimiter and splits two top level expressions.
func (edf *exprDelimFinder) processNext(tk token.Token) bool {
	switch {
	case edf.opensNesting(tk):
	case edf.closesNesting(tk):
	case edf.nests() != 0:
	case tk.Type == token.TT_VALUE_DELIM:
		return true
	}

	return false
}

// opensNesting returns true if the brace is an opener to a nested set of
// expressions. It will store the token on the brace stack if it is.
func (edf *exprDelimFinder) opensNesting(tk token.Token) bool {
	switch tk.Type {
	case token.TT_CURVED_OPEN, token.TT_SQUARE_OPEN:
		edf.braceStack = append(edf.braceStack, tk)
		return true
	}
	return false
}

// closesNesting returns true if the token is a closing brace that matches the
// last found opening brace. It will remove the opening token from the brace
// stack if true.
func (edf *exprDelimFinder) closesNesting(tk token.Token) bool {
	last := edf.nests() - 1
	if last < 0 {
		return false
	}

	openType := edf.braceStack[last].Type

	switch {
	case openType == token.TT_CURVED_OPEN && tk.Type == token.TT_CURVED_CLOSE:
	case openType == token.TT_SQUARE_OPEN && tk.Type == token.TT_SQUARE_CLOSE:
	default:
		return false
	}

	edf.braceStack = edf.braceStack[:last]
	return true
}
