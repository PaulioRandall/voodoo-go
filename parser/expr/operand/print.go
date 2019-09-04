package operand

import (
	"math"

	"github.com/PaulioRandall/voodoo-go/parser/expr"
)

// RuneTable satisfies the Expr interface.
func (o operand) RuneTable() expr.RuneTable {
	rt := expr.NewRuneTable(2, 8)
	rt.Imprint(0, 0, `Operand`)

	s := o.t.Text()
	if len(s) > 8 {
		rt.Imprint(1, 0, s[0:8])
	} else {
		col := float64(8 - len(s))
		col = math.Floor(col / 2)
		rt.Imprint(1, int(col), s)
	}

	rt.Filler(' ')
	return rt
}
