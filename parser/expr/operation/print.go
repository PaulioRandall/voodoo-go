package operation

import (
	"github.com/PaulioRandall/voodoo-go/parser/expr"
)

// RuneTable satisfies the Expr interface.
func (o operation) RuneTable() expr.RuneTable {
	return nil
}
