package interpreter

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/lexer/scanner"
	"github.com/PaulioRandall/voodoo-go/lexer/strimmer"
	"github.com/PaulioRandall/voodoo-go/scroll"
	"github.com/PaulioRandall/voodoo-go/token"
)

// ExitCode represents a program exit code
type ExitCode int

// Execute runs a Voodoo scroll.
func Execute(sc *scroll.Scroll, scArgs []string) (ExitCode, error) {

	for i, line := range sc.Lines {
		if i == 0 {
			continue // Ignoring first line: shebang
		}

		tks, err := scanner.Scan(line)
		if err != nil {
			err = err.SetLine(i + 1)
			return 1, err
		}

		tks = strimmer.Strim(tks)
		token.PrintlnTokenValues(tks)
		token.PrintlnTokenTypes(tks)
		fmt.Println()
	}

	return 0, nil
}
