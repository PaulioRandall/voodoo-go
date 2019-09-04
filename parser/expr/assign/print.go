package assign

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
)

// RuneTable satisfies the Expr interface.
func (a assign) RuneTable() expr.RuneTable {
	return nil
}
