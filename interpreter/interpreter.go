package interpreter

import (
	"github.com/PaulioRandall/voodoo-go/lexer"
	"github.com/PaulioRandall/voodoo-go/scroll"
	"github.com/PaulioRandall/voodoo-go/shared"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// Execute runs a Voodoo scroll.
func Execute(sc *scroll.Scroll, scArgs []string) (shared.ExitCode, error) {

	line := sc.Next(nil)
	line = sc.Next(line) // Ignoring first line, shebang

	for line != nil {

		s, err := lexer.ScanLine(line.Val, line.Num)
		if err != nil {
			return 1, err
		}

		symbol.PrintlnSymbols(s)

		line = sc.Next(line)
	}

	return 0, nil
}
