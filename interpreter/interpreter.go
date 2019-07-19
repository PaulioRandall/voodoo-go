package interpreter

import (
	"github.com/PaulioRandall/voodoo-go/lexer"
	"github.com/PaulioRandall/voodoo-go/scroll"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// ExitCode represents a program exit code
type ExitCode int

// Execute runs a Voodoo scroll.
func Execute(sc *scroll.Scroll, scArgs []string) (ExitCode, error) {

	line := sc.Next(nil)
	line = sc.Next(line) // Ignoring first line: shebang

	for line != nil {

		s, err := lexer.ScanLine(line.Val, line.Num)
		if err != nil {
			return 1, err
		}

		symbol.PrintlnSymbols(s)

		// NEXT:

		line = sc.Next(line)
	}

	return 0, nil
}
