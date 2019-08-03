package new_parser

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/scroll"
)

// ParseFault represents a generic fault whilst parsing
// an expression.
type ParseFault struct {
	Msgs []string
}

// Print satisfies the Fault interface.
func (err ParseFault) Print(sc *scroll.Scroll, line int) {
	fmt.Print("[PARSE ERROR] `")
	fmt.Print(sc.File)
	fmt.Println("`")

	fmt.Printf("%3d: %s\n", line, sc.Lines[line])

	for _, m := range err.Msgs {
		fmt.Print(`...`)
		fmt.Println(m)
	}

	fmt.Println()
}
