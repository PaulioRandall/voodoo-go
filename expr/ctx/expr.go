package ctx

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/scroll"
)

// Expression represents an expression that results in a value.
type Expression interface {

	// Evaluate evaluates the expression returning the resultant
	// value if there is one.
	Evaluate(*Context) (Value, fault.Fault)
}

// EvalFault represents a generic fault during an expressions
// evaluation.
type EvalFault struct {
	ExprType string
	Msgs     []string
}

// Print satisfies the Fault interface.
func (err EvalFault) Print(sc *scroll.Scroll, line int) {
	fmt.Print("[EVALUATION ERROR] `")
	fmt.Print(sc.File)
	fmt.Println("`")

	fmt.Printf("%3d: %s\n", line, sc.Lines[line])

	fmt.Print(`Evaluation type: `)
	fmt.Println(err.ExprType)

	for _, m := range err.Msgs {
		fmt.Print(`...`)
		fmt.Println(m)
	}

	fmt.Println()
}
