package parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses the input into a stack of instructions followed by a
// stack of values.
func Parse(in []Token) (*ExeStack, *ValStack, Fault) {

	exes := EmptyExeStack()
	vals := EmptyValStack()

	ids, in, err := preParseAssignment(in, exes, vals)
	if err != nil {
		return nil, nil, err
	}

	for len(in) > 0 {
		var exe Exe
		tk := in[0]

		switch {
		case isValueOrID(tk):
			vals.Push(tk)
			in = in[1:]
		case isValueSeparator(tk):
			in = in[1:]
		case isOperator(tk):
			exe, in = parseOperation(in)
			exes.Push(exe)
		default:
			return nil, nil, notImplemented()
		}
	}

	if ids > 0 {
		reverseStacks(ids, exes, vals)
	}

	return exes, vals, nil
}

// reverseStacks reverses the input stacks so the firt items added
// are the first to be processed. If an assignment operation is
// present then it is sunk to the bottom of the stack.
func reverseStacks(ids int, exes *ExeStack, vals *ValStack) {
	expr := exes.Len() - 1
	exes.Sink(expr)
	exes.Reverse()

	expr = vals.Len() - ids
	vals.Sink(expr)
	vals.Reverse()
}

// preParseAssignment parses the initial assignment in a statement.
func preParseAssignment(in []Token, exes *ExeStack, vals *ValStack) (int, []Token, Fault) {
	if !containsAssignment(in) {
		return 0, in, nil
	}

	size := len(in)
	ids := 0

	for i := 0; i < size; i += 2 {
		if (size - i) < 2 {
			return -1, nil, badAssignment("Odd number of assignment tokens")
		}

		id := in[i]
		punc := in[i+1]

		if id.Type == token.IDENTIFIER {
			ids++
			vals.Push(id)
		} else {
			return -1, nil, badAssignment("Expected identifier")
		}

		if isValueSeparator(punc) {
			continue
		}

		if isAssignment(punc) {
			e := Exe{
				Token:  punc,
				Params: ids * 2,
			}
			exes.Push(e)
			return ids, in[i+2:], nil
		}

		return -1, nil, badAssignment("Unexpected token")
	}

	return -1, nil, badAssignment("Unexpected end of statement")
}

// containsAssignment returns true if the input token array results
// in an assignment but not a match.
func containsAssignment(in []Token) bool {
	for _, tk := range in {
		if tk.Type == token.ASSIGNMENT {
			return true
		}

		if tk.Type == token.LOGICAL_MATCH {
			return false
		}
	}

	return false
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

// parseOperation parses an operation.
func parseOperation(in []Token) (Exe, []Token) {
	exe := Exe{
		Token:   in[0],
		Params:  2,
		Returns: 1,
	}
	return exe, in[1:]
}

// isValueSeparator returns true if the token is a value separator.
func isValueSeparator(tk Token) bool {
	return tk.Type == token.SEPARATOR_VALUE
}

// isAssignment returns true if the token is an assignment.
func isAssignment(tk Token) bool {
	return tk.Type == token.ASSIGNMENT
}

// isValueOrID returns true if the token is a value or identifier.
func isValueOrID(tk Token) bool {
	switch tk.Type {
	case token.IDENTIFIER:
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

// badAssignment returns a fault if the tokens that make up the
// initial assignment part of a statement are not valid.
func badAssignment(m string) Fault {
	return ParseFault{
		Type: `Bad assignment`,
		Msgs: []string{
			m,
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
