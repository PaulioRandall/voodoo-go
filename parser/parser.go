package parser

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Instruction represents a scroll instruction.
type Instruction struct {
	Token   token.Token
	Params  int // Number of parameters to pop from the value stack
	Returns int // Number of parameters to push on to the value stack
}

// Parse parses the input into a stack of instructions followed by a
// stack of values.
func Parse(in []token.Token) ([]Instruction, []token.Token, fault.Fault) {

	exeStack := []Instruction{}
	valueStack := []token.Token{}
	assign := false

	for len(in) > 0 {
		tk := in[0]

		switch {
		case isValueOrID(tk):
			valueStack = append(valueStack, tk)
			in = in[1:]
		case isAssignment(tk):
			assign = true
			fallthrough
		case isOperator(tk):
			var exe Instruction
			exe, in = parseOperation(in)
			exeStack = append(exeStack, exe)
		default:
			return nil, nil, notImplemented()
		}
	}

	if assign {
		val := valueStack[0]
		valueStack = valueStack[1:]
		valueStack = append(valueStack, val)

		exe := exeStack[0]
		exeStack = exeStack[1:]
		exeStack = append(exeStack, exe)
	}

	return exeStack, valueStack, nil
}

// reverseInstructions reverses an array of instructions.
func reverseInstructions(in []Instruction) []Instruction {
	for i := len(in)/2 - 1; i >= 0; i-- {
		opp := len(in) - 1 - i
		in[i], in[opp] = in[opp], in[i]
	}
	return in
}

// reverseTokens reverses an array of tokens.
func reverseTokens(in []token.Token) []token.Token {
	for i := len(in)/2 - 1; i >= 0; i-- {
		opp := len(in) - 1 - i
		in[i], in[opp] = in[opp], in[i]
	}
	return in
}

// parseOperation parses an operation.
func parseOperation(in []token.Token) (Instruction, []token.Token) {
	exe := Instruction{
		Token:   in[0],
		Params:  2,
		Returns: 1,
	}
	return exe, in[1:]
}

// isAssignment returns true if the token is an assignment.
func isAssignment(tk token.Token) bool {
	return tk.Type == token.ASSIGNMENT
}

// isOperator returns true if the token is an operation.
func isOperator(tk token.Token) bool {
	switch tk.Type {
	case token.CALC_ADD:
		return true
	case token.CALC_SUBTRACT:
		return true
	}

	return false
}

// isValueOrID returns true if the token is a value or identifier.
func isValueOrID(tk token.Token) bool {
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
func requireMin(in []token.Token, min int) fault.Fault {
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
func notImplemented() fault.Fault {
	return ParseFault{
		Type: `Not implemented`,
		Msgs: []string{
			`I haven't coded that path yet!`,
		},
	}
}

// missingTokens returns a fault if there are no tokens supplied for
// parsing when some were expected.
func missingTokens() fault.Fault {
	return ParseFault{
		Type: `Missing tokens`,
		Msgs: []string{
			`I expected some tokens to parse`,
		},
	}
}
