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

		l, err := lexer.ScanLine(line.Val, line.Num)
		if err != nil {
			return 1, err
		}

		lexeme.PrintlnLexemes(l)

		// NEXT: Strimmer -> Normalises lexemes by:
		//                  => Removing whitespace lexemes
		//                  => Removing comment lexemes
		//                  => Removing quote marks from string literals
		//                  => Removing underscores from numbers
		//                  => Converting all letters to lowercase
		// NEXT: Syntax Analyser

		line = sc.Next(line)
	}

	return 0, nil
}
