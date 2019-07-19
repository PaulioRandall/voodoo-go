package interpreter

import (
	"github.com/PaulioRandall/voodoo-go/lexeme"
	"github.com/PaulioRandall/voodoo-go/lexer"
	"github.com/PaulioRandall/voodoo-go/scroll"
)

// ExitCode represents a program exit code
type ExitCode int

// Execute runs a Voodoo scroll.
func Execute(sc *scroll.Scroll, scArgs []string) (ExitCode, error) {

	line := sc.Next(nil)
	line = sc.Next(line) // Ignoring first line: shebang

	for line != nil {

		lex, err := lexer.ScanLine(line.Val, line.Num)
		if err != nil {
			return 1, err
		}

		lexeme.PrintlnSymbols(lex)

		// NEXT: Rename 'symbol' to 'lexeme'

		line = sc.Next(line)
	}

	return 0, nil
}
