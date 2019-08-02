package parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses the input into a stack of instructions followed by a
// stack of values.
func Parse(in []Token) (*ExeStack, *ValStack, Fault) {

	exes := EmptyExeStack()
	vals := EmptyValStack()
	assign := false

	for len(in) > 0 {
		tk := in[0]

		switch {
		case isValueOrID(tk):
			vals.Push(tk)
			in = in[1:]
		case isAssignment(tk):
			assign = true
			fallthrough
		case isOperator(tk):
			var exe Exe
			exe, in = parseOperation(in)
			exes.Push(exe)
		default:
			return nil, nil, notImplemented()
		}
	}

	if assign {
		allButOne := exes.Len() - 1
		exes.Sink(allButOne)
		exes.Reverse()

		allButOne = vals.Len() - 1
		vals.Sink(allButOne)
		vals.Reverse()
	}

	return exes, vals, nil
}

// parseOperation parses an operation.
func parseOperation(in []Token) (Exe, []Token) {
	exe := Exe{
		Token:   in[0],
		Params:  2,
		Returns: 1,
	}
	return exe, in[1:]
}

// isAssignment returns true if the token is an assignment.
func isAssignment(tk Token) bool {
	return tk.Type == token.ASSIGNMENT
}

// isOperator returns true if the token is an operation.
func isOperator(tk Token) bool {
	switch tk.Type {
	case token.CALC_ADD:
		return true
	case token.CALC_SUBTRACT:
		return true
	}

	return false
}

// isValueOrID returns true if the token is a value or identifier.
func isValueOrID(tk Token) bool {
	switch tk.Type {
	case token.IDENTIFIER_EXPLICIT:
		return true
	case token.LITERAL_NUMBER:
		return true
	}

	return false
}

// requireMin returns a fault if the length of the token array
// is less than the minimum number required.
func requireMin(in []Token, min int) Fault {
	if len(in) >= min {
		return nil
	}

	return ParseFault{
		Type: `Not implemented`,
		Msgs: []string{
			`I haven't coded that path yet!`,
		},
	}
}

// notImplemented returns a fault if there's no implementation for a
// particular arrangement of tokens.
func notImplemented() Fault {
	return ParseFault{
		Type: `Not implemented`,
		Msgs: []string{
			`I haven't coded that path yet!`,
		},
	}
}

// missingTokens returns a fault if there are no tokens supplied for
// parsing when some were expected.
func missingTokens() Fault {
	return ParseFault{
		Type: `Missing tokens`,
		Msgs: []string{
			`I expected some tokens to parse`,
		},
	}
}
