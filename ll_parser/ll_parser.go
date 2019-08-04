package ll_parser

import (
	"strconv"

	"github.com/PaulioRandall/voodoo-go/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses the input token array into a statement.
func Parse(in []token.Token) (ctx.Statement, fault.Fault) {

	s := newStack(5)
	express := false

loop:
	for i, tk := range in {
		switch {
		case tk.Type == token.END_OF_STATMENT:
			break loop

		case tk.Type == token.IDENTIFIER:
			v := newIDValue(tk)

			s.Push(v)
			if express {
				buildOperation(s)
				express = false
			}

		case tk.Type == token.ASSIGNMENT:
			s.Push(tk)

		case token.IsOperator(tk.Type):
			s.Push(tk)
			t := lookAhead(in, i)
			if t == token.LITERAL_NUMBER || t == token.IDENTIFIER {
				express = true
			}

		case tk.Type == token.LITERAL_NUMBER:
			v, err := newNumberValue(tk)
			if err != nil {
				return nil, err
			}

			s.Push(v)
			if express {
				buildOperation(s)
				express = false
			}

		default:
			return nil, notImplemented()
		}
	}

	buildAssignment(s)
	return s.Pop().(ctx.Statement), nil
}

// newIDValue creates a new identifier value.
func newIDValue(tk token.Token) ctx.Value {
	return ctx.Value{
		Token: tk,
		Val:   tk.Val,
		Type:  ctx.IDENTIFIER_TYPE,
	}
}

// lookAhead returns the type of the next token.
func lookAhead(in []token.Token, i int) token.TokenType {
	return in[i+1].Type
}

// buildAssignment takes the last three items from the stack
// and builds an assignment statement from them.
func buildAssignment(s *stack) {
	ass := ctx.Assignment{
		Right:    s.Pop().(ctx.Expression),
		Operator: s.Pop().(token.Token),
		Left:     s.Pop().(ctx.Value),
	}

	s.Push(ass)
}

// buildOperation takes the last three items from the stack
// and builds an expression from them.
func buildOperation(s *stack) {
	exp := ctx.Operation{
		Right:    s.Pop().(ctx.Expression),
		Operator: s.Pop().(token.Token),
		Left:     s.Pop().(ctx.Expression),
	}

	s.Push(exp)
}

// newNumberValue creates a new number value.
func newNumberValue(tk token.Token) (ctx.Value, fault.Fault) {
	f, err := parseNumber(tk.Val)
	if err != nil {
		return ctx.Value{}, err
	}

	out := ctx.Value{
		Token: tk,
		Val:   f,
		Type:  ctx.NUMBER_TYPE,
	}

	return out, nil
}

// parseNumber parses a number value.
func parseNumber(n string) (float64, fault.Fault) {
	f, err := strconv.ParseFloat(n, 64)
	if err != nil {
		// TODO: Fault
		return 0, notImplemented()
	}

	return f, nil
}

// notImplemented returns a fault if there's no implementation for a
// particular arrangement of tokens.
func notImplemented() fault.Fault {
	return ParseFault{
		Msgs: []string{
			`I haven't coded that path yet!`,
		},
	}
}
