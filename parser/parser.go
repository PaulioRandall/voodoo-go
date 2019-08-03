package parser

import (
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses a token array into a statement.
func Parse(in []Token) (Statement, Fault) {

	a, in := divideOnAssignment(in)
	if a == nil {
		return nil, notImplemented()
	}

	op, left, err := parseAssignment(a)
	if err != nil {
		return nil, err
	}

	right, err := parseExpressions(in)
	if err != nil {
		return nil, err
	}

	out := Assignment{
		Left:     left,
		Operator: op,
		Right:    right,
	}

	return out, nil
}

// notImplemented returns a fault if there's no implementation for a
// particular arrangement of tokens.
func notImplemented() Fault {
	return ParseFault{
		Msgs: []string{
			`I haven't coded that path yet!`,
		},
	}
}

// divideOnLastOperator returns the first part of the
// token array before the last operator followed by the remainder.
// Nil is returned as the first token array if there are no
// operators.
func divideOnLastOperator(in []Token) ([]Token, Token, []Token) {

	size := len(in)
	op := -1

	for i, tk := range in {
		if token.IsOperator(tk.Type) {
			op = i
		}
	}

	if op == -1 {
		return in, Token{}, nil
	}

	if op == 0 || op+1 >= size {
		// Operator is on the edge of the array which is not valid syntax.
		return nil, Token{}, nil
	}

	return in[:op], in[op], in[op+1:]
}

// parseOperatorExpression creates an operator expression  from the
// left and right input.
func parseOperatorExpression(left []Token, op Token, right []Token) (Expression, Fault) {
	size := len(right)

	switch {
	case size == 1:
		expr, err := parseExpression(left)
		if err != nil {
			// TODO: Fault
		}

		out := Operation{
			Left:     expr,
			Operator: op,
			Right:    Value{right[0]},
		}
		return out, nil
	}

	return nil, notImplemented()
}

// parseExpression parses the whole token array as a single
// expression.
func parseExpression(in []Token) (Expression, Fault) {

	// NEXT: Whats the next test case?

	var left []Token
	var op Token
	var right []Token

	for {

		left, op, right = divideOnLastOperator(in)
		if left == nil {
			// TODO: Fault
		}

		if op == (Token{}) {
			// TODO: Fault
		}

		if right == nil {
			break
		}

		return parseOperatorExpression(left, op, right)
	}

	if len(left) == 1 {
		out := Value{left[0]}
		return out, nil
	}

	return nil, notImplemented()
}

// parseExpressions parses the expression part of a statement
// to produce one or many expressions for the right side.
func parseExpressions(in []Token) (Expression, Fault) {

	split := splitOnToken(in, token.SEPARATOR_VALUE)
	exprs := make([]Expression, len(split))

	for i, v := range split {
		expr, err := parseExpression(v)
		if err != nil {
			return nil, err
		}

		exprs[i] = expr
	}

	out := Join{
		Exprs: exprs,
	}

	return out, nil
}

// splitOnToken splits the token array into slices on the
// tokens with the specified token type.
func splitOnToken(in []Token, delim token.TokenType) [][]Token {

	out := [][]Token{}
	start := 0
	size := len(in)

	// TODO: Don't split if within spell or function param braces

	for i := 0; i < size; i++ {
		if in[i].Type == delim {
			out = append(out, in[start:i])
			start = i + 1
		}
	}

	if start < size {
		out = append(out, in[start:size])
	}

	return out
}

// containsAssignment returns true if the input token array
// contains an assignment.
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

// parseAssignment parses the assignment part of a statment
// to produce an expression for the left side.
func parseAssignment(in []Token) (Token, Expression, Fault) {

	size := len(in)
	ids := make([]Token, size/2)

	assign := in[size-1]
	in = in[:size-1]

	split := splitOnToken(in, token.SEPARATOR_VALUE)

	for i, id := range split {
		if len(id) != 1 {
			// TODO: Fault
		}

		if id[0].Type != token.IDENTIFIER {
			// TODO: Fault
		}

		ids[i] = id[0]
	}

	out := List{
		Tokens: ids,
	}

	return assign, out, nil
}

// validateDelimiter validates the passed token is a value
// delimiter or assignment returning a fault if not.
func validateDelimiter(tk Token) Fault {
	if tk.Type != token.ASSIGNMENT && tk.Type != token.SEPARATOR_VALUE {
		return ParseFault{
			Msgs: []string{
				`Unexpected token type`,
				`Expected a value delimiter or assignment token`,
			},
		}
	}

	return nil
}

// validateIdentifier validates the passed token is an identifier
// returning a fault if not.
func validateIdentifier(tk Token) Fault {
	if tk.Type != token.IDENTIFIER {
		return ParseFault{
			Msgs: []string{
				`Can't assign to non-identifier`,
				`Expected an identifier`,
			},
		}
	}

	return nil
}

// isEven returns true if the input is odd.
func isEven(i int) bool {
	return i == 0 || i%2 == 0
}

// divideOnAssignment returns the assignment part of the
// token array or nil if there is no assignment part.
func divideOnAssignment(in []Token) ([]Token, []Token) {
	for i, tk := range in {
		if tk.Type == token.ASSIGNMENT {
			i++
			return in[:i], in[i:]
		}

		if tk.Type == token.LOGICAL_MATCH {
			return nil, in
		}
	}

	return nil, in
}
