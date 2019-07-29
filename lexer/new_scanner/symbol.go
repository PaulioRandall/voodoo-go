package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// scanSymbol scans one or two rune symbols returning one of:
// - comparison operator
// - arithmetic operator
// - logical operator
// - opening or closing brace or bracket
// - value or key-value separator
// - void identifier
// - number range generator
//
// Assumes that the input is not empty.
func scanSymbol(in []rune) (tk *token.Token, out []rune, err fault.Fault) {

	var t token.TokenType
	c := 0

	set := func(tkType token.TokenType, runeCount int) {
		t = tkType
		c = runeCount
	}

	switch {
	case startsWith(in, `<-`):
		set(token.ASSIGNMENT, 2)
	case startsWith(in, `<=`):
		set(token.CMP_LESS_THAN_OR_EQUAL, 2)
	case startsWith(in, `<`):
		set(token.CMP_LESS_THAN, 1)
	case startsWith(in, `>=`):
		set(token.CMP_GREATER_THAN_OR_EQUAL, 2)
	case startsWith(in, `>`):
		set(token.CMP_GREATER_THAN, 1)
	case startsWith(in, `==`):
		set(token.CMP_EQUAL, 2)
	case startsWith(in, `!=`):
		set(token.CMP_NOT_EQUAL, 2)
	case startsWith(in, `=>`):
		set(token.LOGICAL_MATCH, 2)
	case startsWith(in, `!`):
		set(token.LOGICAL_NOT, 1)
	case startsWith(in, `||`):
		set(token.LOGICAL_OR, 2)
	case startsWith(in, `&&`):
		set(token.LOGICAL_AND, 2)
	case startsWith(in, `+`):
		set(token.CALC_ADD, 1)
	case startsWith(in, `-`):
		set(token.CALC_SUBTRACT, 1)
	case startsWith(in, `*`):
		set(token.CALC_MULTIPLY, 1)
	case startsWith(in, `/`):
		set(token.CALC_DIVIDE, 1)
	case startsWith(in, `%`):
		set(token.CALC_MODULO, 1)
	case startsWith(in, `(`):
		set(token.PAREN_CURVY_OPEN, 1)
	case startsWith(in, `)`):
		set(token.PAREN_CURVY_CLOSE, 1)
	case startsWith(in, `[`):
		set(token.PAREN_SQUARE_OPEN, 1)
	case startsWith(in, `]`):
		set(token.PAREN_SQUARE_CLOSE, 1)
	case startsWith(in, `,`):
		set(token.SEPARATOR_VALUE, 1)
	case startsWith(in, `:`):
		set(token.SEPARATOR_KEY_VALUE, 1)
	case startsWith(in, `_`):
		set(token.VOID, 1)
	default:
		m := "I don't know what this symbol means '" + string(in[0]) + "'"
		err = fault.Sym(m)
		return
	}

	tk = &token.Token{
		Val:  string(in[:c]),
		Type: t,
	}

	out = in[c:]
	return
}
