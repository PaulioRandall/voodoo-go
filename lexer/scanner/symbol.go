package scanner

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// scanSymbol scans one or two rune symbols returning one of:
// - comparison operator
// - arithmetic operator
// - logical operator
// - opening or closing brace or bracket
// - value or key-value separator
// - void identifier
// - number range generator
func scanSymbol(itr *runer.RuneItr) (l *symbol.Lexeme, err fault.Fault) {

	start := itr.Index()
	var t symbol.SymbolType
	c := 0

	set := func(lexType symbol.SymbolType, runeCount int) {
		t = lexType
		c = runeCount
	}

	switch {
	case itr.IsNextStr(`<-`):
		set(symbol.ASSIGNMENT, 2)
	case itr.IsNextStr(`<=`):
		set(symbol.CMP_LESS_THAN_OR_EQUAL, 2)
	case itr.IsNext('<'):
		set(symbol.CMP_LESS_THAN, 1)
	case itr.IsNextStr(`>=`):
		set(symbol.CMP_GREATER_THAN_OR_EQUAL, 2)
	case itr.IsNext('>'):
		set(symbol.CMP_GREATER_THAN, 1)
	case itr.IsNextStr(`==`):
		set(symbol.CMP_EQUAL, 2)
	case itr.IsNextStr(`!=`):
		set(symbol.CMP_NOT_EQUAL, 2)
	case itr.IsNextStr(`=>`):
		set(symbol.LOGICAL_MATCH, 2)
	case itr.IsNext('!'):
		set(symbol.LOGICAL_NOT, 1)
	case itr.IsNextStr(`||`):
		set(symbol.LOGICAL_OR, 2)
	case itr.IsNextStr(`&&`):
		set(symbol.LOGICAL_AND, 2)
	case itr.IsNext('+'):
		set(symbol.CALC_ADD, 1)
	case itr.IsNext('-'):
		set(symbol.CALC_SUBTRACT, 1)
	case itr.IsNext('*'):
		set(symbol.CALC_MULTIPLY, 1)
	case itr.IsNext('/'):
		set(symbol.CALC_DIVIDE, 1)
	case itr.IsNext('%'):
		set(symbol.CALC_MODULO, 1)
	case itr.IsNext('('):
		set(symbol.PAREN_CURVY_OPEN, 1)
	case itr.IsNext(')'):
		set(symbol.PAREN_CURVY_CLOSE, 1)
	case itr.IsNext('['):
		set(symbol.PAREN_SQUARE_OPEN, 1)
	case itr.IsNext(']'):
		set(symbol.PAREN_SQUARE_CLOSE, 1)
	case itr.IsNext(','):
		set(symbol.SEPARATOR_VALUE, 1)
	case itr.IsNext(':'):
		set(symbol.SEPARATOR_KEY_VALUE, 1)
	case itr.IsNextStr(`..`):
		set(symbol.RANGE, 2)
	case itr.IsNext('_'):
		set(symbol.VOID, 1)
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

	l = &symbol.Lexeme{
		Val:   s,
		Start: start,
		End:   itr.Index(),
		Type:  t,
	}

	return
}
