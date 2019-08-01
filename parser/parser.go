package parser

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/expr/exprs"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Parse parses the input into an expression.
func Parse(in []token.Token) (out ctx.Expression, err fault.Fault) {
	return parse(in)
}

// parse parses the input into an expression.
func parse(in []token.Token) (out ctx.Expression, err fault.Fault) {

	if len(in) == 0 {
		err = missingTokens()
		return
	}

	if len(in) == 1 {
		return parseToken(in[0])
	}

	if in[0].Type == token.IDENTIFIER_EXPLICIT && in[1].Type == token.ASSIGNMENT {
		return parseAssignment(in)
	}

	return nil, notImplemented()
}

// parseToken parses an expression containing a single token.
func parseToken(in token.Token) (out ctx.Expression, err fault.Fault) {
	println("parseToken")

	switch in.Type {
	case token.LITERAL_NUMBER:
		out = exprs.Number{
			Number: in,
		}
	default:
		err = notImplemented()
	}

	return
}

// parseAssignment parses an assignment expression.
func parseAssignment(in []token.Token) (out ctx.Expression, err fault.Fault) {
	println("parseAssignment")

	err = requireMin(in, 3)
	if err != nil {
		return
	}

	right, err := parse(in[2:])
	if err != nil {
		return
	}

	out = exprs.Assignment{
		Identifier: in[0],
		Operator:   in[1],
		Right:      right,
	}

	return
}

// reverseTokens reverses an array of tokens.
func reverseTokens(in []token.Token) []token.Token {
	size := len(in)
	out := make([]token.Token, size)

	for i, v := range in {
		i = size - 1 - i
		out[i] = v
	}

	return out
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
