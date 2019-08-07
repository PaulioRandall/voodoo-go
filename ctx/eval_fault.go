package ctx

import (
	"github.com/PaulioRandall/voodoo-go/scroll"
)

// EvalFault represents a generic fault during an expressions
// evaluation.
type EvalFault struct {
	Msgs []string
}

// Print satisfies the Fault interface.
func (err EvalFault) Print(sc *scroll.Scroll, line int) {
	/*
		fmt.Print("[EVALUATION ERROR] `")
		fmt.Print(sc.File)
		fmt.Println("`")

		fmt.Printf("%3d: %s\n", line, sc.Lines[line])

		for _, m := range err.Msgs {
			fmt.Print(`...`)
			fmt.Println(m)
		}

		fmt.Println()
	*/
}
