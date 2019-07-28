package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
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
func scanSymbol(itr *runer.RuneItr) (tk *token.Token, err fault.Fault) {

	start := itr.Index()
	var t token.TokenType
	c := 0

	set := func(lexType token.TokenType, runeCount int) {
		t = lexType
		c = runeCount
	}

	switch {
	case itr.IsNextStr(`<-`):
		set(token.ASSIGNMENT, 2)
	case itr.IsNextStr(`<=`):
		set(token.CMP_LESS_THAN_OR_EQUAL, 2)
	case itr.IsNext('<'):
		set(token.CMP_LESS_THAN, 1)
	case itr.IsNextStr(`>=`):
		set(token.CMP_GREATER_THAN_OR_EQUAL, 2)
	case itr.IsNext('>'):
		set(token.CMP_GREATER_THAN, 1)
	case itr.IsNextStr(`==`):
		set(token.CMP_EQUAL, 2)
	case itr.IsNextStr(`!=`):
		set(token.CMP_NOT_EQUAL, 2)
	case itr.IsNextStr(`=>`):
		set(token.LOGICAL_MATCH, 2)
	case itr.IsNext('!'):
		set(token.LOGICAL_NOT, 1)
	case itr.IsNextStr(`||`):
		set(token.LOGICAL_OR, 2)
	case itr.IsNextStr(`&&`):
		set(token.LOGICAL_AND, 2)
	case itr.IsNext('+'):
		set(token.CALC_ADD, 1)
	case itr.IsNext('-'):
		set(token.CALC_SUBTRACT, 1)
	case itr.IsNext('*'):
		set(token.CALC_MULTIPLY, 1)
	case itr.IsNext('/'):
		set(token.CALC_DIVIDE, 1)
	case itr.IsNext('%'):
		set(token.CALC_MODULO, 1)
	case itr.IsNext('('):
		set(token.PAREN_CURVY_OPEN, 1)
	case itr.IsNext(')'):
		set(token.PAREN_CURVY_CLOSE, 1)
	case itr.IsNext('['):
		set(token.PAREN_SQUARE_OPEN, 1)
	case itr.IsNext(']'):
		set(token.PAREN_SQUARE_CLOSE, 1)
	case itr.IsNext(','):
		set(token.SEPARATOR_VALUE, 1)
	case itr.IsNext(':'):
		set(token.SEPARATOR_KEY_VALUE, 1)
	case itr.IsNextStr(`..`):
		set(token.RANGE, 2)
	case itr.IsNext('_'):
		set(token.VOID, 1)
	default:
		ru := itr.NextRune()
		m := "I don't know what this symbol means '" + string(ru) + "'"
		err = fault.Sym(m).From(start).To(itr.Index())
		return
	}

	s, e := itr.NextStr(c)
	if e != nil {
		err = fault.Sym(err.Error()).From(start).To(itr.Index())
		return
	}

	tk = &token.Token{
		Val:   s,
		Start: start,
		End:   itr.Index(),
		Type:  t,
	}

	return
}
