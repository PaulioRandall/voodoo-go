package interpreter

import (
	lx "github.com/PaulioRandall/voodoo-go/lexer"
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// Execute runs a Voodoo scroll.
func Execute(scroll *sc.Scroll, scrollArgs []string) (sh.ExitCode, error) {

	line := scroll.Next(nil)
	line = scroll.Next(line) // Ignoring first line, shebang

	for line != nil {

		s, err := lx.ScanLine(line.Val, line.Num)
		if err != nil {
			return 1, err
		}

		lx.PrintlnSymbols(s)

		line = scroll.Next(line)
	}

	return 0, nil
}
